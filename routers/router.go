package routers

import (
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/controllers/api"
	"github.com/elipZis/beego-realworld-example-app/controllers/web"
)

func init() {
	apis()
	webs()
}

func apis() {
	//Set up the api namespace
	apiNs := beego.NewNamespace("/api",
		beego.NSInclude(
			&api.ArticleController{},
			&api.AuthController{},
			&api.CommentController{},
			&api.ProfileController{},
			&api.UserController{},
		),
	)

	//Register the API NS
	beego.AddNamespace(apiNs)
}

func webs() {
	//Set up the web routes
	beego.Router("/", &web.MainController{})
}
