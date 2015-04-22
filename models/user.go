package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id          int64 `orm:"auto"`
	Name        string
	Password    string
	Phone       string
	Address     string
	Group       string
	Createdtime time.Time `orm:"auto_now_add;index;type(datetime)"`
	Updatedtime time.Time `orm:"index;type(datetime)"`
	Status      bool
}

func (u *User) Getcount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("user").Count()
	if err != nil {
		return count, err
	} else {
		return count, nil
	}
}

func (u *User) Get(sid, eid int64) ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.Filter("id", sid, eid).All(&users)
	return users, err
}

func (u *User) Add(user User) error {
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (u *User) Del(id int64) error {
	user := &User{Id: id}
	o := orm.NewOrm()
	_, err := o.Delete(user)
	return err
}
