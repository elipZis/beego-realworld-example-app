package repositories

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/elipZis/beego-realworld-example-app/models"
)

type UserRepository struct {
}

func (repo *UserRepository) FindAll() []*models.User {
	o := orm.NewOrm()
	var users []*models.User
	_, err := o.QueryTable(models.User{}).All(&users)
	if err != nil {
		fmt.Println(err)
	}
	return users
}
