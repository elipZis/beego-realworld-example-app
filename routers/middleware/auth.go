package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/elipZis/beego-realworld-example-app/services"
	"net/http"
)

func IsAuthenticated(ctx *context.Context) {
	var jwtService services.JwtService
	claims, err := jwtService.ValidateTokenFromContext(ctx)
	if err != nil {
		ctx.Abort(http.StatusUnauthorized, "Unauthorized!")
	}
	//ctx.SetCookie("token", tokenString,
	//	Expires: expirationTime,
	//})
	beego.Info(claims)
}
