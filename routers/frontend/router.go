package frontend

import (
	"goweb2/controllers/frontendController/homeController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &homeController.HomeController{}, "get:Index")
}
