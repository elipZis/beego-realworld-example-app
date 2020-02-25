package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	//Set up the api namespace
	apiNs := beego.NewNamespace("/api")//It should verify the encrypted request in the production using
	//beego.NSCond(func(ctx *context.Context) bool {
	//	if ua := ctx.Input.Request.UserAgent(); ua != "" {
	//		return true
	//	}
	//	return false
	//}),
	//beego.NSBefore(auth),
	//beego.NSGet("/notallowed", func(ctx *context.Context) {
	//	ctx.Output.Body([]byte("notAllowed"))
	//}),
	//beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
	//beego.NSRouter("/changepassword", &UserController{}),
	//beego.NSNamespace("/shop",
	//	beego.NSBefore(sentry),
	//	beego.NSGet("/:id", func(ctx *context.Context) {
	//		ctx.Output.Body([]byte("notAllowed"))
	//	}),
	//),
	//beego.NSNamespace("/cms",
	//	beego.NSInclude(
	//		&controllers.MainController{},
	//		&controllers.CMSController{},
	//		&controllers.BlockController{},
	//	),
	//),

	//Register the API NS
	beego.AddNamespace(apiNs)
}
