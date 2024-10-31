package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq" // Postgres driver
)

type ReallocationController struct {
	beego.Controller
}
