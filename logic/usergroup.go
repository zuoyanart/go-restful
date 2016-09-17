package logic

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"rmApi/model"
)

/**
 * 获取一条usergroup
 * @method UserGroupGet
 * @param  {[type]} id string [description]
 */
func UserGroupGet(id string) model.ApiJson {
	var usergroup model.UserGroup
	usergroup, err := model.UserGroupGet(id)
	if err != nil {
		if err == mgo.ErrNotFound { //数据不存在，则需要创建
			group := []model.UserGroupInfo{} //固定分组信息
			group = append(group, model.UserGroupInfo{
				Name:  "商务合作",
				State: 0,
				Count: 0,
			})
			group = append(group, model.UserGroupInfo{
				Name:  "好友",
				State: 0,
				Count: 0,
			})
			group = append(group, model.UserGroupInfo{
				Name:  "同学",
				State: 0,
				Count: 0,
			})
			group = append(group, model.UserGroupInfo{
				Name:  "家族",
				State: 0,
				Count: 0,
			})
			group = append(group, model.UserGroupInfo{
				Name:  "同事",
				State: 0,
				Count: 0,
			})
			group = append(group, model.UserGroupInfo{
				Name:  "重要人物",
				State: 0,
				Count: 0,
			})
			group = append(group, model.UserGroupInfo{
				Name:  "陌生人",
				State: 0,
				Count: 0,
			})

			usergroup = model.UserGroup{
				Id:    bson.ObjectIdHex(id),
				Group: group,
			}
			model.UserGroupCreate(usergroup)
      return model.ApiJson{State: true, Msg: usergroup}
		} else {
			return model.ApiJson{State: false, Msg: err.Error()}
		}
	} else {
		return model.ApiJson{State: true, Msg: usergroup}
	}
}
