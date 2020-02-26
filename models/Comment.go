package models

type Comment struct {
	*HasTimestamps

	Id     int
	Body   string `orm:"type(text)"`
	Author *User  `orm:"rel(fk)"`
}
