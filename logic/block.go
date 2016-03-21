package logic
import(
	"pizzaCmsApi/model"
  "strings"
)
/**
 * 删除用户
 * @method UserDele
 * @param  {[type]} ids string [description]
 */
func BlockDele(ids string) model.ApiJson {
	idsArr := strings.Split(ids, ",")
	length := len(idsArr)
	if length > 0 {
		var idsInt = make([]int, length, length)
		for i, id := range idsArr {
			idsInt[i] = Tools.ParseInt(id, 0)
		}
		return model.BlockDele(idsInt)
	} else {
		return model.ApiJson{State: false, Msg: "id is error"}
	}
}
