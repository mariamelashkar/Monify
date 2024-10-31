package models

import (
	//"github.com/beego/beego/v2/client/orm"
)

type Actions struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	ActionName string `orm:"column(action_name);size(255)"`
	ActionType string `orm:"column(action_type);size(200)"`
}

func (A *Actions) TableName() string {
	return "actions"
}
