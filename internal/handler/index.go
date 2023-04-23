package handler

import (
	"github.com/Powehi-cs/seckill/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Index 首页展示和逻辑处理
func Index(ctx *gin.Context) {
	ctx.JSON(200, utils.GetGinH(utils.Success, "这是首页"))
}
