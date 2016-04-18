package logic

import (
	"pizzaCmsApi/model"
)

/**
 * 判断用户是否登录
 * @method UserLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func UsergroupCreate(usergroup model.Usergroup) model.ApiJson {
	_, err := model.UserGroupGetByName(usergroup.Name)
	if err != nil && err.Error() == "record not found" { //用户组不存在
		return model.UsergroupCreate(usergroup)
	} else {
		return model.ApiJson{State: false, Msg: "group name is exist"}
	}
}

/**
 * 判断用户是否登录
 * @method UserLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func UsergroupUpdate(usergroup model.Usergroup) model.ApiJson {
	_, err := model.UserGroupGetByName(usergroup.Name)
	if err != nil && err.Error() == "record not found" { //用户组不存在
		return model.UsergroupUpdate(usergroup)
	} else {
		return model.ApiJson{State: false, Msg: "group name is exist"}
	}
}
