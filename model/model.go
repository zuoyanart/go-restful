package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:huabinglan@@227@(192.168.1.117:3306)/pizzaCms??charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}
	DB.DB().SetMaxIdleConns(10) //最大连接数
	DB.DB().SetMaxOpenConns(100)
	DB.SingularTable(true) //// Disable table name's pluralization
	DB.LogMode(true)

	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Article{}, &Node{})
}

/**
 * 用于api输出信息的个耍
 */
type ApiJson struct {
	State bool        `json:"state"`
	Msg   interface{} `json:"msg"`
	Count int         `json:count`
}
