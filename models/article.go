package models

import "github.com/elipZis/beego-realworld-example-app/models/traits"

type Article struct {
	Id             int    `json:"-"`
	Slug           string `orm:"unique"`
	Title          string
	Description    string `orm:"null;type(text)"json:",omitempty"`
	Body           string `orm:"null;type(text)"json:",omitempty"`
	TagList        string `orm:"null;type(json)"json:",omitempty"`
	Favorited      bool   `orm:"default(false)"`
	FavoritesCount int    `orm:"null"json:",omitempty"`
	Author         *User  `orm:"rel(fk)"`

	traits.HasTimestamps
}
