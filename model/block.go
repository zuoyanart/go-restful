package model

import ()

type Block struct {
	Id      int    `json:"id" gorm:"primary_key;AUTO_INCREMENT" validate:"omitempty,min=1"` //主键id
	Title   string `json:"title" sql:"type:varchar(100);default:''"`                           //标题
	Content string `json:"content" validate:"omitempty,min=1,max=20000"`                   //描述
}

func (u Block) TableName() string {
	return "pz_block"
}

/**
 * 获取block
 * @method BlockGet
 * @param  {[type]} id int [description]
 */
func BlockGet(id int) ApiJson {
	var block Block
	err := DB.First(&block, id).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	}
	return ApiJson{State: true, Msg: block}
}

/**
 * 更新block
 * @method BlockUpdate
 * @param  {[type]}    block Block [description]
 */
func BlockUpdate(block Block) ApiJson {
	err := DB.Model(&block).UpdateColumns(map[string]interface{}{"title": block.Title, "content": block.Content}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	} else {
		return ApiJson{State: true}
	}
}

/**
 * 创建block
 * @method BlockCreate
 * @param  {[type]}    block Block [description]
 */
func BlockCreate(block Block) ApiJson {
	DB.Save(&block)
	return ApiJson{State: true, Msg: block.Id}
}

/**
 * page block
 * @method BlockPage
 * @param  {[type]}  kw string [description]
 * @param  {[type]}  cp int    [description]
 * @param  {[type]}  mp int    [description]
 */
func BlockPage(kw string, cp int, mp int) ApiJson {
	var blocks []Block
	var count int
	DB.Table("pz_block").Select("*").Where("title like ?", "%"+kw+"%").Count(&count).Offset((cp - 1) * mp).Limit(mp).Find(&blocks)
	return ApiJson{State: true, Msg: blocks, Count: count}
}

/**
 * 删除模块
 * @method BlockDele
 * @param  {[type]} ids int[] [description]
 */
func BlockDel(id []int) ApiJson {
	err := DB.Where("id in (?)", id).Delete(Block{}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err.Error()}
	} else {
		return ApiJson{State: true}
	}
}
