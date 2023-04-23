package handler

import (
	"github.com/Powehi-cs/seckill/internal/utils"
	"github.com/gin-gonic/gin"
)

// AdminLogin 管理员登录
func AdminLogin(ctx *gin.Context) {

}

// AdminLoginPage 管理员登录页面
func AdminLoginPage(ctx *gin.Context) {
	if utils.Check(ctx) {
		ctx.JSON(302, "/admin/")
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
