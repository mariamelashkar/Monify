package controllers

import (
	"monify/models"

	beego "github.com/beego/beego/v2/server/web"
)


type RoleController struct {
	beego.Controller
}

func (c *RoleController) GetRole() {
	roleID := c.Ctx.Input.Param(":id")
	// fmt.Println(roleID)

	if roleID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Rule ID is required"}
		c.ServeJSON()
		return
	}

	rule, err := models.GetRole(roleID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "Rule retrieved successfully",
		"rule":    rule,
	}
	c.ServeJSON()
}

func (c *RoleController) GetAllRoles() {

	roles, err := models.GetAllRoles()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "All rules retrieved successfully",
		"rules":   roles,
	}
	c.ServeJSON()
}

func (c *RoleController) UpdateRole() {
	roleID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody
	// fmt.Println(roleID)

	updatedRule, err := models.UpdateRole(roleID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedRule
	c.ServeJSON()
}

func (c *RoleController) CreateRule() {
	requestBody := c.Ctx.Input.RequestBody

	id, err := models.CreateRole(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message": "Rule created successfully",
		"id":      id,
	}
	c.ServeJSON()
}

func (c *RoleController) SoftDeleteRule() {
	roleID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedRule, err := models.SoftDeleteRole(roleID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedRule
	c.ServeJSON()
}
