package models

import (
	// "collection/utils"
	// "encoding/json"
	// "errors"
	"time"
)

type Branch struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	Name          string     `orm:"column(name);size(255);null" json:"name"`
	BranchCode    int        `orm:"column(branch_code);unique" json:"branch_code"`
	Governate     string     `orm:"column(governate);size(200);null" json:"governate"`
	Longitude     float64    `orm:"column(longitude);null" json:"longitude"`
	Latitude      float64    `orm:"column(latitude);null" json:"latitude"`
	Address       string     `orm:"column(address);type(text);null" json:"address"`
	FlexValue     string     `orm:"column(flex_value);size(150);null" json:"flex_value"`
	BranchCompany string     `orm:"column(branch_company);size(200);null" json:"branch_company"`
	CreatedAt     time.Time  `orm:"column(created_at);type(datetime);auto_now_add"`
	CreatedBy     string     `orm:"column(created_by);size(70)"`
	UpdatedBy *string     `orm:"column(updated_by)"`
	UpdatedAt     *time.Time `orm:"column(updated_at);auto_now;type(datetime);null"`
	Actions       string     `orm:"column(action);size(200);default(created)"`
	Deleted       bool       `orm:"column(deleted);default(false)"`
}

func (U *Branch) TableName() string {
	return "branches"
}



// // ---------- CRUD Operations for Branch Model ----------//
