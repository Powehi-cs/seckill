package main

import (
	"github.com/Powehi-cs/seckill/internal/api"
	"github.com/Powehi-cs/seckill/internal/config"
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/Powehi-cs/seckill/pkg/rabbitMQ"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	api.Router(router)      // 设置路由
	config.ReadConfig()     // 读取配置文件
	database.MySQLConnect() // 初始化mysql

	db := database.GetDataBase()
	err := db.AutoMigrate(&model.User{}, &model.Product{}) // 绑定模型
	errors.PrintInStdout(err)

	database.RedisConnect() // 初始化redis

	// 初始化生产者
	rabbitMQ.NewMQ("seckill")

	err = router.Run(viper.GetString("server.ip") + ":" + viper.GetString("server.port"))
	errors.PrintInStdout(err)
}
