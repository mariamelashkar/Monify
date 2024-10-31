package controllers

import (
	"collection/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
    "encoding/json"
)
type LateCustomersSnapController struct {
	beego.Controller
}

func (c *LateCustomersSnapController) GetLateCustomerSnap() {
    CustomerID := c.Ctx.Input.Param(":id")  
    fmt.Println(CustomerID)
  
    if CustomerID == "" {
        c.Ctx.Output.SetStatus(400)
        c.Data["json"] = map[string]string{"error": "Snap ID is required"}
        c.ServeJSON()
        return
    }


    lateCustomerSnap, err := models.GetLateCustomersSnap(CustomerID)
    if err != nil {
        c.Ctx.Output.SetStatus(404)
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    lateLoanSnap, err := models.GetLateLoansSnap(CustomerID)
    if err != nil {
        c.Ctx.Output.SetStatus(404)
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Ctx.Output.SetStatus(200)
    c.Data["json"] = map[string]interface{}{
        "message":   "Late customer snap retrieved successfully",
        "Customer":  lateCustomerSnap,
        "Loans":lateLoanSnap,
    }
    c.ServeJSON()
}

func (c *LateCustomersSnapController) GetAllLateCustomersSnap() {
	var filter models.LateCustomersSnapFilter
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &filter); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid filter parameters"}
		c.ServeJSON()
		return
	}

	lateCustomersSnap, err := models.GetAllLateCustomersSnap(&filter)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message":   "Late customer snap retrieved successfully",
		"Customer":  lateCustomersSnap,
	}
	c.ServeJSON()
}

func (c *LateCustomersSnapController) UpdateLateCustomersSnap() {
    requestBody := c.Ctx.Input.RequestBody
        lateCustomersSnapID := c.Ctx.Input.Param(":id")  
    fmt.Println(lateCustomersSnapID)

   updatedSnap, err := models.UpdateLateCustomersSnap(lateCustomersSnapID, requestBody)
    if err != nil {
        c.Ctx.Output.SetStatus(400)
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Ctx.Output.SetStatus(200)
    c.Data["json"] = updatedSnap
    c.ServeJSON()
}
func (c *LateCustomersSnapController) CreateLateCustomersSnap() {
    requestBody := c.Ctx.Input.RequestBody

    id, err := models.CreateLateCustomersSnap(requestBody)
    if err != nil {
        c.Ctx.Output.SetStatus(400)
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Ctx.Output.SetStatus(201)
    c.Data["json"] = map[string]interface{}{
        "message": "Late customer snap created successfully",
        "id":      id,
    }
    c.ServeJSON()
}

func (c *LateCustomersSnapController) SoftDeleteLateCustomer() {
    lateCustomerID := c.Ctx.Input.Param(":id")
    requestBody := c.Ctx.Input.RequestBody    

    updatedCustomer, err := models.SoftDeleteLateCustomer(lateCustomerID, requestBody)
    if err != nil {
        c.Ctx.Output.SetStatus(400)
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Ctx.Output.SetStatus(200)
    c.Data["json"] = updatedCustomer
    c.ServeJSON()
}
