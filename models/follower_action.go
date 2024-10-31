package models

import (
	"time"

	// "github.com/beego/beego/v2/client/orm"
)

type FollowerAction struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	Ticket     *TicketTable `orm:"column(ticket_id);rel(fk);unique"`
	//Follower   *Followers   `orm:"column(follower_id);rel(fk)"`
	Action     *Actions     `orm:"column(action_id);rel(fk)"`
	Lat        string       `orm:"column(lat);size(150);null"`
	Lang       string       `orm:"column(lang);size(150);null"`
	AddDate    time.Time    `orm:"column(add_date);type(date);auto_now_add"`
	UpdateDate time.Time    `orm:"column(update_date);type(date);auto_now"`
	Date       time.Time    `orm:"column(date);type(date);auto_now_add;unique"`
}

func (F *FollowerAction) TableName() string {
	return "follower_action"
}

// func init() {
// 	orm.RegisterModel(new(FollowerAction))
// }

