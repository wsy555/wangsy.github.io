/**********************************************
** @Des: 文章
** @Author: wangsy
** @Date:   2017-12-09 14:17:37
***********************************************/
package controllers

import (
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"PPGo_ApiAdmin/models"
)

type NewsController struct {
	BaseController
}

func (self *NewsController) List() {
	self.Data["pageTitle"] = "资讯管理"
	class_id ,_ := self.GetInt("class_id")
	pageSize := 10
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.ClassGetList(1, pageSize, filters...)
	classList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["class_name"] = v.ClassName
		row["desc"] = v.Desc
		row["linkurl"] = v.Linkurl
		row["orderid"] = v.Orderid
		row["count"] = count
		classList[k] = row
	}
	self.Data["news_class"] = classList
	self.Data["class_id"] = class_id

	self.display()
}

func (self *NewsController) Add() {
	self.Data["pageTitle"] = "新增资讯"

	pageSize := 10
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.ClassGetList(1, pageSize, filters...)

	classList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["class_name"] = v.ClassName
		row["desc"] = v.Desc
		row["linkurl"] = v.Linkurl
		row["orderid"] = v.Orderid
		row["count"] = count
		classList[k] = row
	}
	self.Data["news_class"] = classList

	self.display()
}

func (self *NewsController) Edit() {
	self.Data["pageTitle"] = "编辑资讯"
	id, _ := self.GetInt("id", 0)
	News, _ := models.NewsGetById(id)

	pageSize := 10
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.ClassGetList(1, pageSize, filters...)

	classList := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["class_name"] = v.ClassName
		row["desc"] = v.Desc
		row["linkurl"] = v.Linkurl
		row["orderid"] = v.Orderid
		row["count"] = count
		classList[k] = row
	}

	row := make(map[string]interface{})
	row["id"] = News.Id
	row["title"] = News.Title
	row["class_id"] = News.ClassId
	row["orderid"] = News.Orderid
	row["keywords"] = News.Keywords
	row["used"] = News.Used
	row["desc"] = News.Desc
	row["content"] = News.Content
	row["pic_url"] = News.Picurl
	row["media"] = News.Media
	row["author"] = News.Author
	row["posttime"] = beego.Date(time.Unix(News.Posttime, 0), "Y/m/d")

	self.Data["news_class"] = classList
	self.Data["news"] = row
	self.display()
}

func (self *NewsController) AjaxDel() {

	NewsId, _ := self.GetInt("id")
	News, _ := models.NewsGetById(NewsId)
	News.Updatetime = time.Now().Unix()
	News.Status = 2
	News.Id = NewsId

	if err := News.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *NewsController) AjaxSave() {
	id, _ := self.GetInt("id")
	classId, _ := self.GetInt("class_id")
	orderid, _ := self.GetInt("orderid")

	News := new(models.InfoList)
	News.Title = strings.TrimSpace(self.GetString("title"))
	News.Author = strings.TrimSpace(self.GetString("author"))
	News.Keywords = strings.TrimSpace(self.GetString("keywords"))
	News.Used,_ = self.GetInt("used")
	News.Desc = strings.TrimSpace(self.GetString("desc"))
	News.Content = strings.TrimSpace(self.GetString("content"))
	News.ClassId = classId
	News.Orderid = orderid
	News.Updatetime = time.Now().Unix()
	News.Picurl = strings.TrimSpace(self.GetString("pic_url"))
	News.Media = strings.TrimSpace(self.GetString("media"))
	News.Status = 1
	if id == 0 {
		News.Posttime = time.Now().Unix()
		if _, err := models.NewsAdd(News); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}
	News.Id = id
	if err := News.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *NewsController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}
	class_id ,_ := self.GetInt("class_id")

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if class_id > 0 {
		filters = append(filters, "class_id", class_id)
	}

	result, count := models.NewsGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["class_id"] = v.ClassId
		row["orderid"] = v.Orderid
		row["keywords"] = v.Keywords
		row["used"] = v.Used
		row["desc"] = v.Desc
		row["pic_url"] = v.Picurl
		row["author"] = v.Author
		row["posttime"] = beego.Date(time.Unix(v.Posttime, 0), "Y-m-d")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

//上传图片
func (self *NewsController) Upload() {

	filepath := "static/upload/" + beego.Date(time.Now(), "Y-m-d") + "/"
	_, err := os.Stat(filepath)
	if err != nil {
		err := os.MkdirAll(filepath, 0777)
		if err != nil {
			self.ajaxMsg("0", MSG_ERR)
		}
	}
	self.UploadFile("upfile", filepath)
}
