package model

type UserAdmin struct {
	ID       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT" `
	Username string `json:"username" sql:"type:varchar(30);default:''" validate:"required,max=30,min=4"`
	Nickname string `json:"nickname" sql:"type:varchar(30);default:''" validate:"required,max=30,min=2"`
	Password string `json:"password" sql:"size:100;default:''" validate:"omitempty,max=25,min=6"`
	State    int    `json:"state" sql:"default:0" validate:"gte=-1,lte=3"`
	Salt     string `json:"salt"`
}

// func init() {
// 	//自动更新表结构，注意：只更新新增的字段和索引
// 	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
// }

func (u UserAdmin) TableName() string {
	return "pz_user"
}

/**
 * 根据user id获取 user
 * @method UserAdminGet
 * @param  {[type]} id int [description]
 */
func UserAdminGet(id int) ApiJson {
	var user UserAdmin
	DB.Select("id,username,nickname,state").First(&user, id)
	return ApiJson{State: true, Msg: user}
}

/**
 * 校验用户登录
 * @method UserAdminCheckLogin
 * @param  {[type]}       username string [description]
 * @param  {[type]}       password string [description]
 */
func UserAdminCheckLogin(username string) UserAdmin {
	var user UserAdmin
	DB.Where("username =  ?", username).First(&user)
	return user
}

/**
 * 更新user信息
 * @method UserAdminUpdate
 * @param  {[type]}   user UserAdmin [description]
 */
func UserAdminUpdate(user UserAdmin) ApiJson {
	var err error
	if user.Password == "" {
		err = DB.Model(&user).UpdateColumns(map[string]interface{}{"Username": user.Username, "nickname": user.Nickname}).Error
	} else {
		var salt = Tools.GetRandomString(10)
		err = DB.Model(&user).UpdateColumns(map[string]interface{}{"Username": user.Username, "Nickname": user.Nickname, "Password": Tools.MD5(user.Password + salt), "salt": salt}).Error
	}
	if err != nil {
		return ApiJson{State: false, Msg: err}
	}
	return ApiJson{State: true}
}

/**
 * 创建user
 * @method UserAdminCreate
 * @param  {[type]}   user UserAdmin [description]
 */
func UserAdminCreate(user UserAdmin) ApiJson {
	var salt = Tools.GetRandomString(10)
	user.Password = Tools.MD5(user.Password + salt)
	user.Salt = salt
	DB.Save(&user)
	return ApiJson{State: true, Msg: user.ID}
}

/**
 * 获取所有的user、
 * @method UserAdminPage
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func UserAdminPage(kw string, cp int, mp int) ApiJson {
	var users []UserAdmin
	var count int
	DB.Table("pz_user").Select("id, username, state").Where("username like ?", "%"+kw+"%").Count(&count).Offset((cp - 1) * mp).Limit(mp).Find(&users)
	return ApiJson{State: true, Msg: users, Count: count}
}

/**
 * 删除用户
 * @method UserAdminDele
 * @param  {[type]} ids int[] [description]
 */
func UserAdminDele(ids []int) ApiJson {
	err := DB.Where("id in (?) ", ids).Delete(UserAdmin{}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
