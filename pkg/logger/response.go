package logger

import (
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, msg string) {
	ctx.JSON(200, gin.H{
		"msg": msg,
	})
}

func Fail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"msg": msg,
	})
}

func Redirect(ctx *gin.Context, path string) {
	ctx.Redirect(302, path)
}
