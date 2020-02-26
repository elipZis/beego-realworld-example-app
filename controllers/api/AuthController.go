package api

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/models"
)

type AuthController struct {
	beego.Controller
}

// @router /users/login [post]
func (this *AuthController) Login() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	//objectid := models.AddOne(ob)
	this.Data["json"] = map[string]interface{}{"User": user}
	this.ServeJSON()
}

// @router /users [post]
func (this *AuthController) Register() {
}
