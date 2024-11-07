package controllers

import (
	"monify/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetUser() {
	userID := c.Ctx.Input.Param(":id")  
	fmt.Println(userID)
  
	if userID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "User ID is required"}
		c.ServeJSON()
		return
	}

	user, err := models.GetUserById(userID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "User retrieved successfully",
		"user":    user,
	}
	c.ServeJSON()
}

func (c *UserController) GetAllUsers() {

	users, err := models.GetAllUsers()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "All users retrieved successfully",
		"users":   users,
	}
	c.ServeJSON()
}

func (c *UserController) UpdateUser() {
	userID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody
	fmt.Println(userID)

	updatedUser, err := models.UpdateUser(userID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedUser
	c.ServeJSON()
}

func (c *UserController) CreateUser() {
	requestBody := c.Ctx.Input.RequestBody

	id, err := models.CreateUser(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message": "User created successfully",
		"id":      id,
	}
	c.ServeJSON()
}

func (c *UserController) SoftDeleteUser() {
	userID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedUser, err := models.SoftDeleteUser(userID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedUser
	c.ServeJSON()
}
