package model

import (
	"encoding/json"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:string;size:256;not null;uniqueIndex" json:"name" form:"name" binding:"required"`
	Password string `gorm:"type:string;size:256;not null;" json:"password" form:"password" binding:"required"`
}

// Create 创建一个用户
func (u *User) Create(ctx *gin.Context) bool {
	rdb := database.GetRedis()
	errRdb := rdb.Set(ctx, u.Name, u, 30*time.Minute).Err() // 30分钟过期

	db := database.GetDataBase()
	errDb := db.Create(u).Error

	return errDb == nil && errRdb == nil
}

// Delete 删除掉对应name值的用户
func (u *User) Delete(ctx *gin.Context) bool {
	rdb := database.GetRedis()
	errRdb := rdb.Del(ctx, u.Name).Err()

	db := database.GetDataBase()
	errDb := db.Where("name = ?", u.Name).Delete(u).Error

	return errDb == nil && errRdb == nil
}

// Update 根据name更新用户密码(延时双删:缓存清除——数据库更新——延时2s——缓存清除)
func (u *User) Update(ctx *gin.Context) bool {
	rdb := database.GetRedis()
	rdb.Del(ctx, u.Name)

	db := database.GetDataBase()
	err := db.Model(u).Where("name = ?", u.Name).Update("password", u.Password).Error
	if err != nil {
		return false
	}

	time.Sleep(2 * time.Second)

	rdb.Del(ctx, u.Name)

	return true
}

func checkPassword(userPassword, realPassword string) bool {
	if userPassword != realPassword {
		return false
	}
	return true
}

// Select 根据name查找用户
func (u *User) Select(ctx *gin.Context) bool {
	rdb := database.GetRedis()
	user := new(User)
	errRdb := rdb.Get(ctx, u.Name).Scan(user)
	if errRdb != nil { // redis中不存在该账号
		db := database.GetDataBase()
		errDb := db.Where("name = ?", u.Name).First(user).Error
		if errDb != nil { // 数据库中不存在该账号
			return false
		}
		// 数据库中存在该账号，核对密码
		if checkPassword(u.Password, user.Password) {
			errRdb = rdb.Set(ctx, u.Name, u, 30*time.Minute).Err() // 将数据库中记录更新到缓存中
			return true
		}
		// 密码错误
		return false
	}

	// redis中存在该账号，核对密码
	return checkPassword(u.Password, user.Password)
}

func (u *User) MarshalBinary() ([]byte, error) {
	data, err := json.Marshal(u)
	return data, err
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
