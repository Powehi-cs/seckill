package handler

import (
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/Powehi-cs/seckill/pkg/rabbitMQ"
	"github.com/Powehi-cs/seckill/pkg/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Register 用户注册
func Register(ctx *gin.Context) {
	// 获取用户登录账号和密码
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		ctx.JSON(200, utils.GetGinH(utils.RegisterFail, "请输入账号和密码"))
		return
	}

	// 存入mysql和redis
	if !user.Create(ctx) {
		ctx.JSON(200, utils.GetGinH(utils.RegisterFail, "重复注册"))
		return
	}

	ctx.Header("Location", "/login")
	ctx.JSON(200, utils.GetGinH(utils.RegisterSuccess, "注册成功"))
}

// RegisterPage 用户注册页面
func RegisterPage(ctx *gin.Context) {
	ctx.JSON(200, utils.GetGinH(utils.Success, "欢迎来到注册页面"))
}

// Login 用户登录
func Login(ctx *gin.Context) {
	// 获取用户登录账号和密码
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		ctx.JSON(200, utils.GetGinH(utils.LoginFail, "请输入账号和密码"))
		return
	}

	// 如果redis或mysql中存在
	if user.Select(ctx) {
		if tokenString, ok := utils.IssueToken(user.Name); ok {
			ctx.Header("Location", "/")
			ctx.Header("Token", tokenString)
			ctx.Set("name", user.Name)
			ctx.JSON(200, utils.GetGinH(utils.LoginSuccess, "登录成功"))
		} else {
			ctx.JSON(200, utils.GetGinH(utils.TokenFail, "Token生成失败"))
		}
		return
	}
	ctx.JSON(200, utils.GetGinH(utils.LoginFail, "用户名或者密码错误"))
}

// LoginPage 用户登录页面
func LoginPage(ctx *gin.Context) {
	if utils.Check(ctx) {
		ctx.Header("Location", "/")
		ctx.JSON(200, utils.GetGinH(utils.LoginSuccess, "用户登录成功"))
		return
	}
	ctx.JSON(200, utils.GetGinH(utils.LoginFail, "欢迎来到登录页面"))
}

// ProductPage 单个产品秒杀页面
func ProductPage(ctx *gin.Context) {
	ctx.JSON(200, utils.GetGinH(utils.Success, "这是一个产品"))
}

// SecKill 秒杀逻辑
func SecKill(ctx *gin.Context) {
	// 1、查看是否在黑名单中
	//bl := utils.GetBlackList()
	//name, ok := ctx.Get("name")
	//if !ok || bl.Get(name.(string)) {
	//	ctx.JSON(200, utils.GetGinH(utils.OrderFail, "下单失败"))
	//	return
	//}

	// 2、通过redis lua脚本预扣减库存
	if uid, ok := purchase(ctx); ok {
		unLock(ctx, uid)
		//bl.Add(name.(string))
		// 异步的将订单送入rabbitMQ并让mysql去消费订单
		mq := rabbitMQ.GetRabbitMQ()
		mq.PublishSimple(ctx.Param("product_id"))
		ctx.JSON(200, utils.GetGinH(utils.OrderSuccess, "下单成功"))
		return
	}

	// 失败处理
	ctx.JSON(200, utils.GetGinH(utils.OrderFail, "下单失败"))
}

// 释放锁
func unLock(ctx *gin.Context, uid uuid.UUID) {
	_, unlock := utils.GetPairLock()
	rdb := database.GetRedis()
	productID := ctx.Param("product_id")

	_, err := rdb.Eval(ctx, unlock, []string{productID}, uid).Result()
	errors.PrintInStdout(err)
}

// 加锁并对库存做预扣减
func purchase(ctx *gin.Context) (uuid.UUID, bool) {
	lock, _ := utils.GetPairLock()
	rdb := database.GetRedis()
	productID := ctx.Param("product_id")
	uid := uuid.NewV4()

	res, err := rdb.Eval(ctx, lock, []string{productID}, uid).Int()
	errors.PrintInStdout(err)

	for res == 0 {
		res, _ = rdb.Eval(ctx, lock, []string{productID}, uid).Int()
	}

	if res == 1 {
		return uid, true
	} else {
		return uid, false
	}
}
