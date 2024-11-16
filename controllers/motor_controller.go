package controllers

import (
	"motorMarketplace/models"
	"os"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

// MotorController handles motor-related actions
type MotorController struct {
	web.Controller
}

// Get all motors
func (c *MotorController) GetAllMotor() {
	o := orm.NewOrm()
	var motors []models.Motor
	o.QueryTable("motor").All(&motors)
	c.Data["Motor"] = motors
	c.TplName = "motor/index.html"
}

// Create motor form
func (c *MotorController) NewMotor() {
	c.TplName = "motor/new.html"
}

// Save new motor (with image upload)
func (c *MotorController) CreateMotor() {
	o := orm.NewOrm()
	price, _ := strconv.ParseFloat(c.GetString("price"), 64)
	year, _ := strconv.Atoi(c.GetString("year"))

	// Handling the image upload
	_, header, err := c.GetFile("image") // Get the uploaded file
	if err != nil {
		c.Ctx.WriteString("Error uploading image")
		return
	}
	defer c.Ctx.Request.Body.Close()

	// Save the image to the static/uploads folder
	imagePath := "/static/uploads/" + header.Filename
	err = c.SaveToFile("image", "static/uploads/"+header.Filename)
	if err != nil {
		c.Ctx.WriteString("Error saving image")
		return
	}

	motor := models.Motor{
		Brand: c.GetString("brand"),
		Model: c.GetString("model"),
		Year:  year,
		Price: price,
		Image: imagePath, // Save image path in database
	}

	o.Insert(&motor)
	c.Redirect("/motor", 302)
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

// Update motor (with image update)
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

		// Handling image update
		_, header, err := c.GetFile("image") // Check if a new image is uploaded
		if err == nil {
			defer c.Ctx.Request.Body.Close()

			// Remove the old image file if exists (optional)
			if motor.Image != "" {
				oldImagePath := "static" + motor.Image
				_ = os.Remove(oldImagePath) // Remove old image
			}

			// Save the new image
			imagePath := "/static/uploads/" + header.Filename
			err = c.SaveToFile("image", "static/uploads/"+header.Filename)
			if err != nil {
				c.Ctx.WriteString("Error saving image")
				return
			}
			motor.Image = imagePath // Update the image path in the database
		}

		o.Update(&motor)
	}
	c.Redirect("/motor", 302)
}

// Delete motor
func (c *MotorController) DeleteMotor() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	motor := models.Motor{Id: id}
	o.Read(&motor)

	// Optionally delete the image file if exists
	if motor.Image != "" {
		imagePath := "static" + motor.Image
		_ = os.Remove(imagePath) // Remove the image from the filesystem
	}

	o.Delete(&motor)
	c.Redirect("/motor", 302)
}
