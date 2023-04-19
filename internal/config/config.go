package config

import (
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

// ReadConfig 读取配置文件
func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	path := getPath()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	errors.PrintInStdout(err)
}

// 获取配置文件路径
func getPath() string {
	getwd, err := os.Getwd()
	errors.PrintInStdout(err)
	return getwd + "/../configs"
}
