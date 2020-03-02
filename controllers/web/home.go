package web

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

// @router / [get]
func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "pages/home.tpl"
}
