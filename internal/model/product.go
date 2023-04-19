package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	name      string // 商品名称
	price     string // 商品价格
	inventory int    // 商品库存
}
