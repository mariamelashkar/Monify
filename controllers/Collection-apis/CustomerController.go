package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq" // Postgres driver
)

type CustomerController struct {
	beego.Controller
}
