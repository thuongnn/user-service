package models

import "time"

type Book struct {
	ID           int64     `orm:"pk;auto;column(id)" json:"id"`
	OwnerID      int64     `orm:"column(owner_id)" json:"owner_id"`
	Name         string    `orm:"column(name)" json:"name"`
	Description  string    `orm:"column(description)" json:"description"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool      `orm:"column(deleted)" json:"deleted"`
}
