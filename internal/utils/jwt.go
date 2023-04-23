package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var Secret = []byte(viper.GetString("server.secret"))

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

func ParseToken(tokenString string) (jwt.MapClaims, bool) {
	if len(tokenString) <= 7 {
		return nil, false
	}
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("加密算法错误: %s", token.Header["alg"])
		}
		return Secret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}

func Check(ctx *gin.Context) bool {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		return false
	}

	// 如果token解析成功，则继续
	if claims, ok := ParseToken(tokenString); ok {
		ctx.Set("name", claims["name"])
		return true
	}

	return false
}
