package models

import (
	"monify/utils"
	"encoding/json"
	"errors"
	"time"
)

type UserTeamHierarchy struct {
	Id          string     `orm:"column(id);pk;size(70)"`
	UserId      *User      `orm:"column(user_id);rel(fk)"`
	TeamId      *Team      `orm:"column(team_id);rel(fk)"`
	HierarchyId *Hierarchy `orm:"column(hierarchy_id);rel(fk)"`
	CreatedAt   time.Time  `orm:"column(created_at);type(datetime);auto_now_add"`
	CreatedBy   string     `orm:"column(created_by)"`
	UpdatedBy   *string    `orm:"column(updated_by);size(70)"`
	UpdatedAt   *time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	Actions     string     `orm:"column(action);size(200);default(created)"`
	Deleted     bool       `orm:"column(deleted);default(false)"`
}

func (utr *UserTeamHierarchy) TableName() string {
	return "user_team_hierarchy"
}

// // ---------- CRUD Operations for user_heirarchy Model ----------//
func GetUserTeamHierarchyById(userTeamHierarchyID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT 
		uth.id AS user_team_hierarchy_id,
		uth.created_at AS user_team_hierarchy_created_at,
		uth.created_by AS user_team_hierarchy_created_by,
		uth.updated_at AS user_team_hierarchy_updated_at,
		uth.updated_by AS user_team_hierarchy_updated_by,
		uth.action AS user_team_hierarchy_action,
		uth.deleted AS user_team_hierarchy_deleted,
		
		u.id AS user_id,
		u.username AS user_name,
		u.name AS user_full_name,
		u.email AS user_email,
		
		t.id AS team_id,
		t.name AS team_name,
		
		h.id AS hierarchy_id,
		h.name AS hierarchy_name,
		h.level AS hierarchy_level
	FROM user_team_hierarchy uth
	LEFT JOIN users u ON uth.user_id = u.id
	LEFT JOIN teams t ON uth.team_id = t.id
	LEFT JOIN hierarchy h ON uth.hierarchy_id = h.id
	WHERE uth.id = $1 AND uth.deleted = false
	`

	params := []interface{}{userTeamHierarchyID}
	Result, err := utils.QueryExecuter(db, "user_team_hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("no user team hierarchy found with the given ID")
	}

	userTeamHierarchyData := map[string]interface{}{
		"id":              Result[0]["user_team_hierarchy_id"],
		"created_at":      Result[0]["user_team_hierarchy_created_at"],
		"created_by":      Result[0]["user_team_hierarchy_created_by"],
		"updated_at":      Result[0]["user_team_hierarchy_updated_at"],
		"updated_by":      Result[0]["user_team_hierarchy_updated_by"],
		"action":          Result[0]["user_team_hierarchy_action"],
		"deleted":         Result[0]["user_team_hierarchy_deleted"],
		"user": map[string]interface{}{
			"id":   Result[0]["user_id"],
			"name": Result[0]["user_name"],
			"full_name": Result[0]["user_full_name"],
			"email": Result[0]["user_email"],
		},
		"team": map[string]interface{}{
			"id":   Result[0]["team_id"],
			"name": Result[0]["team_name"],
		},
		"hierarchy": map[string]interface{}{
			"id":    Result[0]["hierarchy_id"],
			"name":  Result[0]["hierarchy_name"],
			"level": Result[0]["hierarchy_level"],
		},
	}

	return userTeamHierarchyData, nil
}

func GetAllUserTeamHierarchies()([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT 
		uth.id AS user_team_hierarchy_id,
		uth.created_at AS user_team_hierarchy_created_at,
		uth.created_by AS user_team_hierarchy_created_by,
		uth.updated_at AS user_team_hierarchy_updated_at,
		uth.updated_by AS user_team_hierarchy_updated_by,
		uth.action AS user_team_hierarchy_action,
		uth.deleted AS user_team_hierarchy_deleted,
		
		u.id AS user_id,
		u.username AS user_name,
		u.name AS user_full_name,
		u.email AS user_email,
		
		t.id AS team_id,
		t.name AS team_name,
		
		h.id AS hierarchy_id,
		h.name AS hierarchy_name,
		h.level AS hierarchy_level
	FROM user_team_hierarchy uth
	LEFT JOIN users u ON uth.user_id = u.id
	LEFT JOIN teams t ON uth.team_id = t.id
	LEFT JOIN hierarchy h ON uth.hierarchy_id = h.id
	WHERE uth.deleted = false
	`

	Result, err := utils.QueryExecuter(db, "user_team_hierarchy_result", query, nil)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("no user team hierarchies found")
	}

	var userTeamHierarchies []map[string]interface{}
	for _, row := range Result {
		userTeamHierarchyData := map[string]interface{}{
			"id":              row["user_team_hierarchy_id"],
			"created_at":      row["user_team_hierarchy_created_at"],
			"created_by":      row["user_team_hierarchy_created_by"],
			"updated_at":      row["user_team_hierarchy_updated_at"],
			"updated_by":      row["user_team_hierarchy_updated_by"],
			"action":          row["user_team_hierarchy_action"],
			"deleted":         row["user_team_hierarchy_deleted"],
			"user": map[string]interface{}{
				"id":    row["user_id"],
				"name":  row["user_name"],
				"full_name": row["user_full_name"],
				"email": row["user_email"],
			},
			"team": map[string]interface{}{
				"id":   row["team_id"],
				"name": row["team_name"],
			},
			"hierarchy": map[string]interface{}{
				"id":    row["hierarchy_id"],
				"name":  row["hierarchy_name"],
				"level": row["hierarchy_level"],
			},
		}
		userTeamHierarchies = append(userTeamHierarchies, userTeamHierarchyData)
	}

	return userTeamHierarchies, nil
}

