package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Motor struct {
	Id    int    `orm:"auto"`
	Brand string `orm:"size(100)"`
	Model string `orm:"size(100)"`
	Year  int
	Price float64
}

func init() {
	orm.RegisterModel(new(Motor))
}
