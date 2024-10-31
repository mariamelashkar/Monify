package controllers

import (
	"collection/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type HierarchyController struct {
	beego.Controller
}

func (c *HierarchyController) GetHierarchy() {
	hierarchyID := c.Ctx.Input.Param(":id")
	fmt.Println(hierarchyID)

	if hierarchyID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Hierarchy ID is required"}
		c.ServeJSON()
		return
	}

	hierarchy, err := models.GetHierarchy(hierarchyID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message":   "Hierarchy retrieved successfully",
		"hierarchy": hierarchy,
	}
	c.ServeJSON()
}

func (c *HierarchyController) GetAllHierarchies() {

	hierarchies, err := models.GetAllHierarchies()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message":     "All hierarchies retrieved successfully",
		"hierarchies": hierarchies,
	}
	c.ServeJSON()
}

func (c *HierarchyController) CreateHierarchy() {
	requestBody := c.Ctx.Input.RequestBody

	id, err := models.CreateHierarchy(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message": "Hierarchy created successfully",
		"id":      id,
	}
	c.ServeJSON()
}

func (c *HierarchyController) UpdateHierarchy() {
	hierarchyID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody
	fmt.Println(hierarchyID)

	updatedHierarchy, err := models.UpdateHierarchy(hierarchyID, requestBody)
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

func (c *HierarchyController) SoftDeleteHierarchy() {
	hierarchyID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedHierarchy, err := models.SoftDeleteHierarchy(hierarchyID, requestBody)
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
