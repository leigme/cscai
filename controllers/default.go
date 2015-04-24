package controllers

import (
	"cscai/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

type Introduction struct {
	Title    string
	Synopsis string
}

type RightList struct {
	Title  string
	Link   string
	Active bool
}

type Ajaxdish struct {
	models.Dish
	Lead        string
	Updatedtime string
}

func (c *MainController) Get() {
	c.Data["PagTitle"] = "首 页"
	j := 10

	var introduction Introduction
	introduction.Title = "世界，你好！"
	introduction.Synopsis = "长沙菜美食风味小站！"
	c.Data["Introduction"] = introduction

	dishes := Getajax()
	c.Data["Dishes"] = dishes

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
		ajaxdish.Name = dish.Name
		ajaxdish.Synopsis = dish.Synopsis
		ajaxdish.Picurl = dish.Picurl
		ajaxdish.Price = dish.Price
		ajaxdish.Click = dish.Click
		ajaxdish.Updatedtime = dish.Updatedtime.Format("2006-01-02")
		ajaxdishes[i] = ajaxdish
	}
	return ajaxdishes
}
