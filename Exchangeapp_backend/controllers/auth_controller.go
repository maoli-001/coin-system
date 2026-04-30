package controllers

import (
	"Exchangeapp/global"
	"Exchangeapp/models"
	"Exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册接口
// 接收前端床底的JSON数据，成功返回token，失败返回错误
func Register(ctx *gin.Context) {
	//声明一个User类型的变量user
	var user models.User

	//绑定前端传来的用户信息到user,有问题返回400错误
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//调用工具函数加密密码
	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//将加密后的密码赋值给user
	user.Password = hashedPwd

	//JWT令牌问题
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//自动迁移数据库表结构，不存在表就创建存在则更新
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//将用户数据插入数据库
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//所有步骤成功，返回200和token
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// 登录接口
func Login(ctx *gin.Context) {
	//声明用户输入的结构体，同时用标签确保前后端间大小写对应问题
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	//将用户输入绑定到input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	//查询数据库中是否有该用户
	if err := global.Db.Where("username=?", input.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	//验证密码的函数应写utils.go中，这里直接调用
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	//生成JWT令牌
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
