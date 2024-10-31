package controllers

import (
	"collection/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type TeamController struct {
	beego.Controller
}

func (c *TeamController) GetTeam() {
	teamID := c.Ctx.Input.Param(":id")
	fmt.Println(teamID)

	if teamID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Team ID is required"}
		c.ServeJSON()
		return
	}

	team, err := models.GetTeam(teamID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "Team retrieved successfully",
		"team":    team,
	}
	c.ServeJSON()
}

func (c *TeamController) GetAllTeams() {

	teams, err := models.GetAllTeams()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message": "All teams retrieved successfully",
		"teams":   teams,
	}
	c.ServeJSON()
}

func (c *TeamController) CreateTeam() {
	requestBody := c.Ctx.Input.RequestBody

	id, err := models.CreateTeam(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message": "Team created successfully",
		"id":      id,
	}
	c.ServeJSON()
}

func (c *TeamController) UpdateTeam() {
	teamID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedTeam, err := models.UpdateTeam(teamID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedTeam
	c.ServeJSON()
}

func (c *TeamController) SoftDeleteTeam() {
	teamID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedTeam, err := models.SoftDeleteTeam(teamID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedTeam
	c.ServeJSON()
}
