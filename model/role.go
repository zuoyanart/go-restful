package model

import ()

type Role struct {
	Id      int    `json:"id" gorm:"primary_key;AUTO_INCREMENT" validate:"omitempty,min=1"` //主键id
	Groupid int    `json:"groupid" validate:"omitempty,min=1"`                              //用户组id
	Name    string `json:"name" validate:"omitempty,min=1,max=510"`                         //角色名称
	Des     string `json:"des" validate:"omitempty,min=1,max=2000"`                         //角色描述
	State   int    `json:"state" validate:"omitempty,min=1"`                                //角色状态
}

func (u Role) TableName() string {
	return "pz_role"
}

/**
 * 获取role
 * @method RoleGet
 * @param  {[type]} id int [description]
 */
func RoleGet(id int) ApiJson {
	var role Role
	err := DB.First(&role, id).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	}
	return ApiJson{State: true, Msg: role}
}
/**
 * 获取用户组角色信息通过角色名称
 * @method UserGroupGetByName
 * @param  {[type]}           name string [description]
 */
func RoleGetByName(name string) (Role, error) {
	var role Role
	err := DB.Where("name = ?", name).First(&role).Error
	return role, err
}
/**
 * 更新role
 * @method RoleUpdate
 * @param  {[type]}    role Role [description]
 */
func RoleUpdate(role Role) ApiJson {
	err := DB.Model(&role).UpdateColumns(map[string]interface{}{"name": role.Name, "des": role.Des, "state": role.State}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 创建role
 * @method RoleCreate
 * @param  {[type]}    role Role [description]
 */
func RoleCreate(role Role) ApiJson {
	DB.Save(&role)
	return ApiJson{State: true, Msg: role.Id}
}

/**
 * page role
 * @method RolePage
 * @param  {[type]}  kw string [description]
 * @param  {[type]}  cp int    [description]
 * @param  {[type]}  mp int    [description]
 */
func RolePage(groupid int, kw string, cp int, mp int) ApiJson {
	var roles []Role
	DB.Table("pz_role").Select("*").Where("name like ? and groupid = ?", "%"+kw+"%", groupid).Offset((cp - 1) * mp).Limit(mp).Find(&roles)
	return ApiJson{State: true, Msg: roles, Count: 0}
}

/**
 * 删除用户组角色
 * @method RoleDele
 * @param  {[type]} ids int[] [description]
 */
func RoleDel(id int) ApiJson {
	err := DB.Where("id = ?", id).Delete(Role{}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
