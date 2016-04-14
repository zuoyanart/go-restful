package model

import ()

type Comment struct {
	Id        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT" validate:"omitempty,min=1"` //主键id
	Articleid int    `json:"articleid" validate:"omitempty,min=1"`                            //文章id
	Addtime   int    `json:"addtime" validate:"omitempty,min=1"`                              //添加时间
	Content   string `json:"content" validate:"omitempty,min=1,max=2000"`                     //评论内容
	Uid       int    `json:"uid" validate:"omitempty,min=1"`                                  //用户id
	Username  string `json:"username" validate:"omitempty,min=1,max=60"`                      //用户昵称
}

func (u Comment) TableName() string {
	return "pz_comment"
}

/**
 * 获取comment
 * @method CommentGet
 * @param  {[type]} id int [description]
 */
func CommentGet(id int) ApiJson {
	var comment Comment
	err := DB.First(&comment, id).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	}
	return ApiJson{State: true, Msg: comment}
}

/**
 * 更新comment
 * @method CommentUpdate
 * @param  {[type]}    comment Comment [description]
 */
func CommentUpdate(comment Comment) ApiJson {
	err := DB.Model(&comment).UpdateColumns(map[string]interface{}{"articleid": comment.Articleid, "addtime": comment.Addtime, "content": comment.Content, "uid": comment.Uid, "username": comment.Username}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 创建comment
 * @method CommentCreate
 * @param  {[type]}    comment Comment [description]
 */
func CommentCreate(comment Comment) ApiJson {
	DB.Save(&comment)
	return ApiJson{State: true, Msg: comment.Id}
}

/**
 * page comment
 * @method CommentPage
 * @param  {[type]}  kw string [description]
 * @param  {[type]}  cp int    [description]
 * @param  {[type]}  mp int    [description]
 */
func CommentPage(cp int, mp int) ApiJson {
	var comments []Comment
	var count int
	DB.Table("pz_comment").Select("*").Count(&count).Offset((cp - 1) * mp).Limit(mp).Find(&comments)
	return ApiJson{State: true, Msg: comments, Count: count}
}

/**
 * 删除文章评论
 * @method CommentDele
 * @param  {[type]} ids int[] [description]
 */
func CommentDel(id int) ApiJson {
	err := DB.Where("id = ?", id).Delete(Comment{}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
