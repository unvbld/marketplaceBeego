package main

import (
	"motorMarketplace/controllers"
	_ "motorMarketplace/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

func init() {
	// Ambil konfigurasi dari app.conf
	dbUser := web.AppConfig.DefaultString("db_user", "root")
	dbPass := web.AppConfig.DefaultString("db_password", "")
	dbName := web.AppConfig.DefaultString("db_name", "motorMarketplace")
	dbHost := web.AppConfig.DefaultString("db_host", "127.0.0.1")
	dbPort := web.AppConfig.DefaultString("db_port", "3306")

	// Format koneksi MySQL
	dataSource := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"

	// Registrasi database dengan alias 'default'
	orm.RegisterDataBase("default", "mysql", dataSource)

	// Aktifkan mode debug untuk ORM
	orm.Debug = true
}

func main() {
	web.Router("/", &controllers.MotorController{}, "get:GetAllMotors")
	web.Router("/motors", &controllers.MotorController{}, "get:GetAllMotors")
	web.Router("/motors/new", &controllers.MotorController{}, "get:NewMotor")
	web.Router("/motors", &controllers.MotorController{}, "post:CreateMotor")
	web.Router("/motors/edit/:id", &controllers.MotorController{}, "get:EditMotor")
	web.Router("/motors/edit/:id", &controllers.MotorController{}, "post:UpdateMotor")
	web.Router("/motors/delete/:id", &controllers.MotorController{}, "get:DeleteMotor")
	web.Run()
}
