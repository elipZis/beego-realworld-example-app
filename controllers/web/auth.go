package web

import (
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/models"
	"github.com/elipZis/beego-realworld-example-app/services"
)

type AuthController struct {
	beego.Controller
	*services.FormService
}

// @router /login [get]
func (this *AuthController) GetLogin() {
	this.TplName = "pages/auth/login.tpl"
}

// @router /login [post]
func (this *AuthController) Login() {
	this.TplName = "pages/auth/login.tpl"
	user := models.User{}
	if err := this.ParseForm(&user); err != nil {
		beego.Error(err)
	}
	beego.Info(user)
	this.Redirect("/", 303)
}

// @router /register [get]
func (this *AuthController) GetRegister() {
	this.TplName = "pages/auth/register.tpl"
	this.Data["User"] = models.User{}
}

// @router /register [post]
func (this *AuthController) Register() {
	this.TplName = "pages/auth/register.tpl"

	user, errors := this.FormService.ParseAndValidate(this.Input(), &models.User{})
	// No Errors -> Go Home!
	if errors == nil {
		this.Redirect("/", 303)
	}
	this.Data["Errors"] = errors
	this.Data["User"] = user
}
