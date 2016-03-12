package model

import (
)

type User struct {
	ID       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username" sql:"type:varchar(30);default:''" validate:"required,max=30,min=4"`
	Nickname string `json:"nickname" sql:"type:varchar(30);default:''" validate:"required,max=30,min=2"`
	Password string `json:"password" sql:"size:100;default:''" validate:"required,max=100,min=10"`
	State    int    `json:"state" sql:"default:0" validate:"gte=-1,lte=3"`
	Salt string `json:"salt"`
}

// func init() {
// 	//自动更新表结构，注意：只更新新增的字段和索引
// 	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
// }

func (u User) TableName() string {
	return "pz_user"
}

/**
 * 根据user id获取 user
 * @method UserGet
 * @param  {[type]} id int [description]
 */
func UserGet(id int) ApiJson {
	var user User
	DB.Select("id,username,state").First(&user, id)
	return ApiJson{State: true, Msg: user}
}

/**
 * 校验用户登录
 * @method UserCheckLogin
 * @param  {[type]}       username string [description]
 * @param  {[type]}       password string [description]
 */
func UserCheckLogin(username string) User {
	var user User
	DB.Where("username =  ?", username).First(&user)
	return user
}

/**
 * 更新user信息
 * @method UserUpdate
 * @param  {[type]}   user User [description]
 */
func UserUpdateName(user User) ApiJson {
	err := DB.Model(&user).UpdateColumns(User{Username: user.Username}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	}
	return ApiJson{State: true}
}

/**
 * 创建user
 * @method UserCreate
 * @param  {[type]}   user User [description]
 */
func UserCreate(user User) ApiJson {
	DB.Save(&user)
	return ApiJson{State: true, Msg: user.ID}
}

/**
 * 获取所有的user、
 * @method UserPage
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func UserPage(kw string, cp int, mp int) ApiJson {
	var users []User
	var count int
	DB.Table("pz_user").Select("id, username, state").Where("username like ?", "%"+kw+"%").Count(&count).Offset((cp - 1) * mp).Limit(mp).Find(&users)
	return ApiJson{State: true, Msg: users, Count: count}
}
