package utils

import (
	"github.com/gin-gonic/gin"
)

func GetGinH(rule int, msg string) gin.H {
	var ret gin.H
	switch rule {
	case Error:
		ret = getMsg(msg, 500)
	case Success:
		ret = getMsg(msg, 200)
	case RegisterSuccess:
		ret = getMsg(msg, 302)
	case RegisterFail:
		ret = getMsg(msg, 400)
	case LoginSuccess:
		ret = getMsg(msg, 302)
	case LoginFail:
		ret = getMsg(msg, 400)
	case OrderSuccess:
		ret = getMsg(msg, 200)
	case OrderFail:
		ret = getMsg(msg, 500)
	case PurchaseSuccess:
		ret = getMsg(msg, 200)
	case PurchaseFail:
		ret = getMsg(msg, 500)
	case ForwardSuccess:
		ret = getMsg(msg, 200)
	case ForwardFail:
		ret = getMsg(msg, 500)
	case TokenSuccess:
		ret = getMsg(msg, 200)
	case TokenFail:
		ret = getMsg(msg, 302)
	}

	return ret
}

func getMsg(msg string, code int) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
	}
}
