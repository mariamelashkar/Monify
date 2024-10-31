package models

import (
	"time"

	// "github.com/beego/beego/v2/client/orm"
)

type ActionResult struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	Action          *FollowerAction `orm:"column(action_id);rel(one);unique"`
	Sentimental     int             `orm:"column(sentimental);type(int);default(0)"`
	Desc            string          `orm:"column(desc_);type(text)"`
	Result          *ResultChoose   `orm:"column(result);rel(fk)"`
	CollectedAmount float64         `orm:"column(collected_amount);type(float);default(0)"`
	NextFollowDate  time.Time       `orm:"column(next_follow_date);type(date)"`
	AddDate         time.Time       `orm:"column(add_date);type(timestamp);auto_now_add"`
	UpdateAt        time.Time       `orm:"column(update_at);type(timestamp);auto_now"`
	UpdateUser      *User   `orm:"column(update_user);rel(fk)"`
}

func (A *ActionResult) TableName() string {
	return "action_result"
}

