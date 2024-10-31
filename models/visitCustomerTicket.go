package models

import (
	"time"

	// "github.com/beego/beego/v2/client/orm"
)

type VisitCustomerTicket struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	Ticket    *TicketTable `orm:"column(ticket_id);rel(fk);unique"`
	Lat       string       `orm:"column(lat);size(200)"`
	Lang      string       `orm:"column(lang);size(200)"`
	Timestamp time.Time    `orm:"column(timestamp);type(timestamp);auto_now_add"`
	//Follower  *Followers   `orm:"column(follower_id);rel(fk)"`
	Date      time.Time    `orm:"column(date);type(date);auto_now;unique"`
	DeviceId  string       `orm:"column(device_id);size(200)"`
}

func (V *VisitCustomerTicket) TableName() string {
	return "collection_late_customer_visits"
}
