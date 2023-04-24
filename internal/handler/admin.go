package handler

import (
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/Powehi-cs/seckill/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// AdminLogin 管理员登录
func AdminLogin(ctx *gin.Context) {

}

// AdminLoginPage 管理员登录页面
func AdminLoginPage(ctx *gin.Context) {
	if utils.Check(ctx) {
		ctx.JSON(200, utils.GetGinH(utils.LoginSuccess, "管理员登录页面"))
	}
}

// Search 查找商品
func Search(ctx *gin.Context) {

}

// Insert 增加商品
func Insert(ctx *gin.Context) {

}

// Update 修改商品
func Update(ctx *gin.Context) {

}

// Delete 删除商品
func Delete(ctx *gin.Context) {

}

// SetSecKillProduct 将商品设为秒杀商品
func SetSecKillProduct(ctx *gin.Context) {
	productID := ctx.Param("product_id")
	id, err := strconv.Atoi(productID)
	errors.PrintInStdout(err)

	TransSecKillProduct(ctx, uint(id), 30*time.Minute)
}

// TransSecKillProduct 将mysql中的秒杀商品转移到redis中
func TransSecKillProduct(ctx *gin.Context, productID uint, timeDuration time.Duration) {
	rdb := database.GetRedis()

	var product model.Product
	product.ProductID = productID
	err := product.Select()
	errors.PrintInStdout(err)

	rdb.Set(ctx, strconv.Itoa(int(product.ProductID)), product.Inventory, timeDuration)
}
