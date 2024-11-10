package controllers

import (
	"motorMarketplace/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type MotorController struct {
	web.Controller
}

// Get all motors
func (c *MotorController) GetAllMotors() {
	o := orm.NewOrm()
	var motors []models.Motor
	o.QueryTable("motor").All(&motors)
	c.Data["Motors"] = motors
	c.TplName = "motor/index.html"
}

// Create motor form
func (c *MotorController) NewMotor() {
	c.TplName = "motor/new.html"
}

// Save new motor
func (c *MotorController) CreateMotor() {
	o := orm.NewOrm()
	price, _ := strconv.ParseFloat(c.GetString("price"), 64)
	year, _ := strconv.Atoi(c.GetString("year"))
	motor := models.Motor{
		Brand: c.GetString("brand"),
		Model: c.GetString("model"),
		Year:  year,
		Price: price,
	}
	o.Insert(&motor)
	c.Redirect("/motors", 302)
}

// Edit motor form
func (c *MotorController) EditMotor() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	motor := models.Motor{Id: id}
	o.Read(&motor)
	c.Data["Motor"] = motor
	c.TplName = "motor/edit.html"
}

// Update motor
func (c *MotorController) UpdateMotor() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	price, _ := strconv.ParseFloat(c.GetString("price"), 64)
	year, _ := strconv.Atoi(c.GetString("year"))
	motor := models.Motor{Id: id}
	if o.Read(&motor) == nil {
		motor.Brand = c.GetString("brand")
		motor.Model = c.GetString("model")
		motor.Year = year
		motor.Price = price
		o.Update(&motor)
	}
	c.Redirect("/motors", 302)
}

// Delete motor
func (c *MotorController) DeleteMotor() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	motor := models.Motor{Id: id}
	o.Delete(&motor)
	c.Redirect("/motors", 302)
}
