package traits

import "time"

type HasTimestamps struct {
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"json:"-"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"json:"-"`
}
