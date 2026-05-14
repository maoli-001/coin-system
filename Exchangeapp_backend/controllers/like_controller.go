package controllers

import (
	"Exchangeapp/global"
	"Exchangeapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// 点赞文章/取消点赞
func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")

	//从上下文获取用户名
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
		return
	}
	usernameStr := username.(string)

	//构造Redis里存点赞数和点赞用户的key
	likeKey := "article:" + articleID + ":likes"
	userLikeKey := "article:" + articleID + ":users"
	//用户点赞的文章列表
	userLikedArticlesKey := "user:" + usernameStr + ":likeArticles"

	//判断用户是否已经点赞过
	isLiked, err := global.RedisDB.SIsMember(userLikeKey, usernameStr).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//如果已经点赞过，则取消点赞
	if isLiked {
		//删除用户的点赞记录
		if err := global.RedisDB.SRem(userLikeKey, usernameStr).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//删除用户的点赞文章记录
		global.RedisDB.SRem(userLikedArticlesKey, articleID)

		//更新点赞数
		likes, err := global.RedisDB.Decr(likeKey).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully unliked the article",
			"likes": likes,
			"liked": false})
	} else {
		//如果没有点赞过，则点赞
		//添加用户的点赞记录
		if err := global.RedisDB.SAdd(userLikeKey, usernameStr).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//添加用户的点赞文章记录
		global.RedisDB.SAdd(userLikedArticlesKey, articleID)

		//更新点赞数
		likes, err := global.RedisDB.Incr(likeKey).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked the article",
			"likes": likes,
			"liked": true})
	}
}

// 获取文章点赞数
func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"
	userLikeKey := "article:" + articleID + ":users"

	//从Redis获取这个key的值（点赞数）
	likes, err := global.RedisDB.Get(likeKey).Result()

	//判断Redis里是否有这个key，如果没有，说明点赞数为0
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//获取当前用户的点赞状态
	isLiked := false
	username, exists := ctx.Get("username")
	if exists {
		usernameStr := username.(string)
		isLiked, _ = global.RedisDB.SIsMember(userLikeKey, usernameStr).Result()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"likes":   likes,
		"isLiked": isLiked,
	})
}

// 获取用户的点赞文章列表
func GetUserLikedArticles(ctx *gin.Context) {
	//从上下文获取用户名
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "请先登录"})
		return
	}

	//构造用户名的点赞的Key
	usernameStr := username.(string)
	userLikedArticlesKey := "user:" + usernameStr + ":likeArticles"

	//从缓存中获取用户点赞的文章ID
	articleIDs, err := global.RedisDB.SMembers(userLikedArticlesKey).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//如果没有点赞记录，返回空数组
	if len(articleIDs) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"articles": []models.Article{},
			"total":    0,
		})
		return
	}

	//从MySql中获取文章详情
	var articles []models.Article
	if err := global.Db.Where("id IN ?", articleIDs).Find(&articles).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//返回文章列表
	ctx.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    len(articles),
	})
}
