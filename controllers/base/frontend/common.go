package frontend

import (
   "github.com/astaxie/beego"
)

type FrontendLayout struct {
   beego.Controller
}

func (c *FrontendLayout) Prepare() {
   c.Layout = "frontend/layouts/layout.tpl"
   c.LayoutSections = map[string]string{}
   c.LayoutSections["Header"] = "frontend/includes/header.tpl"
   c.LayoutSections["Slider"] = "frontend/includes/slider.tpl"
   c.LayoutSections["Footer"] = "frontend/includes/footer.tpl"
}