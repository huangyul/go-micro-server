package _6_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

type User struct {
	gorm.Model
	Name  string
	Cards []Card `gorm:"foreignKey:UserID"`
}

type Card struct {
	gorm.Model
	Name   string
	UserID uint
}

func Test_Gorm(t *testing.T) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
		},
	)

	// 连接数据库
	dsn := "root:root@tcp(localhost:33066)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 数据库迁移
	err = db.AutoMigrate(&User{}, &Card{})
	if err != nil {
		panic(err)
	}
	//
	//db.Create(&User{Name: "user"})
	//db.Create(&Card{
	//	Name:   "card1",
	//	UserID: 1,
	//})
	var user User
	db.Preload("Cards").First(&user)
	fmt.Println(user.Cards)
}

// 创建数据
//db.Create(&User{Name: "name"})

// 批量创建数据
//var users = []User{{Name: "name1"}, {Name: "name2"}}
//db.Create(&users)

//var user User
//var users []User
//// 根据主键排序，获取第一条
//db.First(&user)
//// 指定主键
//db.First(&user, 1)
//// 获取数据库第一条数据
//db.Take(&user)
//// 获取主键降序最后一条
//db.Last(&user)
//// 获取全部数据
//db.Find(&users)

// where 查询
//var user User
//// 拼接
//db.Where("name = ?", "name").First(&user)
//// 使用struct
//db.Where(&User{Name: "name"}).First(&user)
//// 使用map
//db.Where(map[string]any{"name": "name"}).First(&user)
//
//fmt.Println(user)

// 更新
//user := &User{
//	Name: "new name",
//}
// 如果有主键就更新，没有主键就新增
//db.Save(&user)
// 只更新特定字段
//db.Model(User{}).Where("name = ?", "test").Limit(1).Update("name", "huang")
