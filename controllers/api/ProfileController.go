package api

import (
	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}

func (c *ProfileController) URLMapping() {
	c.Mapping("Profile", c.Profile)
	c.Mapping("Follow", c.Follow)
	c.Mapping("Unfollow", c.UnFollow)
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
