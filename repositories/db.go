package repositories

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/elipZis/beego-realworld-example-app/models"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", beego.AppConfig.String("db::driver"), beego.AppConfig.String("db::username")+":"+beego.AppConfig.String("db::password")+"@/"+beego.AppConfig.String("db::database")+"?charset=utf8")

	// TODO: DEBUGGING
	err := orm.RunSyncdb("default", true, false)
	if err != nil {
		fmt.Println(err)
	}
}
