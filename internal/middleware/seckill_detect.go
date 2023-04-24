package middleware

import (
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/utils"
	"github.com/gin-gonic/gin"
)

func IsStart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID := ctx.Param("product_id")

		rdb := database.GetRedis()
		_, err := rdb.Get(ctx, "inventory_"+productID).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(200, utils.GetGinH(400, "秒杀尚未开始"))
			return
		}
		ctx.Next()
	}
}
