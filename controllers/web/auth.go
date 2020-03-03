package web

import (
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/models"
	"github.com/elipZis/beego-realworld-example-app/services"
)

type AuthController struct {
	beego.Controller
	*services.FormService
	*services.FormError
}

// @router /login [get,post]
func (this *AuthController) Login() {
	this.TplName = "pages/auth/login.tpl"
	user := models.User{}
	this.Data["User"] = user

	if this.Ctx.Input.Is("POST") {
		user, errors := this.FormService.ParseAndValidate(this.Input(), &user)
		if errors == nil {
			// TODO: Login
			beego.Info(this.GetSession("user"))
			this.SetSession("user", user)

			//
			this.Redirect("/", 303)
		}
		this.Data["Errors"] = errors
	}
}

// @router /register [get,post]
func (this *AuthController) Register() {
	this.TplName = "pages/auth/register.tpl"
	user := models.User{}
	this.Data["User"] = user

	if this.Ctx.Input.Is("POST") {
		user, errors := this.FormService.ParseAndValidate(this.Input(), &user)

		beego.Info(this.GetSession("user"))
		this.SetSession("user", user)

		// No Errors -> Go Home!
		if errors == nil {
			this.Redirect("/", 303)
		}
		this.Data["Errors"] = errors
	}
}
