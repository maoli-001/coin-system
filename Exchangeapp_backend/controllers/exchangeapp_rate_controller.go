package controllers

import (
	"Exchangeapp/global"
	"Exchangeapp/models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建汇率
func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate

	// 绑定前端传来的汇率信息到exchangeRate,有问题返回400错误
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置汇率的创建时间
	exchangeRate.Date = time.Now()

	// 自动迁移数据库表结构，不存在表就创建存在则更新
	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//创建数据库表
	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回创建成功的汇率信息
	ctx.JSON(http.StatusCreated, exchangeRate)
}

// 获取汇率列表
func GetExchangeRates(ctx *gin.Context) {
	//切片对象
	var exchangeRates []models.ExchangeRate

	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, exchangeRates)
}
