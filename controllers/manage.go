package controllers

import (
	"cscai/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type ManageController struct {
	beego.Controller
}

type Dish struct {
	models.Dish
	Updatedtime string
}

func (c *ManageController) Get() {
	op := c.Input().Get("op")
	if len(op) == 0 {
		c.TplNames = "admin/index.html"
	} else {
		id := c.Input().Get("id")
		if len(id) == 0 {
			c.TplNames = "admin/index.html"
		} else {
			switch op {
			case "moduser":
			case "deluser":
			case "modgroup":
			case "delgroup":
			case "moddish":
				beego.Info("进入修改模块！")
				sid := c.Input().Get("id")
				id, _ := strconv.ParseInt(sid, 10, 64)
				var dish models.Dish
				dishone, err := dish.Getone(id)
				var moddish Dish
				moddish.Id = dishone.Id
				moddish.Name = dishone.Name
				moddish.Synopsis = dishone.Synopsis
				moddish.Price = dishone.Price
				moddish.Picurl = dishone.Picurl
				moddish.Updatedtime = dishone.Updatedtime.Format("2006-01-02 15:04")
				moddish.Status = dishone.Status
				if err == nil {
					c.Data["PagTitle"] = "修改"
					c.Data["Pagename"] = "modifydish"
					c.Data["SectionTitle"] = "菜品修改"
					c.Data["Dish"] = moddish
					c.TplNames = "admin/index.html"
				} else {
					beego.Info("跳出修改模块！")
					c.Redirect("admin/manage", 301)
				}
			case "deldish":
			}
		}
	}
}

func (c *ManageController) User() {

}

func (c *ManageController) Group() {

}

func (c *ManageController) Dish() {
	sid := c.Input().Get("id")
	id, _ := strconv.ParseInt(sid, 10, 64)
	name := c.Input().Get("name")
	sprice := c.Input().Get("price")
	price, _ := strconv.ParseInt(sprice, 10, 64)
	picurl := c.Input().Get("picurl")
	supdatedtime := c.Input().Get("updatedtime")
	updatedtime, _ := time.Parse("2006-01-02 15:04", supdatedtime)
	sstatus := c.Input().Get("status")
	var status bool
	if sstatus == "true" {
		status = true
	} else {
		status = false
	}
	synopsis := c.Input().Get("synopsis")
	dish := models.Dish{
		Name:        name,
		Price:       price,
		Picurl:      picurl,
		Updatedtime: updatedtime,
		Status:      status,
		Synopsis:    synopsis,
	}
	if sid == "" {
		dish.Add(dish)
	} else {
		dish.Mod(id, dish)
	}
	c.Redirect("/admin/manage", 301)
}
