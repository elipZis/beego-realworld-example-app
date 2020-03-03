package models

import (
	"github.com/elipZis/beego-realworld-example-app/models/traits"
)

type User struct {
	Id       int64  `json:"-"`
	Email    string `orm:"unique" form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `json:"-" form:"password" valid:"Required; MinSize(6); MaxSize(100)"`
	Token    string `orm:"null;type(text)" json:",omitempty"`
	Username string `orm:"null" json:",omitempty" form:"username" valid:"MaxSize(100)"`
	Bio      string `orm:"null;type(text)" json:",omitempty"`
	Image    string `orm:"null;type(text)" json:",omitempty"`

	traits.HasTimestamps
}
