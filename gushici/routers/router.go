package routers

import (
	"PPGo_ApiAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WwwController{}, "*:Index")
	beego.Router("/show/:id", &controllers.WwwController{}, "*:Show")
	beego.Router("/list/:class_id", &controllers.WwwController{}, "*:List")

	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.Router("/news/list", &controllers.NewsController{}, "*:List")
	beego.Router("/news/edit", &controllers.NewsController{}, "*:Edit")

	beego.Router("/ads/index", &controllers.AdsController{}, "*:Index")
	beego.Router("/ads/show", &controllers.AdsController{}, "*:Show")
	beego.Router("/ads/image_show", &controllers.AdsController{}, "*:ImageShow")

	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.ApiDocController{})
	beego.AutoRouter(&controllers.ApiMonitorController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.CodeController{})

	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.NewsController{})
	beego.ErrorController(&controllers.ErrorController{})
}
