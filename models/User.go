package models

type User struct {
	*HasTimestamps

	Id       int
	Email    string `orm:"unique"`
	Password string
	Token    string `orm:"null;type(text)"`
	Username string `orm:"null"`
	Bio      string `orm:"null;type(text)"`
	Image    string `orm:"null;type(text)"`
}
