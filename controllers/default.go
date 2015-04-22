package controllers

import (
	"cscai/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

type Lead struct {
	Title   string
	Content string
}

type Article struct {
	Title   string
	Content string
}

type RightList struct {
	Title  string
	Link   string
	Active bool
}

type Ajaxdish struct {
	models.Dish
	Createdtime string
	Updatedtime string
}

func (c *MainController) Get() {
	c.Data["PagTitle"] = "首页"
	j := 10

	var lead Lead
	lead.Title = "世界，你好！"
	lead.Content = "长沙菜美食风味小站！"
	c.Data["Lead"] = lead

	articles := Getajax()
	c.Data["Articles"] = articles

	var rightlist RightList
	rightlists := make(map[int]RightList)
	for i := 0; i < j; i++ {
		rightlist.Title = "测试热点" + strconv.Itoa(i)
		rightlist.Link = "/"
		if i == 0 {
			rightlist.Active = true
		}
		rightlists[i] = rightlist
	}
	c.Data["RightLists"] = rightlists
	c.TplNames = "home.html"
}

func Getajax() map[int]Ajaxdish {
	pagesize, _ := strconv.ParseInt(beego.AppConfig.String("homepagesize"), 10, 64)
	var ajaxdish Ajaxdish
	var dish models.Dish
	ajaxdishes := make(map[int]Ajaxdish)
	dishes := make([]*models.Dish, pagesize)
	dishes, _ = dish.Get(pagesize-1, 0)
	for i, dish := range dishes {
		ajaxdish.Id = dish.Id
		ajaxdish.Name = dish.Name
		ajaxdish.Synopsis = dish.Synopsis
		ajaxdish.Picurl = dish.Picurl
		ajaxdish.Price = dish.Price
		ajaxdish.Click = dish.Click
		ajaxdish.Createdtime = dish.Createdtime.Format("2006-01-02")
		ajaxdish.Updatedtime = dish.Updatedtime.Format("2006-01-02")
		ajaxdishes[i] = ajaxdish
	}
	return ajaxdishes
}
