package main

import (
	"github.com/Powehi-cs/seckill/internal/config"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/rabbitMQ"
)

func main() {
	config.ReadConfig()     // 读取配置文件
	database.MySQLConnect() // 初始化mysql

	rabbitMQ.NewMQ("seckill") // 初始化MQ

	mq := rabbitMQ.GetRabbitMQ()
	mq.ConsumeSimple()
}
