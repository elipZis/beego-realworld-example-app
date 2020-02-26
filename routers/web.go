package routers

import (
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/controllers/web"
)

func initWeb() {
	//Set up the web routes
	beego.Router("/", &web.MainController{})
}
