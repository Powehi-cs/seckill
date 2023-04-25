package handler

import (
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/database"
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
		name, _ := ctx.Get("name")
		if name != "yuan cheng" {
			ctx.JSON(200, utils.GetGinH(utils.LoginFail, "您不是管理员"))
			return
		}
		ctx.Header("Location", "/admin")
		ctx.JSON(200, utils.GetGinH(utils.LoginSuccess, "管理员登录成功"))
		return
	}
	ctx.JSON(200, utils.GetGinH(utils.LoginFail, "欢迎来到登录页面"))
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
	if err != nil {
		ctx.JSON(200, utils.GetGinH(utils.Error, "无法将该字符串类型转换为整型"))
		return
	}

	if TransSecKillProduct(ctx, uint(id), 30*time.Minute) {
		ctx.JSON(200, utils.GetGinH(utils.Success, "秒伤商品设置成功"))
		return
	}

	ctx.JSON(200, utils.GetGinH(utils.Fail, "秒杀商品设置失败"))
}

// TransSecKillProduct 将mysql中的秒杀商品转移到redis中
func TransSecKillProduct(ctx *gin.Context, productID uint, timeDuration time.Duration) bool {
	rdb := database.GetRedis()

	var product model.Product
	product.ProductID = productID
	err := product.Select()
	if err != nil {
		return false
	}

	rdb.Set(ctx, "inventory_"+strconv.Itoa(int(product.ProductID)), product.Inventory, timeDuration)

	return true
}
