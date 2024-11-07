package controllers

import (
	"monify/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type UserTeamsController struct {
	beego.Controller 
}

func (c *UserTeamsController) GetUserTeam() {
	userTeamID := c.Ctx.Input.Param(":id")
	fmt.Println(userTeamID)

	if userTeamID == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "User Team ID is required"}
		c.ServeJSON()
		return
	}

	userTeam, err := models.GetUserTeamById(userTeamID)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message":    "User Team retrieved successfully",
		"user_team":  userTeam,
	}
	c.ServeJSON()
}

func (c *UserTeamsController) GetAllUserTeams() {
	userTeams, err := models.GetAllUserTeams()
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = map[string]interface{}{
		"message":    "All user teams retrieved successfully",
		"user_teams": userTeams,
	}
	c.ServeJSON()
}

func (c *UserTeamsController) UpdateUserTeam() {
	userTeamID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody
	fmt.Println(userTeamID)

	updatedUserTeam, err := models.UpdateUserTeam(userTeamID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedUserTeam
	c.ServeJSON()
}

func (c *UserTeamsController) CreateUserTeam() {
	requestBody := c.Ctx.Input.RequestBody

	id, err := models.CreateUserTeam(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"message":    "User Team created successfully",
		"id":         id,
	}
	c.ServeJSON()
}

func (c *UserTeamsController) SoftDeleteUserTeam() {
	userTeamID := c.Ctx.Input.Param(":id")
	requestBody := c.Ctx.Input.RequestBody

	updatedUserTeam, err := models.SoftDeleteUserTeam(userTeamID, requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = updatedUserTeam
	c.ServeJSON()
}
