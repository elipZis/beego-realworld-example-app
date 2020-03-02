package routers

import (
	"github.com/astaxie/beego"
	"github.com/elipZis/beego-realworld-example-app/controllers/api"
	"github.com/elipZis/beego-realworld-example-app/controllers/web"
	"github.com/elipZis/beego-realworld-example-app/routers/filter"
)

func init() {
	apis()
	webs()
}

// All realworld.io Backend API relevant routes
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

// All frontend full-stack urls
// Adapting the frontend realworld.io requirements
func webs() {
	//Set up the web routes
	beego.InsertFilter("/", beego.BeforeExec, filter.AddJwtHeader)
	beego.Include(&web.AuthController{})
	beego.Include(&web.HomeController{})
}
