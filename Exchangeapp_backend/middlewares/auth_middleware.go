package middlewares

import (
	"Exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证token中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//拿到登录信息
		token := ctx.GetHeader("Authorization")
		//如果为空返回错误并终止
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			//结束请求
			ctx.Abort()
			return
		}

		//调用工具函数解析token
		username, err := utils.ParseJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		//将用户名存入上下文
		ctx.Set("username", username)
		//放行
		ctx.Next()
	}
}
