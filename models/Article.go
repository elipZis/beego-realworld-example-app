package models

type Article struct {
	*HasTimestamps

	Id             int
	Slug           string `orm:"unique"`
	Title          string
	Description    string   `orm:"null;type(text)"`
	Body           string   `orm:"null;type(text)"`
	TagList        []string `orm:"null"`
	Favorited      bool     `orm:"default(false)"`
	FavoritesCount int      `orm:"null"`
	Author         *User    `orm:"rel(fk)"`
}
