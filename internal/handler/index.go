package handler

import (
	"github.com/gin-gonic/gin"
)

// Index 首页展示和逻辑处理
func Index(ctx *gin.Context) {
	ctx.JSON(200, "欢迎来到我的世界!")
}
