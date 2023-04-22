package handler

import (
	"github.com/Powehi-cs/seckill/internal/middleware"
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	// 获取用户登录账号和密码
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		logger.Fail(ctx, 400, "请输入账号和密码")
		return
	}

	// 存入mysql和redis
	if !user.Create(ctx) {
		logger.Fail(ctx, 403, "重复注册")
		return
	}

	logger.Success(ctx, "注册成功")
}

// RegisterPage 用户注册页面
func RegisterPage(ctx *gin.Context) {
	logger.Success(ctx, "欢迎来到注册页面")
}

// Login 用户登录
func Login(ctx *gin.Context) {
	// 获取用户登录账号和密码
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		logger.Fail(ctx, 400, "请输入账号和密码")
		return
	}

	// 如果redis或mysql中存在
	if user.Select(ctx) {
		if tokenString, ok := middleware.IssueToken(user.Name); ok {
			ctx.Header("token", tokenString)
			logger.Success(ctx, "登录成功")
			logger.Redirect(ctx, "/")
		} else {
			logger.Fail(ctx, 500, "token生成错误")
		}
		return
	}
	logger.Fail(ctx, 403, "用户或者密码错误!")
}

// LoginPage 用户登录页面
func LoginPage(ctx *gin.Context) {
	logger.Success(ctx, "欢迎来到登录页面")
}

// ProductPage 单个产品秒杀页面
func ProductPage(ctx *gin.Context) {

}

// SecKill 秒杀逻辑
func SecKill(ctx *gin.Context) {

}
