package main

import (
	"github.com/Powehi-cs/seckill/internal/api"
	"github.com/Powehi-cs/seckill/internal/config"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.Router(router)      // 设置路由
	config.ReadConfig()     // 读取配置文件
	database.MySQLConnect() // 初始化mysql
	database.RedisConnect() // 初始化redis
}
