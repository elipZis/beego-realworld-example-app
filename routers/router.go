package routers

import (
	"github.com/elipZis/beego-realworld-example-app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
