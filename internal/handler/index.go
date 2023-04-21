package handler

import (
	"github.com/Powehi-cs/seckill/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Index 首页展示和逻辑处理
func Index(ctx *gin.Context) {
	logger.Success(ctx, "欢迎来到主页面")
}
