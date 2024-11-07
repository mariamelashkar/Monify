package filters

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"monify/utils"
		beego "github.com/beego/beego/v2/server/web"


)
	type CollectionController struct {
	beego.Controller
}


type HierarchyRequest struct {
	CurrentLevel int `json:"current_level"`
}
func (c *CollectionController) GetUsersAtNextHierarchyLevel() ([]map[string]interface{}, error) {

	db, err := utils.GetDBConnection()
	if err != nil {
		log.Println("Error connecting to database:", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error connecting to database"}
		c.ServeJSON()
		return nil, err
	}
    var req HierarchyRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		log.Printf("Error parsing request: %v", err)
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid request data"}
		c.ServeJSON()
		return nil, err
	}
	
	nextLevel := req.CurrentLevel + 1
	
	usersData, err := GetUsersAtNextLevel(db,nextLevel)
	if err != nil {
		log.Printf("Error retrieving users: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Database retrieval failure"}
		c.ServeJSON()
		return nil, err
	}
	
	if len(usersData) == 0 {
		log.Printf("No users found at the next hierarchy level")
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = map[string]string{"error": "No users found at the next level"}
		c.ServeJSON()
		return nil, fmt.Errorf("no users found at level %d", nextLevel)
	}
	
	response := map[string]interface{}{
		"next_level_users": usersData,
	}
	
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = response
	c.ServeJSON()
	
	return usersData, nil
}	