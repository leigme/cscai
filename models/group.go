package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Group struct {
	Id          int64 `orm:"auto"`
	Name        string
	Remark      string
	Createdtime time.Time `orm:"auto_now_add;index;type(datetime)"`
	Updatedtime time.Time `orm:"index;type(datetime)"`
	Status      bool
}

func (g *Group) Getcount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("group").Count()
	if err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func (g *Group) Get(sid, eid int64) ([]*Group, error) {
	o := orm.NewOrm()
	groups := make([]*Group, 0)
	qs := o.QueryTable("group")
	_, err := qs.Filter("id", sid, eid).All(&groups)
	return groups, err
}

func (g *Group) Add(group Group) error {
	o := orm.NewOrm()
	_, err := o.Insert(&group)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (g *Group) Del(id int64) error {
	group := &Group{Id: id}
	o := orm.NewOrm()
	_, err := o.Delete(group)
	return err
}
