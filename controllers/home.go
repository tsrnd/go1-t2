package controllers

import (
	"goweb2/controllers/base/frontend"
)

type HomeController struct {
	frontend.FrontendLayout
}

func (c *HomeController) Get() {
	c.TplName = "frontend/templates/home/content.tpl"
	c.Content()
}

func (c *HomeController) Content() {
    c.Data["Title"] = "Golang's best project in Asiantech | Technology Service Provider | Software Company"
}
