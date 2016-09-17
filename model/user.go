package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id" json:"id"`
	Username     string        `bson:"userName" json:"username" validate:"required,min=1,max=20"`          //用户名
	Nickname     string        `bson:"nickName" json:"nickname" validate:"required,min=1,max=20"`          //昵称
	Password     string        `json:"password" validate:"required,min=6,max=20"`                          //密码
	Photo        string        `json:"photo" validate:"omitempty,min=1,max=100"`                           //形象照片
	Avatar       string        `json:"avatar" validate:"omitempty,min=1,max=100"`                          //头像
	State        int           `json:"state" validate:"omitempty,min=1,max=100"`                           //状态
	Time         time.Time     `json:"time" validate:"omitempty,min=1,max=50"`                             //最后活动时间
	Neoid        int           `json:"neoid" validate:"omitempty,min=0"`                                   //图数据库id
	Feel         string        `json:"feel" validate:"omitempty,min=0,max=10"`                             //励志格言，暂废
	Hope         string        `json:"hope" validate:"omitempty,min=1,max=2000"`                           //自我介绍
	Count        int           `json:"count" validate:"omitempty,min=0,max=1000"`                          //人脉总数
	Privitecount int           `bson:"priviteCount" json:"privitecount" validate:"omitempty,min=0,max=10"` //私有人脉总数
	Defcom       string        `json:"defcom" validate:"omitempty,min=0,max=10"`                           //默认公司名称
	Defschool    string        `json:"defschool" validate:"omitempty,min=0,max=10"`                        //默认学校名称
	Scmail       int           `json:"scmail" validate:"omitempty,min=0,max=10"`                           //能否查看mail
	Scphone      int           `json:"scphone" validate:"omitempty,min=0,max=10"`                          //能否查看phone
	Scqq         int           `json:"scqq" validate:"omitempty,min=0,max=10"`                             //能否查看qq
	Scwx         int           `json:"scwx" validate:"omitempty,min=0,max=10"`                             //否能查看微信
	Ask          string        `json:"ask" validate:"omitempty,min=1,max=100"`                             //找回密码问题
	Answer       string        `json:"answer" validate:"omitempty,min=1,max=100"`                          //找回密码答案
	Industry     string        `json:"industry" validate:"omitempty,min=1,max=500"`                        //所属行业
	Inddes       string        `json:"inddes" validate:"omitempty,min=1,max=2000"`                         //行业描述
}

/**
 * 创建user
 * @method UserCreate
 * @param  {[type]}   user User [description]
 */
func UserCreate(user User) ApiJson {
	session, c := Modb.SwitchC("users")
	defer session.Close()
	// user.Id = bson.NewObjectId()
	err := c.Insert(user)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: map[string]interface{}{"id": user.Id.Hex(), "neoid": user.Neoid}}
	}
}

/**
 * 获取一条user
 * @method UserGet
 * @param  {[type]} id string [description]
 */
func UserGet(id string) (User, error) {
	var user User
	objid := bson.ObjectIdHex(id)
	session, c := Modb.SwitchC("users")
	defer session.Close()
	err := c.FindId(objid).One(&user)
	user.Password = ""
	return user, err
}

/**
 * 获取用户信息
 * @method UserGetByUsername
 * @param  {[type]}          username string [description]
 */
func UserGetByUsername(username string) (User, error) {
	var user User
	session, c := Modb.SwitchC("users")
	defer session.Close()

	err := c.Find(bson.M{"userName": username}).One(&user)
	if err != nil {
		Tools.Logs(err.Error())
	}
	return user, err
}

/**
 * 获取会员列表
 * @method UserPage
 * @param  {[type]} id string [description]
 */
func UserPage(uid string, cp int, mp int) ApiJson {
	var users []User
	session, c := Modb.SwitchC("users")
	defer session.Close()
	err := c.Find(bson.M{"uid": bson.ObjectIdHex(uid)}).Skip((cp - 1) * mp).Sort("-_id").Limit(mp).All(&users) //降序
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true, Msg: users}
	}
}

//删除一条会员
func UserDel(id string, uid string) ApiJson {
	session, c := Modb.SwitchC("users")
	defer session.Close()
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id), "uid": bson.ObjectIdHex(uid)})
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 更新user数据
 * @method UpdateUser
 * @param  {[type]}   query  bson.M [description]
 * @param  {[type]}   change bson.M [description]
 */
func UserUpdate(query bson.M, change bson.M) ApiJson {
	session, c := Modb.SwitchC("users")
	defer session.Close()
	err := c.Update(query, change)
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
