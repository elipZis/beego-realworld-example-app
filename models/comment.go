package models

import "github.com/elipZis/beego-realworld-example-app/models/traits"

type Comment struct {
	Id     int    `json:"-"`
	Body   string `orm:"type(text)"`
	Author *User  `orm:"rel(fk)"`

	traits.HasTimestamps
}
