package handler

import (
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/internal/utils"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	// 获取用户登录账号和密码
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		ctx.JSON(400, "请输入账号和密码")
		return
	}

	// 存入mysql和redis
	if !user.Create(ctx) {
		ctx.JSON(403, "重复注册")
		return
	}

	ctx.JSON(200, "注册成功")
}

// RegisterPage 用户注册页面
func RegisterPage(ctx *gin.Context) {
	ctx.JSON(200, "欢迎来到注册页面")
}

// Login 用户登录
func Login(ctx *gin.Context) {
	// 获取用户登录账号和密码
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		ctx.JSON(400, "请输入账号和密码")
		return
	}

	// 如果redis或mysql中存在
	if user.Select(ctx) {
		if tokenString, ok := utils.IssueToken(user.Name); ok {
			ctx.JSON(302, gin.H{"path": "/", "token": tokenString})
		} else {
			ctx.JSON(500, "token生成错误")
		}
		return
	}
	ctx.JSON(403, "用户或者密码错误！")
}

// LoginPage 用户登录页面
func LoginPage(ctx *gin.Context) {
	if utils.Check(ctx) {
		ctx.JSON(302, "/")
		return
	}

	ctx.JSON(200, "欢迎来到登录界面")
}

// ProductPage 单个产品秒杀页面
func ProductPage(ctx *gin.Context) {
	ctx.JSON(200, "这是一个产品")
}

// SecKill 秒杀逻辑
func SecKill(ctx *gin.Context) {
	ctx.JSON(200, "秒杀成功")
}
