package middleware

import (
	"fmt"
	"github.com/Powehi-cs/seckill/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

// Verify 验证用户是否合法
func Verify(ctx *gin.Context) {
	tokenString := ctx.GetHeader("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("加密算法错误: %s", token.Header["alg"])
		}
		return []byte(viper.GetString("server.secret")), nil
	})

	if err != nil {
		logger.Fail(ctx, 400, err.Error())
		logger.Redirect(ctx, "/login")
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		logger.Success(ctx, "这是主页面")
		logger.Redirect(ctx, "/")
		return
	}
	logger.Fail(ctx, 401, "token验证未通过")
	logger.Redirect(ctx, "/login")
}

type MyCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func IssueToken(ctx *gin.Context, name string) {
	claims := MyCustomClaims{
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "powehi-yc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(viper.GetString("server.secret"))
	if err != nil {
		logger.Fail(ctx, 500, "token生成错误")
		return
	}
	ctx.Set("token", tokenString)
	logger.Success(ctx, "token生成成功")
}
