package controllers

import (
	"cscai/models"
	"github.com/astaxie/beego"
	"time"
	//"path"
	//"strconv"
)

type AdminController struct {
	beego.Controller
}

type Dish struct {
	models.Dish
	Entrytime  string
	Modifytime string
}

func (c *AdminController) Get() {
	c.Data["PagTitle"] = "登录"
	c.TplNames = "admin/signin.html"
}

func (c *AdminController) Login() {
	username := c.Input().Get("inputEmail")
	password := c.Input().Get("inputPassword")
	if username == "admin@admin.me" {
		if password == "password" {
			c.SetSession("IsLogin", true)
			c.Manage()
		} else {
			c.Get()
		}
	} else {
		c.Get()
	}
}

func (c *AdminController) Signout() {
	if c.GetSession("IsLogin") == true {
		c.SetSession("IsLogin", false)
		beego.Info("登出系统！")
		c.Get()
	} else {
		c.SetSession("IsLogin", false)
		beego.Info("没有登录！")
		c.Get()
	}
}

func (c *AdminController) Manage() {
	islogin := c.GetSession("IsLogin")
	beego.Info(islogin)
	if islogin == true {
		beego.Info("登录成功！")
	} else {
		beego.Info("登录失败！")
	}
	c.Data["PagTitle"] = "管理"
	c.Data["SectionTitle"] = "菜品管理"
	var dish Dish
	dishes := make(map[int]Dish, 10)
	for i := 0; i < 10; i++ {
		dish.Id = 1
		dish.Name = "剁椒鱼头"
		dish.Synopsis = "这些菜"
		dish.Price = 6000 / 100
		dish.Click = 12
		dish.Entrytime = time.Now().Format("2006-01-02")
		dish.Modifytime = time.Now().Format("2006-01-02")
		dishes[i] = dish
	}
	c.Data["Dishes"] = dishes

	uploadimageurl := beego.AppConfig.String("uploadimageurl")

	beego.Info(uploadimageurl)
	beego.Info(time.Now().Month())
	c.TplNames = "admin/index.html"
}

func (c *AdminController) Moddish() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id)
	c.Manage()
}

func (c *AdminController) Deldish() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id)
	c.Manage()
}

func (c *AdminController) Adddish() {
	type Dish struct {
		models.Dish
	}
	var dish Dish
	dish.Add("剁椒鱼头", "这些菜", "", time.Now(), time.Now(), 60, true)
	dish.Delete(2)
	dish.Select(beego.AppConfig.String("backstagepagesize"), "")
	beego.Info(beego.AppConfig.String("backstagepagesize"))

	//id := 1
	name := c.Input().Get("name")
	//	_, picurl, err := c.GetFile("picurl")
	synopsis := c.Input().Get("synopsis")
	price := c.Input().Get("price")
	entrytime := c.Input().Get("entrytime")
	status := c.Input().Get("status")
	//	uploadname := strconv.Itoa(id) + picurl.Filename
	//	uploadurl := path.Join("uploads/images/", uploadname)
	//	c.SaveToFile("picurl", uploadurl)

	beego.Info(name)
	//	beego.Info(picurl.Filename)
	//	beego.Info(uploadurl)
	beego.Info(synopsis)
	beego.Info(price)
	beego.Info(entrytime)
	beego.Info(status)
	c.Manage()
}
