package logic

import (
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
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
		if pwd == password {
			user.Password = ""
			return model.ApiJson{State: true, Msg: user}
		} else {
			return model.ApiJson{State: false, Msg: "username or password is error"}
		}
	}
}
