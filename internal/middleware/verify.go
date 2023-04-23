package middleware

import (
	"github.com/Powehi-cs/seckill/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		// 如果token解析成功，则继续
		if claims, ok := utils.ParseToken(tokenString); ok {
			ctx.Set("name", claims["name"])
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(400, "token错误!")
		return
	}
}
