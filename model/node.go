package model

import ()

type Node struct {
	ID       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Pid      int    `json:"pid" sql:"default:0" validate:"required,max=1000,min=1"`
	Name     string `json:"name" sql:"type:varchar(50);default:''" validate:"required,max=40,min=1"`
	Brief    string `json:"brief" sql:"type:varchar(255);default:''"`
	Nodepath string `json:"nodepath" sql:"type:varchar(255);default:''" validate:"required,max=100,min=4"`
	Link     string `json:"link" sql:"type:varchar(100);default:''" validate:"required,max=100,min=4"`
	Weight   int    `json:"weight" sql:"default:0" validate:"required,numeric,max=3000,min=0"`
}

// func init() {
// 	//自动更新表结构，注意：只更新新增的字段和索引
// 	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Node{})
// }

func (u Node) TableName() string {
	return "pz_node"
}

/**
 * 根据node id获取 node
 * @method NodeGet
 * @param  {[type]} id int [description]
 */
func NodeGet(id int) Node {
	var node Node
	DB.First(&node, id)
	return node
}

/**
 * 更新node信息
 * @method NodeUpdate
 * @param  {[type]}   node Node [description]
 */
func NodeUpdate(node Node) ApiJson {
	err := DB.Model(&node).UpdateColumns(Node{Name: node.Name, Brief: node.Brief, Link: node.Link, Weight: node.Weight}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	}
	return ApiJson{State: true}
}

/**
 * 更新node信息
 * @method NodeUpdate
 * @param  {[type]}   node Node [description]
 */
func NodeUpdateNodepath(node Node) ApiJson {
	err := DB.Model(&node).UpdateColumns(Node{Nodepath: node.Nodepath}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	}
	return ApiJson{State: true}
}

/**
 * 创建node
 * @method NodeCreate
 * @param  {[type]}   node Node [description]
 */
func NodeCreate(node Node) int {
	DB.Save(&node)
	return node.ID
}

/**
 * 获取所有的node、
 * @method NodePage
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func NodePage(pid int) ApiJson {
	var nodes []Node
	DB.Where("pid = ?", pid).Find(&nodes)
	return ApiJson{State: true, Msg: nodes}
}
