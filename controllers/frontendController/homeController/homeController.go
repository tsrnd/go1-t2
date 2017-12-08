package homeController

import (
	"goweb2/controllers/frontendController/baseController"
)

type HomeController struct {
	baseController.FrontendLayout
}

func (c *HomeController) Index() {
	c.TplName = "frontend/templates/home/content.tpl"
	c.Content()
}

func (c *HomeController) Content() {
    c.Data["Title"] = "Golang's best project in Asiantech | Technology Service Provider | Software Company"
}