func CreateUserTeamHierarchy(requestBody []byte) (string, error) {
	var userTeamHierarchy map[string]interface{}
	if err := json.Unmarshal(requestBody, &userTeamHierarchy); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if userTeamHierarchy["user_id"] == nil || userTeamHierarchy["team_id"] == nil || userTeamHierarchy["hierarchy_id"] == nil {
		return "", errors.New("user_id, team_id, and hierarchy_id are required")
	}

	userTeamHierarchy["created_at"] = utils.CurrentTime
	userTeamHierarchy["updated_at"] = utils.CurrentTime
	userTeamHierarchy["action"] = "created"
	userTeamHierarchy["deleted"] = false

	query := `
        INSERT INTO user_team_hierarchies (
            id, user_id, team_id, hierarchy_id, created_at, created_by, updated_at, updated_by, action, deleted
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING id
    `

	params := []interface{}{
		userTeamHierarchy["id"], userTeamHierarchy["user_id"], userTeamHierarchy["team_id"],
		userTeamHierarchy["hierarchy_id"], userTeamHierarchy["created_at"], userTeamHierarchy["created_by"],
		userTeamHierarchy["updated_at"], userTeamHierarchy["updated_by"], userTeamHierarchy["action"],
		userTeamHierarchy["deleted"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "user_team_hierarchy_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		}
	}

	return "", errors.New("failed to create user team hierarchy")
}
func UpdateUserTeamHierarchy(userTeamHierarchyID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingHierarchy, err := GetUserTeamHierarchyById(userTeamHierarchyID)
	if err != nil {
		return nil, errors.New("no user team hierarchy found with the given ID")
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
        UPDATE user_team_hierarchies
        SET user_id = $1, team_id = $2, hierarchy_id = $3, updated_at = $4, updated_by = $5, action = $6, deleted = $7
        WHERE id = $8
        RETURNING *
    `

	params := []interface{}{
		updatedHierarchy["user_id"], updatedHierarchy["team_id"], updatedHierarchy["hierarchy_id"],
		updatedHierarchy["updated_at"], updatedHierarchy["updated_by"], updatedHierarchy["action"],
		updatedHierarchy["deleted"], userTeamHierarchyID,
	}

	Result, err := utils.QueryExecuter(db, "user_team_hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update user team hierarchy")
	}

	return Result[0], nil
}
func SoftDeleteUserTeamHierarchy(userTeamHierarchyID string, requestBody []byte) (map[string]interface{}, error) {
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
        UPDATE user_team_hierarchies
        SET deleted = $1, updated_by = $2, updated_at = $3
        WHERE id = $4
        RETURNING id, user_id, team_id, hierarchy_id, created_at, updated_at, deleted
    `

	params := []interface{}{
		softDeleteRequest.Deleted, softDeleteRequest.UpdatedBy, softDeleteRequest.UpdatedAt, userTeamHierarchyID,
	}

	Result, err := utils.QueryExecuter(db, "user_team_hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete user team hierarchy")
	}

	return Result[0], nil
}
