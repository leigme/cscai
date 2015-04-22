package routers

import (
	"cscai/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//前台页面路由
	beego.Router("/", &controllers.MainController{})
	//超级管理路由
	beego.Router("/admin/controller", &controllers.ManageController{})
	beego.Router("/controller/user", &controllers.ManageController{}, "post:User")
	//权限组管理路由
	beego.Router("/controller/group", &controllers.ManageController{}, "post:Group")
	//菜品管理路由
	beego.Router("/controller/dish", &controllers.ManageController{}, "post:Dish")

	//用户登录控制
	beego.Router("/admin/manage", &controllers.AdminController{})
	beego.Router("/admin/signin", &controllers.AdminController{}, "get:Signin")
	beego.Router("/admin/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/admin/signout", &controllers.AdminController{}, "get:Signout")

}
