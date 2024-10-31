package hierarchy

import (
	"collection/models"
	"collection/utils"
	"encoding/json"
	"log"
)

func (c *HierarchyController) DeleteHierarchy() {
	var req struct {
		Id string `json:"id"` 
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid JSON input"}
		c.ServeJSON()
		return
	}

	if req.Id == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Valid hierarchy ID is required"}
		c.ServeJSON()
		return
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to connect to the database"}
		c.ServeJSON()
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Transaction failed"}
		c.ServeJSON()
		return
	}

	var existingHierarchy models.Hierarchy
	checkQuery := `
        SELECT id, level, name, deleted FROM hierarchy WHERE id = $1
    `
	err = tx.QueryRow(checkQuery, req.Id).Scan(&existingHierarchy.Id, &existingHierarchy.Level, &existingHierarchy.Name, &existingHierarchy.Deleted)
	if err != nil {
		log.Printf("Hierarchy not found with ID %s: %v", req.Id, err)
		tx.Rollback()
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Hierarchy not found"}
		c.ServeJSON()
		return
	}

	if existingHierarchy.Deleted {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Hierarchy already deleted"}
		c.ServeJSON()
		return
	}

	updateQuery := `
        UPDATE hierarchy SET deleted = true WHERE id = $1
    `
	_, err = tx.Exec(updateQuery, req.Id)
	if err != nil {
		log.Printf("Failed to soft delete hierarchy with ID %s: %v", req.Id, err)
		tx.Rollback()
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to delete hierarchy"}
		c.ServeJSON()
		return
	}

	var followingLevels []models.Hierarchy
	followingQuery := `
        SELECT id, level FROM hierarchy WHERE level > $1 AND deleted = false
    `
	rows, err := tx.Query(followingQuery, existingHierarchy.Level)
	if err != nil {
		log.Printf("Failed to retrieve following levels after deletion: %v", err)
		tx.Rollback()
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to retrieve following levels"}
		c.ServeJSON()
		return
	}
	defer rows.Close()

	for rows.Next() {
		var h models.Hierarchy
		if err := rows.Scan(&h.Id, &h.Level); err != nil {
			log.Printf("Failed to scan following level hierarchy: %v", err)
			tx.Rollback()
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Failed to scan following levels"}
			c.ServeJSON()
			return
		}
		followingLevels = append(followingLevels, h)
	}

	for _, levelHierarchy := range followingLevels {
		incrementQuery := `
            UPDATE hierarchy SET level = level - 1 WHERE id = $1
        `
		_, err = tx.Exec(incrementQuery, levelHierarchy.Id)
		if err != nil {
			log.Printf("Failed to update following hierarchy level with ID %s: %v", levelHierarchy.Id, err)
			tx.Rollback()
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Failed to update following hierarchy level"}
			c.ServeJSON()
			return
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Transaction commit failed"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(204) 
	c.Data["json"] = map[string]string{"message": "Hierarchy soft deleted and following levels updated successfully"}
	c.ServeJSON()
}
