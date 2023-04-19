package model

import (
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	nickName string // 用户名称
	password string // 用户密码
}

func (u *User) Insert(nickName, password string) {
	db := database.GetDataBase()
	user := &User{nickName: nickName, password: password}
	result := db.Create(user)
	errors.PrintInStdout(result.Error)
}

func (u *User) Update(password string) {
	db := database.GetDataBase()
	db.Model(u).Update("password", password)
}

func (u *User) Select(nickName string) {
	db := database.GetDataBase()
	db.Where("nickName = ?", nickName).First(u)
}

func (u *User) Delete() {
	db := database.GetDataBase()
	db.Where("nickName = ?", u.nickName).Delete(&User{})
}
