package models

import "github.com/beego/beego/v2/client/orm"

type Motor struct {
	Id    int
	Brand string
	Model string
	Year  int
	Price float64
	Image string
}

func init() {
	orm.RegisterModel(new(Motor))
}

func AddMotor(motor Motor) {
	o := orm.NewOrm()
	o.Insert(&motor)
}

func GetMotorsByUserId(userId int) []Motor {
	o := orm.NewOrm()
	var motors []Motor
	o.QueryTable("motor").Filter("UserId", userId).All(&motors)
	return motors
}

func GetMotorById(id int) Motor {
	o := orm.NewOrm()
	motor := Motor{Id: id}
	o.Read(&motor)
	return motor
}

func UpdateMotor(motor Motor) {
	o := orm.NewOrm()
	o.Update(&motor)
}

func DeleteMotor(id int) {
	o := orm.NewOrm()
	o.Delete(&Motor{Id: id})
}
