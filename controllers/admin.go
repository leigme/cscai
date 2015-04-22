package controllers

import (
	"cscai/models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	islogin := c.GetSession("IsLogin")
	beego.Info(islogin)
	if islogin == true {
		beego.Info("登录成功！")
	} else {
		beego.Info("登录失败！")
	}
	c.Data["PagTitle"] = "管理"
	c.Data["Pagename"] = "dish"
	c.Data["SectionTitle"] = "菜品管理"
	var dish models.Dish
	dishes := make([]*models.Dish, 0)
	dishes, _ = dish.Get(10, 0)

	c.Data["Dishes"] = dishes

	c.TplNames = "admin/index.html"
}

func (c *AdminController) Signin() {
	c.Data["PagTitle"] = "登录"
	c.TplNames = "admin/signin.html"
}

func (c *AdminController) Post() {
	username := c.Input().Get("inputEmail")
	password := c.Input().Get("inputPassword")
	if username == "admin@admin.me" {
		if password == "password" {
			c.SetSession("IsLogin", true)
			c.Redirect("/admin/manage", 301)
		} else {
			c.Redirect("/admin/signin", 301)
		}
	} else {
		c.Redirect("/admin/signin", 301)
	}
}

func (c *AdminController) Login() {
	if c.GetSession("IsLogin") == true {

	}
}

func (c *AdminController) Signout() {
	if c.GetSession("IsLogin") == true {
		c.SetSession("IsLogin", false)
		beego.Info("登出系统！")
		c.Redirect("/admin/signin", 301)
	} else {
		c.SetSession("IsLogin", false)
		beego.Info("没有登录！")
		c.Redirect("/admin/signin", 301)
	}
}
