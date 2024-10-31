package models

import (
	"time"

	// "github.com/beego/beego/v2/client/orm"
)

type CurrentFollower struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	Ticket     *TicketTable `orm:"column(ticket_id);rel(fk)"`
//	FollowUser *Followers   `orm:"column(follow_user_id);rel(fk)"`
	AddDate    time.Time    `orm:"column(add_date);type(date);auto_now_add"`
	AddUser    *User `orm:"column(add_user);rel(fk)"`
}

func (C *CurrentFollower) TableName() string {
	return "current_follower"
}

