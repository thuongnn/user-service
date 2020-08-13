package models

import "time"

type User struct {
	UserID       int64     `orm:"pk;auto;column(user_id)" json:"user_id"`
	Username     string    `orm:"column(username)" json:"username"`
	Password     string    `orm:"column(password)" json:"password"`
	Salt         string    `orm:"column(salt)" json:"-"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool      `orm:"column(deleted)" json:"deleted"`
}

// TableName ...
func (u *User) TableName() string {
	return "user"
}