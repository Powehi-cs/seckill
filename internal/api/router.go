package api

import (
	"github.com/Powehi-cs/seckill/internal/handler"
	"github.com/Powehi-cs/seckill/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Router 路由设置
func Router(router *gin.Engine) {
	router.Use(middleware.Cors())

	router.GET("/", handler.Index) // 首页逻辑

	// 用户注册
	router.GET("/register", handler.RegisterPage) // 注册页面
	router.POST("/register", handler.Register)    // 注册逻辑

	// 用户登录
	router.GET("/login", middleware.AuthVerify(), handler.LoginPage) // 登录页面
	router.POST("/login", handler.Login)                             // 登录逻辑

	// 秒杀路由
	products := router.Group("/:user_id/:product_id/").Use(middleware.AuthVerify(), middleware.ConsistentHash())
	{
		// 用户秒杀
		products.GET("", handler.ProductPage) // 单个产品页面
		products.PUT("", handler.SecKill)     // 产品秒杀逻辑
	}

	// 管理员(唯一管理员)
	router.GET("/admin/login", middleware.AuthVerify(), handler.AdminLoginPage) // 管理员登录页面
	router.POST("/admin/login", handler.AdminLogin)                             // 管理员登录逻辑
	admin := router.Group("/admin/").Use(middleware.AuthVerify())
	{
		admin.GET("", handler.Search)    // 管理员查询商品
		admin.POST("", handler.Insert)   // 管理员增加商品
		admin.PUT("", handler.Update)    // 管理员修改商品
		admin.DELETE("", handler.Delete) // 管理员删除商品
	}
}
