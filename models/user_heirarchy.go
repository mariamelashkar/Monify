package models

import (
	"collection/utils"
	"encoding/json"
	"errors"
	"time"
	"fmt"

)

type UserHierarchy struct {
    Id          string     `orm:"column(id);pk;size(70)" json:"id"`
    UserId      *User      `orm:"rel(fk);column(user_id)" json:"user_id"`
    HierarchyId *Hierarchy `orm:"rel(fk);column(hierarchy_id)" json:"hierarchy_id"`
    CreatedAt   time.Time  `orm:"column(created_at);type(datetime);auto_now_add" json:"created_at"`
    CreatedBy   string     `orm:"column(created_by)" json:"created_by"`
    UpdatedBy   *string    `orm:"column(updated_by);size(70)" json:"updated_by,omitempty"`
    UpdatedAt   *time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at,omitempty"`
    Actions     string     `orm:"column(action);size(200);default(created)" json:"actions"`
    Deleted     bool       `orm:"column(deleted);default(false)" json:"deleted"`
}

func (ug *UserHierarchy) TableName() string {
	return "user_hierarchy"
}


// // ---------- CRUD Operations for user_heirarchy Model ----------//
func GetUserHierarchyById(userID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT 
    uh.id, 
    uh.user_id AS user_id, 
    uh.hierarchy_id AS hierarchy_id, 
    uh.created_at, 
    uh.created_by, 
    uh.updated_at, 
    uh.updated_by, 
    uh.actions, 
    uh.deleted,
    p.id AS parent_id, 
    p.user_id AS parent_user_id
FROM user_hierarchy uh
LEFT JOIN user_hierarchy p ON uh.parent_id = p.id
WHERE uh.id = $1 OR uh.parent_id = $1

	`

	params := []interface{}{userID}

	Result, err := utils.QueryExecuter(db, "user_hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no user hierarchy found with the given ID")
	}

	return Result[0], nil
}
func GetAllUserHierarchySnaps() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT 
		id, user_id, hierarchy_id, created_at, created_by, updated_at, updated_by, actions, deleted
	FROM user_hierarchy
	`

	Result, err := utils.QueryExecuter(db, "user_hierarchy_result", query, nil)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no user hierarchies found")
	}

	return Result, nil
}
func CreateUserHierarchy(requestBody []byte) (string, error) {
	var userHierarchyData UserHierarchy
	if err := json.Unmarshal(requestBody, &userHierarchyData); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if userHierarchyData.UserId == nil || userHierarchyData.HierarchyId == nil || userHierarchyData.CreatedBy == "" {
		return "", errors.New("user ID, hierarchy ID, and created by fields are required")
	}

	userID := userHierarchyData.UserId.Id 
	hierarchyID := userHierarchyData.HierarchyId.Id 

	_, err := GetUserHierarchyById(userID)
	if err != nil {
		return "", err 
	}

	_, err = GetHierarchy(hierarchyID)
	if err != nil {
		return "", err 
	}

	now := time.Now()
	userHierarchyData.CreatedAt = now
	userHierarchyData.UpdatedAt = &now
	userHierarchyData.Actions = "created"
	userHierarchyData.Deleted = false

	query := `
		INSERT INTO user_hierarchy (
			user_id, hierarchy_id, created_at, created_by, updated_at, updated_by, actions, deleted
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	params := []interface{}{
		userHierarchyData.UserId.Id, 
		userHierarchyData.HierarchyId.Id, 
		userHierarchyData.CreatedAt,
		userHierarchyData.CreatedBy,
		userHierarchyData.UpdatedAt,
		userHierarchyData.UpdatedBy,
		userHierarchyData.Actions,
		userHierarchyData.Deleted,
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "user_hierarchy_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 && len(Result[0]) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		} else {
			return "", errors.New("failed to retrieve the user hierarchy ID")
		}
	}

	return "", errors.New("no result returned from the query")
}
func UpdateUserHierarchy( userHierarchyID string, requestBody []byte) (map[string]interface{}, error) {
	fmt.Println(userHierarchyID)
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingUserHierarchy, err := GetUserHierarchyById(userHierarchyID)
	if err != nil {
		return nil, errors.New("no user hierarchy found with the given ID")
	}

	fmt.Println(existingUserHierarchy)

	var updatedUserHierarchy map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedUserHierarchy); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	if _, err := GetUserById(updatedUserHierarchy["user_id"].(string)); err != nil {
		return nil, errors.New("invalid user_id provided")
	}

	if _, err := GetHierarchy(updatedUserHierarchy["hierarchy_id"].(string)); err != nil {
		return nil, errors.New("invalid hierarchy_id provided")
	}

	updatedUserHierarchy["id"] = existingUserHierarchy["id"]
	updatedUserHierarchy["created_at"] = existingUserHierarchy["created_at"]
	updatedUserHierarchy["created_by"] = existingUserHierarchy["created_by"]
	updatedUserHierarchy["updated_at"] = utils.CurrentTime

	if updatedUserHierarchy["updated_by"] == nil {
		return nil, errors.New("the 'UpdatedBy' field is required for updating")
	}

	if updatedUserHierarchy["actions"] == "" {
		return nil, errors.New("the 'Actions' field is required for updating")
	}

	query := `
        UPDATE user_hierarchy
        SET 
            user_id = $1, 
            hierarchy_id = $2, 
            updated_at = $3, 
            updated_by = $4, 
            actions = $5, 
            deleted = $6
        WHERE id = $7
        RETURNING *
    `

	params := []interface{}{
		updatedUserHierarchy["user_id"],
		updatedUserHierarchy["hierarchy_id"],
		updatedUserHierarchy["updated_at"],
		updatedUserHierarchy["updated_by"],
		updatedUserHierarchy["actions"],
		updatedUserHierarchy["deleted"],
		userHierarchyID,
	}

	Result, err := utils.QueryExecuter(db, "user_hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update the user hierarchy")
	}

	return Result[0], nil
}
func SoftDeleteUserHierarchy(userHierarchyID string, requestBody []byte) (map[string]interface{}, error) {
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
        UPDATE user_hierarchy
        SET deleted = $1, updated_by = $2, updated_at = $3
        WHERE id = $4
        RETURNING id, user_id, hierarchy_id, created_at, created_by,
                  updated_at, updated_by, actions, deleted
    `

	params := []interface{}{
		softDeleteRequest.Deleted,
		softDeleteRequest.UpdatedBy,
		softDeleteRequest.UpdatedAt,
		userHierarchyID,
	}

	Result, err := utils.QueryExecuter(db, "user_hierarchy_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete the user hierarchy")
	}

	return Result[0], nil
}
