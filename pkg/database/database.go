package database

import (
	"fmt"
	"github.com/Powehi-cs/seckill/internal/model"
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
	err = database.AutoMigrate(&model.User{}, &model.Product{})
	errors.PrintInStdout(err)
}

func GetDataBase() *gorm.DB {
	return database
}

func getDSN() string {
	sql := viper.Get("mysql").(map[string]string)
	user := sql["user"]
	password := sql["password"]
	ip := sql["ip"]
	port := sql["port"]
	dbName := sql["db_name"]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, dbName)
	return dsn
}
