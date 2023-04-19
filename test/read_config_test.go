package test

import (
	"github.com/Powehi-cs/seckill/internal/config"
	"testing"
)

func TestReadConfig(t *testing.T) {
	config.ReadConfig() // 要将getPath()函数改成 getwd + "/../configs"
}
