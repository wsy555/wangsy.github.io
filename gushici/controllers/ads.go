/**********************************************
** @Des: 文章
** @Author: wangsy
** @Date:   2017-12-09 14:17:37
***********************************************/
package controllers


type AdsController struct {
	BaseController
}

func (self *AdsController) Index() {
	self.Data["siteName"] = "2018年你靠什么吃饭"
	self.Data["pageTitle"] = "2018年你靠什么吃饭"
	self.display()
}

func (self *AdsController) Show() {
	self.Data["siteName"] = "2018年你靠什么吃饭"
	self.Data["pageTitle"] = "2018年你靠什么吃饭"

	self.TplName = "ads/show.html"
	self.display()
}

func (self *AdsController) ImageShow() {
	self.Data["siteName"] = "2018年你靠什么吃饭"
	self.Data["pageTitle"] = "2018年你靠什么吃饭"
	self.TplName = "ads/imageshow.html"
}


