package model

import (
	"github.com/Powehi-cs/seckill/pkg/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:256;not null;uniqueIndex"`
	Password string `gorm:"type:string;size:256;not null;"`
}

// Create 创建一个用户
func (u *User) Create() error {
	db := database.GetDataBase()
	result := db.Create(u)
	return result.Error
}

// Delete 删除掉对应name值的用户
func (u *User) Delete() error {
	db := database.GetDataBase()
	result := db.Where("name = ?", u.Name).Delete(u)
	return result.Error
}

// Update 根据name更新用户密码
func (u *User) Update() error {
	db := database.GetDataBase()
	result := db.Model(u).Where("name = ?", u.Name).Update("password", u.Password)
	return result.Error
}

// Select 根据name查找用户
func (u *User) Select() error {
	db := database.GetDataBase()
	result := db.Where("name = ?", u.Name).First(u)
	return result.Error
}
