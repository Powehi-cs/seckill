package utils

import (
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// TransSecKillProduct 将mysql中的秒杀商品转移到redis中
func TransSecKillProduct(ctx *gin.Context, productID uint, timeDuration time.Duration) {
	rdb := database.GetRedis()

	var product model.Product
	product.ProductID = productID
	err := product.Select()
	errors.PrintInStdout(err)

	rdb.Set(ctx, strconv.Itoa(int(product.ProductID)), product.Inventory, timeDuration)
}
