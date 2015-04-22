package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	orm.RegisterModel(new(Dish), new(User), new(Shop), new(Group))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "cscai:cscai@/cscai?charset=utf8", 30)
}
