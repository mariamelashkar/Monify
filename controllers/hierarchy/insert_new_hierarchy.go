package hierarchy

import (
	"monify/models"
	"monify/utils"
	"encoding/json"
	"log"
	beego "github.com/beego/beego/v2/server/web"
)

type HierarchyController struct {
	beego.Controller
}

func (c *HierarchyController) CreateHierarchy() {
	var req struct { 
		Hierarchy   models.Hierarchy `json:"hierarchy"`
		FormerLevel *int             `json:"former_level,omitempty"`
		NextLevel   *int             `json:"next_level,omitempty"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid JSON input"}
		c.ServeJSON()
		return
	}

	if req.Hierarchy.Name == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Hierarchy name is required"}
		c.ServeJSON()
		return
	}

	req.Hierarchy.CreatedAt = utils.CurrentTime
	req.Hierarchy.UpdatedAt = &utils.CurrentTime

	db, err := utils.GetDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to connect to the database"}
		c.ServeJSON()
		return
	}
	defer db.Close()

	// var formerHierarchyID, nextHierarchyID int

	if req.FormerLevel != nil {
		var formerHierarchy models.Hierarchy
		formerQuery := `
            SELECT id FROM hierarchy WHERE level = $1
        `
		err = db.QueryRow(formerQuery, *req.FormerLevel).Scan(&formerHierarchy.Id)
		if err != nil {
			log.Printf("Failed to retrieve hierarchy with level %d: %v", *req.FormerLevel, err)
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = map[string]string{"error": "Former hierarchy not found at the provided level"}
			c.ServeJSON()
			return
		}
		// formerHierarchyID = formerHierarchy.Id
	}

	if req.NextLevel != nil {
		var nextHierarchy models.Hierarchy
		nextQuery := `
            SELECT id FROM hierarchy WHERE level = $1
        `
		err = db.QueryRow(nextQuery, *req.NextLevel).Scan(&nextHierarchy.Id)
		if err != nil {
			log.Printf("Failed to retrieve hierarchy with level %d: %v", *req.NextLevel, err)
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = map[string]string{"error": "Next hierarchy not found at the provided level"}
			c.ServeJSON()
			return
		}
		// nextHierarchyID = nextHierarchy.Id
	}
	var existingHierarchyId string
	checkDuplicateQuery := `
        SELECT id FROM hierarchy WHERE name = $1 AND delted != true
    `
	err = db.QueryRow(checkDuplicateQuery, req.Hierarchy.Name).Scan(&existingHierarchyId)
	if err == nil {
		log.Printf("Duplicate hierarchy name found: %s", req.Hierarchy.Name)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Hierarchy name already exists"}
		c.ServeJSON()
		return
	}
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Transaction failed"}
		c.ServeJSON()
		return
	}

	if req.NextLevel != nil {
		shiftQuery := `
            UPDATE hierarchy
            SET level = level + 1
            WHERE level >= $1
        `
		_, err = tx.Exec(shiftQuery, *req.NextLevel)
		if err != nil {
			log.Printf("Failed to shift hierarchy levels: %v", err)
			tx.Rollback()
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Failed to shift hierarchy levels"}
			c.ServeJSON()
			return
		}

		req.Hierarchy.Level = *req.NextLevel
	} else if req.FormerLevel != nil {
		req.Hierarchy.Level = *req.FormerLevel + 1
	}

	insertQuery := `
        INSERT INTO hierarchy (id,name, level, created_at, updated_at)
        VALUES ($1, $2, $3, $4,$5)
        RETURNING id
    `
	var newHierarchyID int
	err = tx.QueryRow(insertQuery,req.Hierarchy.Id, req.Hierarchy.Name, req.Hierarchy.Level, req.Hierarchy.CreatedAt, req.Hierarchy.UpdatedAt).Scan(&newHierarchyID)
	if err != nil {
		log.Printf("Failed to insert new hierarchy: %v", err)
		tx.Rollback()
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to insert new hierarchy"}
		c.ServeJSON()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Transaction commit failed"}
		c.ServeJSON()
		return
	}

	updatedHierarchies := []models.Hierarchy{}
	updatedQuery := `
        SELECT id, name, level FROM hierarchy ORDER BY level
    `
	rows, err := db.Query(updatedQuery)
	if err != nil {
		log.Printf("Failed to retrieve updated hierarchies: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to retrieve updated hierarchies"}
		c.ServeJSON()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var h models.Hierarchy
		err = rows.Scan(&h.Id, &h.Name, &h.Level)
		if err != nil {
			log.Printf("Failed to scan hierarchy: %v", err)
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Failed to scan updated hierarchies"}
			c.ServeJSON()
			return
		}
		updatedHierarchies = append(updatedHierarchies, h)
	}

	response := map[string]interface{}{
		"message":             "Hierarchy created successfully",
		"new_hierarchy_id":    newHierarchyID,
		"new_hierarchy_name":  newHierarchyID,
		"updated_hierarchies": updatedHierarchies,
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = response
	c.ServeJSON()
}
