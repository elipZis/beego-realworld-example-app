package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Article), new(Comment), new(User))
}
