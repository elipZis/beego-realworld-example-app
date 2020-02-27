package main

import (
	"github.com/astaxie/beego"
	_ "github.com/elipZis/beego-realworld-example-app/repositories"
	_ "github.com/elipZis/beego-realworld-example-app/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	beego.Run()
}
