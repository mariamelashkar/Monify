package models

import (
	"collection/utils"
	"encoding/json"
	"errors"
	"time"
)
type Hierarchy struct {
	Id        string     `orm:"column(id);pk;size(70)" json:"id"`
	Level     int        `orm:"column(level);type(int)" json:"level"`
	Name      string     `orm:"column(name);size(200)" json:"name"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);auto_now_add" json:"created_at"`
	CreatedBy string     `orm:"column(created_by)" json:"created_by"`
	UpdatedBy *string    `orm:"column(updated_by);size(70)" json:"updated_by,omitempty"`
	UpdatedAt *time.Time `orm:"column(updated_at);auto_now;type(datetime);null" json:"updated_at,omitempty"`
	Actions   string     `orm:"column(action);size(200);default(created)" json:"action"`
	Deleted   bool       `orm:"column(deleted);default(false)" json:"deleted"`
}


func (r *Hierarchy) TableName() string {
	return "hierarchy"
}

// //('collection manager'),
// // 		('area supervisor'),
// //         ('teamleader'),
// //         ('supervisor'),
// //         ('collection officer')

// // ---------- CRUD Operations for Hierarchy Model ----------//
func GetHierarchy(hierarchyID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT id, level, name, created_at, created_by, updated_at, updated_by, action, deleted
	FROM hierarchy
	WHERE id = $1
	`

	params := []interface{}{hierarchyID}
	Result, err := utils.QueryExecuter(db, "hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no hierarchy found with the given ID")
	}
	return Result[0], nil
}
func GetAllHierarchies() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT id, level, name, created_at, created_by, updated_at, updated_by, action, deleted
	FROM hierarchy
	`

	Result, err := utils.QueryExecuter(db, "hierarchy_result", query, nil)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no hierarchies found")
	}
	return Result, nil
}
func CreateHierarchy( requestBody []byte) (string, error) {
	var hierarchy map[string]interface{}
	if err := json.Unmarshal(requestBody, &hierarchy); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if hierarchy["level"] == nil || hierarchy["name"] == nil {
		return "", errors.New("level and name are required")
	}

	hierarchy["created_at"] = utils.CurrentTime
	hierarchy["updated_at"] = utils.CurrentTime
	hierarchy["actions"] = "created"
	hierarchy["deleted"] = false

	query := `
		INSERT INTO hierarchy (id, level, name, created_at, created_by, updated_at, updated_by, action, deleted)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	params := []interface{}{
		hierarchy["id"],
		hierarchy["level"],
		hierarchy["name"],
		hierarchy["created_at"],
		hierarchy["created_by"],
		hierarchy["updated_at"],
		hierarchy["updated_by"],
		hierarchy["actions"],
		hierarchy["deleted"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "hierarchy_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 && len(Result[0]) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		}
	}
	return "", errors.New("failed to create hierarchy")
}
func UpdateHierarchy(hierarchyID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingHierarchy, err := GetHierarchy(hierarchyID)
	if err != nil {
		return nil, errors.New("no hierarchy found with the given ID")
	}

	var updatedHierarchy map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedHierarchy); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	updatedHierarchy["id"] = existingHierarchy["id"]
	updatedHierarchy["created_at"] = existingHierarchy["created_at"]
	updatedHierarchy["created_by"] = existingHierarchy["created_by"]
	updatedHierarchy["updated_at"] = utils.CurrentTime

	query := `
		UPDATE hierarchy
		SET level = $1, name = $2, updated_at = $3, updated_by = $4, action = $5, deleted = $6
		WHERE id = $7
		RETURNING *
	`

	params := []interface{}{
		updatedHierarchy["level"],
		updatedHierarchy["name"],
		updatedHierarchy["updated_at"],
		updatedHierarchy["updated_by"],
		updatedHierarchy["action"],
		updatedHierarchy["deleted"],
		hierarchyID,
	}

	Result, err := utils.QueryExecuter(db, "hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update hierarchy")
	}

	return Result[0], nil
}
func SoftDeleteHierarchy(hierarchyID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var softDeleteRequest struct {
		Deleted   bool       `json:"deleted"`
		UpdatedBy string     `json:"updated_by"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	if err := json.Unmarshal(requestBody, &softDeleteRequest); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	softDeleteRequest.UpdatedAt = &utils.CurrentTime

	query := `
		UPDATE hierarchy
		SET deleted = $1, updated_by = $2, updated_at = $3
		WHERE id = $4
		RETURNING *
	`

	params := []interface{}{
		softDeleteRequest.Deleted,
		softDeleteRequest.UpdatedBy,
		softDeleteRequest.UpdatedAt,
		hierarchyID,
	}

	Result, err := utils.QueryExecuter(db, "hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete the hierarchy")
	}

	return Result[0], nil
}
