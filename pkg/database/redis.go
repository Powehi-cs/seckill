package database

import (
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strconv"
)

var rdb *redis.Client

func RedisConnect() {
	rdb = redis.NewClient(getOptions())
}

func GetRedis() *redis.Client {
	return rdb
}

func getOptions() *redis.Options {
	sql := viper.Get("redis").(map[string]string)
	ip := sql["ip"]
	port := sql["port"]
	password := sql["password"]
	db, err := strconv.Atoi(sql["DB"])
	errors.PrintInStdout(err)
	options := &redis.Options{
		Addr:     ip + ":" + port,
		Password: password,
		DB:       db,
	}
	return options
}
