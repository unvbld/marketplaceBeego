package routers

import (
	"motorMarketplace/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// Route for listing all motors
	web.Router("/motor", &controllers.MotorController{}, "get:GetAllMotor")

	// Route to show form to create a new motor
	web.Router("/motor/new", &controllers.MotorController{}, "get:NewMotor")

	// Route for creating a new motor (POST)
	web.Router("/motor/create", &controllers.MotorController{}, "post:CreateMotor")

	// Route to show form for editing a motor (GET)
	web.Router("/motor/edit/:id", &controllers.MotorController{}, "get:EditMotor")

	// Route for updating motor data (POST)
	web.Router("/motor/update/:id", &controllers.MotorController{}, "post:UpdateMotor")

	// Route for deleting a motor (POST)
	web.Router("/motor/delete/:id", &controllers.MotorController{}, "post:DeleteMotor")
}
