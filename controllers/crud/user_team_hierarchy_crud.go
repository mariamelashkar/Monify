package controllers

import (
	"monify/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type UserTeamHierarchyController struct {
	beego.Controller
}

func (c *UserTeamHierarchyController) GetUserTeamHierarchy() {
	hierarchyID := c.Ctx.Input.Param(":id")
	fmt.Println(hierarchyID)

	if hierarchyID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Hierarchy ID is required"}
		c.ServeJSON()
		return
	}

	hierarchy, err := models.GetUserTeamHierarchyById(hierarchyID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "User team hierarchy retrieved successfully",
		"hierarchy":    hierarchy,
	}
	c.ServeJSON()
}

func (c *UserTeamHierarchyController) GetAllUserTeamHierarchies() {
	hierarchies, err := models.GetAllUserTeamHierarchies()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "All user team hierarchies retrieved successfully",
		"hierarchies":   hierarchies,
	}
	c.ServeJSON()
}

func (c *UserTeamHierarchyController) UpdateUserTeamHierarchy() {
	hierarchyID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody
	fmt.Println(hierarchyID)

	updatedHierarchy, err := models.UpdateUserTeamHierarchy(hierarchyID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedHierarchy
	c.ServeJSON()
}

func (c *UserTeamHierarchyController) CreateUserTeamHierarchy() {
	requestBody := c.Ctx.Input.RequestBody

	id, err := models.CreateUserTeamHierarchy(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message": "User team hierarchy created successfully",
		"id":      id,
	}
	c.ServeJSON()
}

func (c *UserTeamHierarchyController) SoftDeleteUserTeamHierarchy() {
	hierarchyID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedHierarchy, err := models.SoftDeleteUserTeamHierarchy(hierarchyID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedHierarchy
	c.ServeJSON()
}
