package routers

import (
	"motorMarketplace/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// Menangani URL root
	web.Router("/", &controllers.MotorController{}, "get:GetAllMotors")

	// Menampilkan daftar motor
	web.Router("/motors", &controllers.MotorController{}, "get:GetAllMotors")

	// Form untuk menambah motor baru
	web.Router("/motors/new", &controllers.MotorController{}, "get:NewMotor")

	// Menambahkan motor baru (post request)
	web.Router("/motors", &controllers.MotorController{}, "post:CreateMotor")

	// Form untuk mengedit motor berdasarkan ID
	web.Router("/motors/edit/:id", &controllers.MotorController{}, "get:EditMotor")
	web.Router("/motors/edit/:id", &controllers.MotorController{}, "post:UpdateMotor")

	// Menghapus motor berdasarkan ID
	web.Router("/motors/delete/:id", &controllers.MotorController{}, "get:DeleteMotor")
}
