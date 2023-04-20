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
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	user := model.User{
		Name:     name,
		Password: password,
	}

	// 存入mysql和redis
	if user.Create(ctx) {
		logger.Fail(ctx, 403, "重复注册")
		return
	}
}

// RegisterPage 用户注册页面
func RegisterPage(ctx *gin.Context) {
	logger.Success(ctx, "注册成功")
}

// Login 用户登录
func Login(ctx *gin.Context) {
	// 获取用户登录账号和密码
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	user := model.User{
		Name:     name,
		Password: password,
	}

	// 如果redis和mysql中存在
	if user.Select(ctx) {
		middleware.IssueToken(ctx, user.Name)
		return
	}

	logger.Fail(ctx, 403, "用户不存在, 请先登录")
	logger.Redirect(ctx, "/register")
}

// LoginPage 用户登录页面
func LoginPage(ctx *gin.Context) {
	logger.Success(ctx, "登录成功")
}

// ProductPage 单个产品秒杀页面
func ProductPage(ctx *gin.Context) {

}

// SecKill 秒杀逻辑
func SecKill(ctx *gin.Context) {

}
