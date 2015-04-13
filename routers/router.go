package routers

import (
	"cscai/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//前台页面路由
	beego.Router("/", &controllers.MainController{})
	//超级管理路由
	beego.Router("/manage/addusername", &controllers.ManageController{}, "post:Addusername")
	beego.Router("/manage/modusername/?:id", &controllers.ManageController{}, "get:Modusername")
	beego.Router("/manage/delusername/?:id", &controllers.ManageController{}, "get:Delusername")
	//权限组管理路由
	beego.Router("/manage/addgroup", &controllers.ManageController{}, "post:Addgroup")
	beego.Router("/manage/modgroup/?:id", &controllers.ManageController{}, "get:Modgroup")
	beego.Router("/manage/delgroup/?:id", &controllers.ManageController{}, "get:Delgroup")
	//用户登录路由
	beego.Router("/admin/signin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/admin/signout", &controllers.AdminController{}, "get:Signout")
	//菜品管理路由
	beego.Router("/admin/adddish", &controllers.AdminController{}, "post:Adddish")
	beego.Router("/admin/moddish/?:id", &controllers.AdminController{}, "get:Moddish")
	beego.Router("/admin/deldish/?:id", &controllers.AdminController{}, "get:Deldish")

	//绕过登录的测试管理页面
	beego.Router("/admin/manage", &controllers.AdminController{}, "get:Manage")
}
