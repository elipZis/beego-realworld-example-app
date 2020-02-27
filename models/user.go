package models

import "github.com/elipZis/beego-realworld-example-app/models/traits"

type User struct {
	Id       int    `json:"-"`
	Email    string `orm:"unique"`
	Password string `json:"-"`
	Token    string `orm:"null;type(text)"json:",omitempty"`
	Username string `orm:"null"json:",omitempty"`
	Bio      string `orm:"null;type(text)"json:",omitempty"`
	Image    string `orm:"null;type(text)"json:",omitempty"`

	traits.HasTimestamps
}
