package controllers

import (
	"Exchangeapp/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// 点赞文章
func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")
	//构造Redis里存点赞数的key
	likeKey := "article:" + articleID + ":likes"

	//key+1
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked the article"})
}

// 获取文章点赞数
func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"

	//从Redis获取这个key的值（点赞数）
	likes, err := global.RedisDB.Get(likeKey).Result()

	//判断Redis里是否有这个key，如果没有，说明点赞数为0
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
