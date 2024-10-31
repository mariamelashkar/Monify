package controllers

import (
	"collection/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UserHierarchyController struct {
	beego.Controller
}

func (c *UserHierarchyController) GetUserHierarchy() {
	userHierarchyID := c.Ctx.Input.Param(":id")

	if userHierarchyID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "User ID is required"}
		c.ServeJSON()
		return
	}

	userHierarchy, err := models.GetUserHierarchyById(userHierarchyID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = userHierarchy
	c.ServeJSON()
}

func (c *UserHierarchyController) GetAllUserHierarchies() {

	userHierarchies, err := models.GetAllHierarchies()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message":         "User Hierarchies retrieved successfully",
		"user_hierarchies": userHierarchies,
	}
	c.ServeJSON()
}

func (c *UserHierarchyController) CreateUserHierarchy() {
	requestBody := c.Ctx.Input.RequestBody

	createdID, err := models.CreateUserHierarchy(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message": "User Hierarchy created successfully",
		"id":      createdID,
	}
	c.ServeJSON()
}

func (c *UserHierarchyController) UpdateUserHierarchy() {
	userHierarchyID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	if userHierarchyID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "User Hierarchy ID is required"}
		c.ServeJSON()
		return
	}

	updatedUserHierarchy, err := models.UpdateUserHierarchy(userHierarchyID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedUserHierarchy
	c.ServeJSON()
}

func (c *UserHierarchyController) DeleteUserHierarchy() {
	userHierarchyID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	if userHierarchyID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "User Hierarchy ID is required"}
		c.ServeJSON()
		return
	}

	updatedUserHierarchy, err := models.UpdateUserHierarchy(userHierarchyID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "User Hierarchy deleted successfully",
		"user_hierarchy": updatedUserHierarchy,
	}
	c.ServeJSON()
}
