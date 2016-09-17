package model

import (
	"gopkg.in/mgo.v2/bson"
)

type UserGroupInfo struct {
	Name  string `json:"name"`
	State int    `json:"state"`
	Count int    `json:"count"`
}

type UserGroup struct {
	Id    bson.ObjectId   `bson:"_id" json:"id" validate:"omitempty,min=10,max=14"` //
	Group []UserGroupInfo `json:"group"`                                            //
}

/**
 * 创建usergroup
 * @method UserGroupCreate
 * @param  {[type]}   usergroup UserGroup [description]
 */
func UserGroupCreate(usergroup UserGroup) ApiJson {
	session, c := Modb.SwitchC("usergroups")
	defer session.Close()
	// usergroup.Id = bson.NewObjectId()
	err := c.Insert(usergroup)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: usergroup.Id.Hex()}
	}
}

/**
 * 获取一条usergroup
 * @method UserGroupGet
 * @param  {[type]} id string [description]
 */
func UserGroupGet(id string) (UserGroup, error) {
	var usergroup UserGroup
	objid := bson.ObjectIdHex(id)
	session, c := Modb.SwitchC("usergroups")
	defer session.Close()
	err := c.FindId(objid).One(&usergroup)
	return usergroup, err
}

/**
 * 获取用户分组列表
 * @method UserGroupPage
 * @param  {[type]} id string [description]
 */
func UserGroupPage(uid string, cp int, mp int) ApiJson {
	var usergroups []UserGroup
	session, c := Modb.SwitchC("usergroups")
	defer session.Close()
	err := c.Find(bson.M{"uid": bson.ObjectIdHex(uid)}).Skip((cp - 1) * mp).Sort("-_id").Limit(mp).All(&usergroups) //降序
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: usergroups}
	}
}

//删除一条用户分组
func UserGroupDel(id string, uid string) ApiJson {
	session, c := Modb.SwitchC("usergroups")
	defer session.Close()
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id), "uid": bson.ObjectIdHex(uid)})
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 更新usergroup数据
 * @method UpdateUserGroup
 * @param  {[type]}   query  bson.M [description]
 * @param  {[type]}   change bson.M [description]
 */
func UserGroupUpdate(query bson.M, change bson.M) ApiJson {
	session, c := Modb.SwitchC("usergroups")
	defer session.Close()
	err := c.Update(query, change)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
