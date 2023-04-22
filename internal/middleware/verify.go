package middleware

import (
	"fmt"
	"github.com/Powehi-cs/seckill/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var Secret = []byte(viper.GetString("server.secret"))

// AuthVerify 验证用户是否合法
func AuthVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("加密算法错误: %s", token.Header["alg"])
			}
			return Secret, nil
		})

		if err != nil {
			ctx.Next()
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			logger.Redirect(ctx, "/")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

type MyCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func IssueToken(name string) (string, bool) {
	claims := MyCustomClaims{
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "powehi-yc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return tokenString, false
	}
	return tokenString, true
}
