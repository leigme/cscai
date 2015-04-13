package models

import (
	"time"
)

type Dish struct {
	Id         int64 `orm:"auto"`
	Name       string
	Synopsis   string `orm:"type(text)"`
	Picurl     string
	Price      int64
	Shopid     int64
	Click      int64
	Entrytime  time.Time `orm:"type(date)"`
	Modifytime time.Time `orm:"type(date)"`
	Status     bool
}
