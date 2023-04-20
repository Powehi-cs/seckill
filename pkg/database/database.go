package database

import (
	"fmt"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func MySQLConnect() {
	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errors.PrintInStdout(err)
	database = db
}

func GetDataBase() *gorm.DB {
	return database
}

func getDSN() string {
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	ip := viper.GetString("mysql.ip")
	port := viper.GetString("mysql.port")
	dbName := viper.GetString("mysql.db_name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbName)
	return dsn
}
