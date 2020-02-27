package api

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/models"
	"github.com/elipZis/beego-realworld-example-app/repositories"
	"github.com/elipZis/beego-realworld-example-app/services"
)

type AuthController struct {
	beego.Controller
	*services.JwtService
	*repositories.UserRepository
}

// @Import github.com/elipZis/beego-realworld-example-app/routers/middleware
// @Filter /users/login beego.BeforeRouter middleware.IsAuthenticated
// @router /users/login [post]
func (this *AuthController) Login() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	token, err := this.JwtService.GenerateToken(&user)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"User": user, "token": token, "Users": this.UserRepository.FindAll()}
	}
	this.ServeJSON()
}

// @router /users [post]
func (this *AuthController) Register() {
}
