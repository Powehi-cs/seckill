package test

import (
	"github.com/Powehi-cs/seckill/internal/config"
	"github.com/Powehi-cs/seckill/internal/model"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	config.ReadConfig()
	database.MySQLConnect()
	db := database.GetDataBase()
	err := db.AutoMigrate(&model.User{}, &model.Product{}) // 绑定模型
	errors.PrintInStdout(err)
	user := model.User{
		Name:     "mofan",
		Password: "mysonmyson",
	}
	err = user.Create()
	errors.PrintInStdout(err)
	//user.Password = "mofanmofanmofanmofan"
	//err = user.Update()
	////errors.PrintInStdout(err)
	//
	//err = user.Delete()
	//errors.PrintInStdout(err)
}

func TestProduct(t *testing.T) {
	config.ReadConfig()
	database.MySQLConnect()
	db := database.GetDataBase()
	err := db.AutoMigrate(&model.User{}, &model.Product{}) // 绑定模型
	errors.PrintInStdout(err)

	product := model.Product{
		ProductID: 1,
		Name:      "书本",
		Price:     128,
		Inventory: 100,
	}

	//err = product.Create()

	product.Name = "哈利波特"
	err = product.Update()

	product2 := model.Product{
		ProductID: 1,
	}

	err = product2.Select()
	log.Println(product2.Name, product2.Inventory)

	product.Delete()
}
