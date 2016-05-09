package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pizzaCmsApi/config"
	"pizzaCmsApi/tools"

)
var (
	DB     *gorm.DB
	Tools  *tools.Tools
	Config *config.Config
)

func init() {
	Config := config.New()
	Tools = tools.New()

	var err error
	DB, err = gorm.Open("mysql", Config.Mysql.Connect)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}
	DB.DB().SetMaxIdleConns(Config.Mysql.MaxIdle) //最大连接数
	DB.DB().SetMaxOpenConns(Config.Mysql.MaxOpen)
	DB.SingularTable(true) //// Disable table name's pluralization
	DB.LogMode(true)
	DB.Set("gorm:table_options", "ENGINE=InnoDB")//.AutoMigrate(&User{}, &Article{}, &Node{})

}

/**
 * 用于api输出信息的个耍
 */
type ApiJson struct {
	State bool        `json:"state"`
	Msg   interface{} `json:"msg"`
	Count int         `json:"count"`
}
