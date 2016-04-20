package model

import ()

type Actionmenu struct {
	Id    int    `json:"id" gorm:"primary_key;AUTO_INCREMENT" validate:"omitempty,min=1"` //主键id
	Name  string `json:"name" validate:"omitempty,min=1,max=510"`                         //栏目模块名称
	Link  string `json:"link" validate:"omitempty,min=1,max=510"`                         //栏目模块编码
	State int    `json:"state" validate:"omitempty,min=1"`                                //栏目模块状态
	Pid   int    `json:"pid" validate:"omitempty,min=1"`                                  //父节点
	Icon  string `json:"icon" validate:"omitempty,min=1,max=100"`                         //图标
}

func (u Actionmenu) TableName() string {
	return "pz_actionMenu"
}

/**
 * 获取actionmenu
 * @method ActionmenuGet
 * @param  {[type]} id int [description]
 */
func ActionmenuGet(id int) ApiJson {
	var actionmenu Actionmenu
	err := DB.First(&actionmenu, id).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	}
	return ApiJson{State: true, Msg: actionmenu}
}

/**
 * 更新actionmenu
 * @method ActionmenuUpdate
 * @param  {[type]}    actionmenu Actionmenu [description]
 */
func ActionmenuUpdate(actionmenu Actionmenu) ApiJson {
	err := DB.Model(&actionmenu).UpdateColumns(map[string]interface{}{"name": actionmenu.Name, "link": actionmenu.Link, "state": actionmenu.State, "icon": actionmenu.Icon}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 创建actionmenu
 * @method ActionmenuCreate
 * @param  {[type]}    actionmenu Actionmenu [description]
 */
func ActionmenuCreate(actionmenu Actionmenu) ApiJson {
	DB.Save(&actionmenu)
	return ApiJson{State: true, Msg: actionmenu.Id}
}

/**
 * page actionmenu
 * @method ActionmenuPage
 * @param  {[type]}  kw string [description]
 * @param  {[type]}  cp int    [description]
 * @param  {[type]}  mp int    [description]
 */
func ActionmenuPage(pid int, cp int, mp int) ApiJson {
	var actionmenus []Actionmenu
	var count int
	DB.Table("pz_actionMenu").Select("*").Where("pid = ?", pid).Offset((cp - 1) * mp).Limit(mp).Find(&actionmenus)
	return ApiJson{State: true, Msg: actionmenus, Count: count}
}

/**
 * 删除部门栏目
 * @method ActionmenuDele
 * @param  {[type]} ids int[] [description]
 */
func ActionmenuDel(id int, uid int) ApiJson {
	err := DB.Where("id = ? and uid = ?", id, uid).Delete(Actionmenu{}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
