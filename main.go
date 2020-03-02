package main

import (
	"github.com/Masterminds/sprig"
	"github.com/astaxie/beego"
	_ "github.com/elipZis/beego-realworld-example-app/repositories"
	_ "github.com/elipZis/beego-realworld-example-app/routers"
	_ "github.com/joho/godotenv/autoload"
	"github.com/leekchan/gtf"
)

func main() {
	// Register all GTF functions -- Must be before Run()
	for k, v := range gtf.GtfFuncMap {
		beego.AddFuncMap(k, v)
	}
	// Register all sprig functions -- Must be before Run()
	for k, v := range sprig.FuncMap() {
		beego.AddFuncMap(k, v)
	}

	beego.Run()
}
