package test

import (
	"github.com/Powehi-cs/seckill/internal/config"
	"github.com/Powehi-cs/seckill/pkg/database"
	"testing"
)

func TestMySQLConnect(t *testing.T) {
	config.ReadConfig()
	database.MySQLConnect()
}
