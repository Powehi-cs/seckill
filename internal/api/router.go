package api

import (
	"github.com/Powehi-cs/seckill/internal/handler"
	"github.com/Powehi-cs/seckill/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Router 路由设置
func Router(router *gin.Engine) {

	router.Use(middleware.Verify) // 过滤器，验证用户合法性

	// 首页
	router.GET("/", handler.Index) // 首页逻辑

	// 用户路由
	users := router.Group("/users/")
	{
		// 用户注册
		users.GET("register", handler.RegisterPage) // 注册页面
		users.POST("register", handler.Register)    // 注册逻辑
		// 用户登录
		users.GET("login", handler.LoginPage) // 登录页面
		users.POST("login", handler.Login)    // 登录逻辑
	}

	// 秒杀路由
	products := router.Group("/:user_id/:product_id/")
	{
		// 用户秒杀
		products.GET("", handler.ProductPage) // 单个产品页面
		products.PUT("", handler.SecKill)     // 产品秒杀逻辑
	}

	// 管理员(唯一管理员)
	admin := router.Group("/admin/")
	{
		admin.GET("login", handler.AdminLoginPage) // 管理员登录页面
		admin.POST("login", handler.AdminLogin)    // 管理员登录逻辑

		admin.GET("", handler.Search)    // 管理员查询商品
		admin.POST("", handler.Insert)   // 管理员增加商品
		admin.PUT("", handler.Update)    // 管理员修改商品
		admin.DELETE("", handler.Delete) // 管理员删除商品
	}
}
