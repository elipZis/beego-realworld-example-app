package routers

import (
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/controllers/web"
)

func init() {
	beego.Router("/", &web.MainController{})
}
