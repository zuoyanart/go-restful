package logic

import (
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
)

/**
 * 创建node
 * @method NodeCreate
 * @param  {[type]}   node model.Node [description]
 */
func NodeCreate(node model.Node) model.ApiJson {
	nodeParent := model.NodeGet(node.Pid)
	if nodeParent.Nodepath != "" { //获取到数据
		node.Nodepath = nodeParent.Nodepath
		id := model.NodeCreate(node)
    node.ID = id;
    node.Nodepath = node.Nodepath + tools.ParseString(id) + ","
    return model.NodeUpdateNodepath(node);
	} else {
    	return model.ApiJson{State: false, Msg: "parent node not exist"}
  }
}

/**
 * 根据node id获取 node
 * @method NodeGet
 * @param  {[type]} id int [description]
 */
func NodeGet(id int) model.ApiJson {
	node := model.NodeGet(id)
	return model.ApiJson{State: true, Msg: node}
}
