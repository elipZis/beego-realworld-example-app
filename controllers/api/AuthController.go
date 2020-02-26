package api

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
}

// @router /users/login [post]
func (this *AuthController) Login() {

}

// @router /users [post]
func (this *AuthController) Register() {
}
