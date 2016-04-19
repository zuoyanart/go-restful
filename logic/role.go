package logic

import (
	"pizzaCmsApi/model"
)

/**
 * 判断角色是否存在
 * @method UserLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func RoleCreate(role model.Role) model.ApiJson {
	_, err := model.RoleGetByName(role.Name)
	if err != nil && err.Error() == "record not found" { //用户组角色不存在
		return model.RoleCreate(role)
	} else {
		return model.ApiJson{State: false, Msg: "role name is exist"}
	}
}

/**
 * 判断用户是否登录
 * @method UserLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func RoleUpdate(role model.Role) model.ApiJson {
	_, err := model.RoleGetByName(role.Name)
	if err != nil && err.Error() == "record not found" { //用户组角色不存在
		return model.RoleUpdate(role)
	} else {
		return model.ApiJson{State: false, Msg: "role name is exist"}
	}
}
