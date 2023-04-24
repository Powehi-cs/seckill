package database

import (
	"github.com/Powehi-cs/seckill/pkg/utils"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func RedisConnect() {
	rdb = redis.NewClient(getOptions())
	utils.InitLua()
}

func GetRedis() *redis.Client {
	return rdb
}

func getOptions() *redis.Options {
	ip := viper.GetString("redis.ip")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.Db")
	options := &redis.Options{
		Addr:     ip + ":" + port,
		Password: password,
		DB:       db,
	}
	return options
}
