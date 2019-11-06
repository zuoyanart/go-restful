package logic

import (
	"strings"

	"pizzaCmsApi/model"
)

/**
 * 删除文章
 * @method UserDele
 * @param  {[type]} ids string [description]
 */
func ArticleDele(ids string) model.ApiJson {
	idsArr := strings.Split(ids, ",")
	length := len(idsArr)
	if length > 0 {
		var idsInt = make([]int, length, length)
		for i, id := range idsArr {
			idsInt[i] = Tools.ParseInt(id, 0)
		}
		return model.ArticleDele(idsInt)
	} else {
		return model.ApiJson{State: false, Msg: "id is error"}
	}
}

/**
 * 审核文章
 * @method UserDele
 * @param  {[type]} ids string [description]
 */
func ArticlePass(ids string, pass int) model.ApiJson {
	idsArr := strings.Split(ids, ",")
	length := len(idsArr)
	if length > 0 {
		var idsInt = make([]int, length, length)
		for i, id := range idsArr {
			idsInt[i] = Tools.ParseInt(id, 0)
		}
		return model.ArticlePass(idsInt, pass)
	}

	return model.ApiJson{State: false, Msg: "id is error"}

}
