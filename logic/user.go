package logic

import (
	"pizzaCmsApi/model"
  "pizzaCmsApi/tools"
  "log"
)

/**
 * 判断用户是否登录
 * @method UserLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
func UserCheckLogin(username string, password string) (model.ApiJson) {
    pwd := tools.SHA1(tools.MD5(password))
    user := model.UserCheckLogin(username, pwd);
  	log.Printf("password=" + pwd)
    if user.ID == 0 {
      return model.ApiJson{State:false, Msg: "user is no exist"}
    } else {
      user.Password = ""
      return model.ApiJson{State: true, Msg: user}
    }
}
