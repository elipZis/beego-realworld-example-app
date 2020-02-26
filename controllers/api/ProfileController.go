package api

import (
	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

// @router /profiles/:username [get]
func (this *ProfileController) Profile(username string) {

}

// @router /profiles/:username/follow [post]
func (this *ProfileController) Follow(username string) {
}

// @router /profiles/:username/follow [delete]
func (this *ProfileController) UnFollow(username string) {
}
