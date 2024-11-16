package routers

import (
	"motorMarketplace/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/motor", &controllers.MotorController{}, "get:GetAllMotor")
	web.Router("/motor/new", &controllers.MotorController{}, "get:NewMotor")
	web.Router("/motor/create", &controllers.MotorController{}, "post:CreateMotor")
	web.Router("/motor/edit/:id", &controllers.MotorController{}, "get:EditMotor")
	web.Router("/motor/update/:id", &controllers.MotorController{}, "post:UpdateMotor")
	web.Router("/motor/delete/:id", &controllers.MotorController{}, "post:DeleteMotor")
}
