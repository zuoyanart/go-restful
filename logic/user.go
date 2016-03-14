package logic

import (
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
	"strings"
)

/**
 * 判断用户是否登录
 * @method UserLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func UserCheckLogin(username string, password string) model.ApiJson {
	user := model.UserCheckLogin(username)
	if user.ID == 0 {
		return model.ApiJson{State: false, Msg: "user is no exist"}
	} else {
		pwd := tools.MD5(password + user.Salt)
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
 * @method UserDele
 * @param  {[type]} ids string [description]
 */
func UserDele(ids string) model.ApiJson {
	idsArr := strings.Split(ids, ",")
	length := len(idsArr)
	if length > 0 {
		var idsInt = make([]int, length, length)
		for i, id := range idsArr {
			idsInt[i] = tools.ParseInt(id, 0)
		}
		return model.UserDele(idsInt)
	} else {
		return model.ApiJson{State: false, Msg: "id is error"}
	}
}
