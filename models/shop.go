package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Shop struct {
	Id          int64 `orm:"auto"`
	Name        string
	Synopsis    string `orm:"size(20000)"`
	Picurl      string
	Price       int64
	Phone       string
	Address     string
	Coordinate  string
	Userid      int64
	Click       int64     `orm:"index"`
	Luad        int64     `orm:"index"`
	Createdtime time.Time `orm:"auto_now_add;index;type(datetime)"`
	Updatedtime time.Time `orm:"index;type(datetime)"`
	Status      bool
}

func (s *Shop) Getcount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("shop").Count()
	if err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func (s *Shop) Get(sid, eid int64) ([]*Shop, error) {
	o := orm.NewOrm()
	shops := make([]*Shop, 0)
	qs := o.QueryTable("shop")
	_, err := qs.Filter("id", sid, eid).All(&shops)
	return shops, err
}

func (s *Shop) Add(shop Shop) error {
	o := orm.NewOrm()
	_, err := o.Insert(&shop)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Shop) Del(id int64) error {
	shop := &Shop{Id: id}
	o := orm.NewOrm()
	_, err := o.Delete(shop)
	return err
}
