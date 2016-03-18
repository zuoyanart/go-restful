## useage

model.go

```js
package model
import (
 "pizzaCmsApi/mongodb"
)

var (
	Modb *mongodb.Mongodb //mongodb
)

func init() {
Modb = mongodb.New(ConnectStr)
}

```

need.go

```js
package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Need struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Uid     bson.ObjectId `json:"uid"`
	Nid     int32         `json:"nid"`
	Title   string        `json:"title"`
	Tag     string        `json:"tag"`
	Des     string        `json:"des"`
	Cat     int32         `json:"cat"`
	Imgs    string        `json:"imgs"`
	Endtime int64         `json:"endtime"` //最后更新时间
	Cmt     int32         `json:"cmt"`     //评论总数
	Prov    string        `json:"prov"`
	City    string        `json:"city"`
	State   int32         `json:"state"` //分享状态，0为未分享，1为有偿分享,2为暂时不能分享，3为关闭
}

/**
 * 创建need
 * @method NeedCreate
 * @param  {[type]}   need Need [description]
 */
func NeedCreate(need Need) ApiJson {
	session, c := Modb.SwitchC("need")
	defer session.Close()
	need.Id = bson.NewObjectId()
	err := c.Insert(need)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: need.Id.Hex()}
	}
}

/**
 * 获取一条need
 * @method NeedGet
 * @param  {[type]} id string [description]
 */
func NeedGet(id string) ApiJson {
	var need Need
	objid := bson.ObjectIdHex(id)
	session, c := Modb.SwitchC("need")
	defer session.Close()
	err := c.FindId(objid).One(&need)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: need}
	}
}

/**
 * 获取默认所有的供需
 * @method NeedPage
 * @param  {[type]} id string [description]
 */
func NeedPage(uid string, skip int, mp int) ApiJson {
	var needs []Need
	session, c := Modb.SwitchC("need")
	defer session.Close()
	err := c.Find(bson.M{"uid": bson.ObjectIdHex(uid)}).Skip(skip).Sort("-_id").Limit(mp).All(&needs) //降序
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: needs}
	}
}

//删除一条需求
func NeedDel(id string) ApiJson {
	session, c := Modb.SwitchC("need")
	defer session.Close()
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id), "state": 0})
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 更新need数据
 * @method UpdateNeed
 * @param  {[type]}   query  bson.M [description]
 * @param  {[type]}   change bson.M [description]
 */
func NeedUpdate(query bson.M, change bson.M) ApiJson {
	session, c := Modb.SwitchC("need")
	defer session.Close()
	err := c.Update(query, change)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}


```