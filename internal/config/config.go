package config

import (
	"fmt"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"strings"
)

// ReadConfig 读取配置文件
func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	rootPath := GetPath()
	viper.AddConfigPath(rootPath + "/configs/")
	err := viper.ReadInConfig()
	errors.PrintInStdout(err)
}

// GetPath 获取配置文件路径
func GetPath() string {
	getwd, err := os.Getwd()
	errors.PrintInStdout(err)
	fmt.Println(path.Dir(getwd))
	osType := runtime.GOOS
	var files []string
	if osType == "windows" {
		files = strings.Split(getwd, `\`)
	} else {
		files = strings.Split(getwd, `/`)
	}
	for i := len(files) - 1; i >= 0; i-- {
		if files[i] == "seckill" {
			getwd = strings.Join(files[:i+1], "/")
			break
		}
	}
	return getwd
}
