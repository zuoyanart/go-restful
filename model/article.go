package model

import (
// "github.com/jinzhu/gorm"

)

type Article struct {
	ID         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title      string `json:"title" sql:"type:varchar(50);default:''"`
	Timg       string `json:"timg" sql:"type:varchar(100);default:''"`
	Content    string `json:"content" sql:"type:varchar(10000);default:''"`
	Brief      string `json:"brief" sql:"type:varchar(255);default:''"`
	Nodeid     int    `json:"nodeid" sql:"default:0"`
	Count      int    `json:"count" sql:"default:0"`
	Reco       int    `json:"reco" sql:"default:0"`
	Createtime string `json:"createtime"`
	Uid        int    `json:"uid" sql:"default:0"`
	Pass       int    `json:"pass" sql:"default:0"`
	Source     string `json:"source" sql:"type:varchar(100);default:''"`
	Tags       string `json:"tags" sql:"type:varchar(100);default:''"`
	Link       string `json:"link" sql:"type:varchar(100);default:''"`
	Comment    int    `json:"comment" sql:"default:0"`
	State      int    `json:"state" sql:"default:0"`
}

type ArticleResults struct {
	Article
	Nodeid   int    `json:"nodeid"`
	Nodename string `json:"nodename"`
	Username string `json:"username"`
}

func (u Article) TableName() string {
	return "pz_article"
}

/**
 * 根据article id获取 article
 * @method ArticleGet
 * @param  {[type]} id int [description]
 */
func ArticleGet(id int) ApiJson {
	var article Article
	DB.First(&article, id)
	return ApiJson{State: true, Msg: article}
}

/**
 * 更新article信息
 * @method ArticleUpdate
 * @param  {[type]}   article Article [description]
 */
func ArticleUpdate(article Article) ApiJson {
	err := DB.Model(&article).UpdateColumns(Article{Title: article.Title, Timg: article.Timg, Content: article.Content, Brief: article.Brief, Nodeid: article.Nodeid, Reco: article.Reco, Pass: article.Pass, Source: article.Source, Tags: article.Tags, Link: article.Link}).Error
	if err != nil {
		return ApiJson{State: false, Msg: err}
	}
	return ApiJson{State: true}
}

/**
 * 创建article
 * @method ArticleCreate
 * @param  {[type]}   article Article [description]
 */
func ArticleCreate(article Article) ApiJson {
	DB.Save(&article)
	return ApiJson{State: true, Msg: article.ID}
}

/**
 * 获取所有的article、
 * @method ArticlePage
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func ArticlePage(kw string, nodeid int, cp int, mp int) ApiJson {
	var articles []ArticleResults
	var count int
	// var param []interface{}
	// param=append("%"+kw+"%", nodeid, mp, (cp-1)*mp)
	DB.Raw("select a.*,b.`name` as nodename,c.username from pz_article as a,pz_node as b,pz_user as c where a.nodeid = b.id and a.uid = c.id and a.title like ? and a.nodeid = ?  limit ? offset ?", "%"+kw+"%", nodeid, mp, (cp-1)*mp).Scan(&articles)

	return ApiJson{State: true, Msg: articles, Count: count}
}
