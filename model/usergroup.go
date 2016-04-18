package model

import ()

type Usergroup struct {
	Id    int    `json:"id" gorm:"primary_key;AUTO_INCREMENT;" validate:"omitempty,min=1"` //主键id
	Name  string `json:"name" sql:"type:varchar(255);default:`                             //用户组名称
	Des   string `json:"des" sql:"type:varchar(1000);default:`                             //用户组描述
	State int    `json:"state" sql:"default:0" validate:"omitempty,min=1"`                 //用户组状态
}

func (u Usergroup) TableName() string {
	return "pz_userGroup"
}

/**
 * 获取usergroup
 * @method UsergroupGet
 * @param  {[type]} id int [description]
 */
func UsergroupGet(id int) ApiJson {
	var usergroup Usergroup
	err := DB.First(&usergroup, id).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	}
	return ApiJson{State: true, Msg: usergroup}
}
/**
 * 获取用户组信息通过组名称
 * @method UserGroupGetByName
 * @param  {[type]}           name string [description]
 */
func UserGroupGetByName(name string) (Usergroup, error) {
	var usergroup Usergroup
	err := DB.Where("name = ?", name).First(&usergroup).Error
	return usergroup, err
}

/**
 * 更新usergroup
 * @method UsergroupUpdate
 * @param  {[type]}    usergroup Usergroup [description]
 */
func UsergroupUpdate(usergroup Usergroup) ApiJson {
	err := DB.Model(&usergroup).UpdateColumns(map[string]interface{}{"name": usergroup.Name, "des": usergroup.Des, "state": usergroup.State}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 创建usergroup
 * @method UsergroupCreate
 * @param  {[type]}    usergroup Usergroup [description]
 */
func UsergroupCreate(usergroup Usergroup) ApiJson {
	DB.Save(&usergroup)
	return ApiJson{State: true, Msg: usergroup.Id}
}

/**
 * page usergroup
 * @method UsergroupPage
 * @param  {[type]}  kw string [description]
 * @param  {[type]}  cp int    [description]
 * @param  {[type]}  mp int    [description]
 */
func UsergroupPage(kw string, cp int, mp int) ApiJson {
	var usergroups []Usergroup
	var count int
	DB.Table("pz_userGroup").Select("*").Where("name like ?", "%"+kw+"%").Count(&count).Offset((cp - 1) * mp).Limit(mp).Find(&usergroups)
	return ApiJson{State: true, Msg: usergroups, Count: count}
}

/**
 * 删除用户组
 * @method UsergroupDele
 * @param  {[type]} ids int[] [description]
 */
func UsergroupDel(id int) ApiJson {
	err := DB.Where("id = ?", id).Delete(Usergroup{}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
