package logic

import (
	"strings"

	"pizzaCmsApi/model"
)

/**
 * 判断用户是否登录
 * @method UserAdminLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func UserAdminCheckLogin(username string, password string) model.ApiJson {
	user := model.UserAdminCheckLogin(username)
	if user.ID == 0 {
		return model.ApiJson{State: false, Msg: "user is no exist"}
	} else {
		pwd := Tools.MD5(password + user.Salt)
		if pwd == user.Password {
			user.Password = ""
			return model.ApiJson{State: true, Msg: user}
		} else {
			return model.ApiJson{State: false, Msg: "username or password is error"}
		}
	}
}

/**
 * 删除用户
 * @method UserAdminDele
 * @param  {[type]} ids string [description]
 */
func UserAdminDele(ids string) model.ApiJson {
	idsArr := strings.Split(ids, ",")
	length := len(idsArr)
	if length > 0 {
		var idsInt = make([]int, length, length)
		for i, id := range idsArr {
			idsInt[i] = Tools.ParseInt(id, 0)
		}
		return model.UserAdminDele(idsInt)
	} else {
		return model.ApiJson{State: false, Msg: "id is error"}
	}
}
