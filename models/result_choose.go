package models

import (
	// "github.com/beego/beego/v2/client/orm"
)

type ResultChoose struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	ResultName string `orm:"column(result_name);size(255);unique"`
	Category   int    `orm:"column(category);type(int);default(0)"`
}

func (R *ResultChoose) TableName() string {
	return "result_choose"
}

// func init() {
// 	orm.RegisterModel(new(ResultChoose))
// }
