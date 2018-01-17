/**********************************************
** @Des: 文章
** @Author: wangsy
** @Date:   2017-12-09 14:17:37
***********************************************/
package controllers

import (
	"PPGo_ApiAdmin/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"math/rand"
)

type WwwController struct {
	BaseController
}

func (self *WwwController) Index() {

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	filters = append(filters, "class_id", 5)
	result, _ := models.NewsGetList(1, 6, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["class_id"] = v.ClassId

		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = v.Picurl
		row["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row["desc"] = string(nameRune[:lth])
		}

		row["linkurl"] = v.Linkurl
		row["author"] = v.Author
		list[k] = row
	}

	filters2 := make([]interface{}, 0)
	filters2 = append(filters2, "status", 1)
	filters2 = append(filters2, "class_id", 3)
	result2, _ := models.NewsGetList(1, 6, filters2...)
	list2 := make([]map[string]interface{}, len(result2))
	for k, v := range result2 {
		row2 := make(map[string]interface{})
		row2["id"] = v.Id
		row2["title"] = v.Title
		row2["class_id"] = v.ClassId
		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row2["picurl"] = v.Picurl
		row2["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row2["desc"] = string(nameRune[:lth])
		}

		row2["linkurl"] = v.Linkurl
		row2["author"] = v.Author
		list2[k] = row2
	}

	//国学经典
	filters3 := make([]interface{}, 0)
	filters3 = append(filters3, "status", 1)
	filters3 = append(filters3, "class_id", 1)
	result3, _ := models.NewsGetList(1, 16, filters3...)
	list3 := make([]map[string]interface{}, len(result3))
	for k, v := range result3 {
		row2 := make(map[string]interface{})
		row2["id"] = v.Id
		row2["title"] = v.Title
		row2["class_id"] = v.ClassId
		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row2["picurl"] = v.Picurl
		row2["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row2["desc"] = string(nameRune[:lth])
		}

		row2["linkurl"] = v.Linkurl
		row2["author"] = v.Author
		list3[k] = row2
	}

	//诗词古韵
	filters4 := make([]interface{}, 0)
	filters4 = append(filters4, "status", 1)
	filters4 = append(filters4, "class_id", 2)
	result4, _ := models.NewsGetList(1, 6, filters4...)
	list4 := make([]map[string]interface{}, len(result4))
	for k, v := range result4 {
		row2 := make(map[string]interface{})
		row2["id"] = v.Id
		row2["title"] = v.Title
		row2["class_id"] = v.ClassId
		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row2["picurl"] = v.Picurl
		row2["media"] = v.Media

		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row2["desc"] = string(nameRune[:lth])
		}

		row2["linkurl"] = v.Linkurl
		row2["author"] = v.Author
		list4[k] = row2
	}


	out := make(map[string]interface{})
	out["list"] = list
	out["list2"] = list2
	out["list3"] = list3
	out["list4"] = list4
	out["class_id"] = 0
	self.Data["data"] = out

	self.Layout = "public/www_layout.html"
	self.display()
}

func (self *WwwController) Show() {

	id, _ := self.GetInt(":id")
	News, _ := models.NewsGetById(id)
	NewsNext, _ := models.NewsGetNextById(id)

	nextRow := make(map[string]interface{})

	if(NewsNext != nil ){
		nextRow["id"] = NewsNext.Id
		nextRow["title"] = NewsNext.Title
	}
	row := make(map[string]interface{})
	row["class_id"] = 0
	if (News != nil) {
		row["id"] = News.Id
		row["title"] = News.Title
		row["class_id"] = News.ClassId
		row["desc"] = News.Desc
		row["content"] = News.Content
		if(string(News.Picurl) == "") {
			var r = rand.Intn(10)
			News.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = News.Picurl
		row["linkurl"] = News.Linkurl
		row["media"] = News.Media
		row["author"] = News.Author
		row["posttime"] = beego.Date(time.Unix(News.Posttime, 0), "Y/m/d")
	}
	fmt.Println(row)
	row["next"] = nextRow
	self.Data["data"] = row

	self.Layout = "public/www_layout.html"
	self.display()
}

func (self *WwwController) List() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 16
	}
	catId, cerr_ := self.GetInt(":class_id")
	fmt.Println(catId)
	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if cerr_ == nil {
		filters = append(filters, "class_id", catId)
	}
	result, count := models.NewsGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["class_id"] = v.ClassId

		if(string(v.Picurl) == "") {
			var r = rand.Intn(10)
			v.Picurl = "/uploads/image/rand" + fmt.Sprintf("%d", r) + ".jpeg"
		}
		row["picurl"] = v.Picurl
		row["media"] = v.Media
		if (v.Desc != "") {
			nameRune := []rune(v.Desc)
			lth := len(nameRune)
			if(lth > 30) {
				lth = 30
			}
			row["desc"] = string(nameRune[:lth])
		}


		row["linkurl"] = v.Linkurl
		row["author"] = v.Author
		row["posttime"] = beego.Date(time.Unix(v.Posttime, 0), "Y-m-d")
		list[k] = row
	}

	classArr := make(map[int]string)
	classArr[5] = "开心儿歌"
	classArr[3] = "儿童古诗"
	classArr[2] = "诗词古韵"
	classArr[1] = "经典国学"

	out := make(map[string]interface{})
	out["count"] = count
	out["class_id"] = catId
	out["page"] = page
	out["class_name"] = classArr[catId]
	out["title"] = classArr[catId]
	out["list"] = list
	self.Data["data"] = out

	self.Layout = "public/www_layout.html"
	self.display()
}


