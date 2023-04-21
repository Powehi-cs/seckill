package model

import (
	"github.com/Powehi-cs/seckill/pkg/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID uint   `gorm:"type:uint;uniqueIndex;not null" form:"product_id" binding:"required"`
	Name      string `gorm:"type:string;size:256;not null" form:"name" binding:"required"`
	Price     int    `gorm:"type:int;not null" form:"price" binding:"required"`
	Inventory int    `gorm:"type:int;not null" form:"inventory" binding:"required"`
}

func (p *Product) Create() error {
	db := database.GetDataBase()
	result := db.Create(p)
	return result.Error
}

func (p *Product) Delete() error {
	db := database.GetDataBase()
	result := db.Where("product_id = ?", p.ProductID).Delete(p)
	return result.Error
}

func (p *Product) Update() error {
	db := database.GetDataBase()
	result := db.Model(p).Where("product_id = ?", p.ProductID).Updates(p)
	return result.Error
}

func (p *Product) Select() error {
	db := database.GetDataBase()
	result := db.Where("product_id = ?", p.ProductID).First(p)
	return result.Error
}
