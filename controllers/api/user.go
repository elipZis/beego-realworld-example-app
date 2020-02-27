package api

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("User", c.User)
	c.Mapping("Update", c.Update)
}

// @router /user [get]
func (this *UserController) User() {

}

// @router /user [put]
func (this *UserController) Update() {
}
