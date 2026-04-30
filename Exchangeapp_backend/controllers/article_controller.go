package controllers

import (
	"Exchangeapp/global"
	"Exchangeapp/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// 缓存文章列表的 key
var cacheKey = "articles"

// 创建文章接口
func CreateArticle(ctx *gin.Context) {
	//创建对象
	var article models.Article

	//绑定前端传来的文章信息到article,有问题返回400错误CreateArticle
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//自动迁移到数据库表结构，不存在就创建，存在就更新
	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//插入数据库
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//删除缓存以便当有用户上传文章时更新缓存
	if err := global.RedisDB.Del(cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//返回创建成功的文章信息
	ctx.JSON(http.StatusCreated, article)
}

// 获取文章列表接口
func GetArticles(ctx *gin.Context) {
	//从缓存中获取文章列表
	cacheData, err := global.RedisDB.Get(cacheKey).Result()
	//判断Redis里是否有缓存,没有缓存就从数据库中那并加入缓存
	if err == redis.Nil {
		//切片对象
		var articles []models.Article

		//查询数据库
		if err := global.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "No articles found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		//序列化为JSON
		articleJSON, err := json.Marshal(articles)

		//判断序列化是否成功
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//存入缓存,缓存时间10分钟
		if err := global.RedisDB.Set(cacheKey, articleJSON, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, articles)
		return

	} else if err != nil { //发生错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	} else { //有缓存
		var articles []models.Article

		//反序列化JSON为go结构体
		if err := json.Unmarshal([]byte(cacheData), &articles); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//返回缓存数据
		ctx.JSON(http.StatusOK, articles)
	}
}

// 根据文章 ID 从数据库查询单篇文章
func GetArticlesByID(ctx *gin.Context) {
	// 拿 URL 里的 ID
	id := ctx.Param("id")

	var article models.Article

	if err := global.Db.Where("id=?", id).First(&article).Error; err != nil {
		//判断错误是不是没找到记录
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	//返回文章信息
	ctx.JSON(http.StatusOK, article)
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	// 拿 URL 里的 ID
	id := ctx.Param("id")
	var article models.Article

	//先查看文章还在吗
	if err := global.Db.Where("id=?", id).First(&article).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	//删除数据库里的文章
	if err := global.Db.Delete(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//删除缓存
	if err := global.RedisDB.Del(cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//返回删除成功的消息
	ctx.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
