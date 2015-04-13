package controllers

import (
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

func (c *MainController) Get() {
	c.Data["PagTitle"] = "首页"
	j := 10

	var lead Lead
	lead.Title = "世界，你好！"
	lead.Content = "长沙菜美食风味小站！"
	c.Data["Lead"] = lead

	var article Article
	articles := make(map[int]Article, j-1)
	for i := 0; i < j; i++ {
		article.Title = "测试标题"
		article.Content = "测试文章内容容纳一些字符！！！"
		articles[i] = article
	}
	c.Data["Articles"] = articles

	var rightlist RightList
	rightlists := make(map[int]RightList, j-1)
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
