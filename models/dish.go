package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Dish struct {
	Id          int64 `orm:"auto"`
	Name        string
	Synopsis    string `orm:"size(20000)"`
	Picurl      string
	Price       int64
	Shopid      int64
	Click       int64
	Createdtime time.Time `orm:"auto_now_add;index;type(datetime)"`
	Updatedtime time.Time `orm:"index;type(datetime)"`
	Status      bool
}

func (d *Dish) Getcount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("dish").Count()
	if err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func (d *Dish) Get(sid, eid int64) ([]*Dish, error) {
	o := orm.NewOrm()
	dishes := make([]*Dish, 0)
	qs := o.QueryTable("dish")
	_, err := qs.OrderBy("-Updatedtime").Limit(sid, eid).All(&dishes)
	return dishes, err
}

func (d *Dish) Getone(id int64) (Dish, error) {
	o := orm.NewOrm()
	var dish Dish
	err := o.QueryTable("dish").Filter("Id", id).One(&dish)
	return dish, err
}

func (d *Dish) Add(dish Dish) error {
	o := orm.NewOrm()
	_, err := o.Insert(&dish)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *Dish) Mod(id int64, dish Dish) error {
	o := orm.NewOrm()
	dishid := Dish{Id: id}
	if o.Read(&dishid) == nil {
		dishid.Name = dish.Name
		dishid.Synopsis = dish.Synopsis
		dishid.Picurl = dish.Picurl
		dishid.Price = dish.Price
		dishid.Shopid = dish.Shopid
		dishid.Updatedtime = dish.Updatedtime
		dishid.Status = dish.Status
		if _, err := o.Update(&dishid); err == nil {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (d *Dish) Del(id int64) error {
	dish := &Dish{Id: id}
	o := orm.NewOrm()
	_, err := o.Delete(dish)
	return err
}
