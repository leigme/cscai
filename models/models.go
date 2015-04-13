package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func (d *Dish) Select(pagsize, sid string) {
	if pagsize == "" && sid != "" {
		id, _ := strconv.ParseInt(sid, 10, 64)
		beego.Info(id)
	} else if pagsize != "" && sid == "" {
		beego.Info("查询所有的菜品！")
	} else {
		beego.Info("参数传递错误！")
	}
}

func (d *Dish) Add(name, synopsis, picurl string, entrytime, modifytime time.Time, price int64, status bool) {
	o := orm.NewOrm()
	d.Name = name
	d.Synopsis = synopsis
	d.Picurl = picurl
	d.Entrytime = entrytime
	d.Modifytime = modifytime
	d.Price = price
	d.Status = status
	o.Insert(d)
	beego.Info("添加一条数据到数据库！")
}

func (d *Dish) Modify(id int64) {
	beego.Info("修改一条数据库的数据！")
}

func (d *Dish) Delete(id int64) {
	beego.Info("删除一条数据库的数据！")
}

func init() {
	orm.RegisterModel(new(Dish), new(Username), new(Shop), new(Group))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "cscai:cscai@/cscai?charset=utf8", 30)
}
