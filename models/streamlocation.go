package models

import (
	"time"
)

type StreamLocation struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	Date     time.Time `orm:"column(date);type(date);auto_now_add"`
	DeviceId string    `orm:"column(device_id);size(200)"`
	Lat      string    `orm:"column(lat);size(200)"`
	Lang     string    `orm:"column(lang);size(200)"`
}

func (S *StreamLocation) TableName() string {
	return "collection_streamlocation"
}

