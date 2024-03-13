package _6_gorm

import (
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
	Name string
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
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	// 创建数据
	db.Create(&User{Name: "name"})

	// 批量创建数据
	//var users = []User{{Name: "name1"}, {Name: "name2"}}
	//db.Create(&users)

	var user User
	var users []User
	// 根据主键排序，获取第一条
	db.First(&user)
	// 指定主键
	db.First(&user, 1)
	// 获取数据库第一条数据
	db.Take(&user)
	// 获取主键降序最后一条
	db.Last(&user)
	// 获取全部数据
	db.Find(&users)
}

// 4-10
