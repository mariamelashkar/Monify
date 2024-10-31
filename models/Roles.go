package models

import (
	"collection/utils"
	"encoding/json"
	"errors"
	"time"
)

type Role struct {
	Id        string     `orm:"column(id);pk;size(70)" json:"id"`
	Name      string     `orm:"column(name);size(200)" json:"name"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);auto_now_add" json:"created_at"`
	CreatedBy string     `orm:"column(created_by)" json:"created_by"`
	UpdatedBy *string    `orm:"column(updated_by);size(70)" json:"updated_by"`
	UpdatedAt *time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
	Actions   string     `orm:"column(action);size(200);default(created)" json:"actions"`
	Deleted   bool       `orm:"column(deleted);default(false)" json:"deleted"`
}


func (r *Role) TableName() string {
	return "user_roles"
}


// //('Managerial'),
// //('collection'),

// // ---------- CRUD Operations for Roles Model ----------//
func GetRole(roleID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
  SELECT 
    u.id,u.password, u.last_login, u.is_superuser, u.username, u.name, u.email, u.is_active, 
    u.date_joined, u.prefix, u.position, u.created_at, u.created_by, u.updated_at, u.updated_by, u.action, u.deleted,
	ur.name As role_name , ur.created_at,ur.created_by,ur.updated_by,ur.updated_at,ur.action,ur.deleted
  FROM users u
	  FULL OUTER JOIN user_roles ur
ON u.role_id = ur.id
WHERE u.role_id=$1
	`

	params := []interface{}{roleID}

	Result, err := utils.QueryExecuter(db, "role_result", query, params)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no role found with the given ID")
	}

	return Result[0], nil
}

func GetAllRoles() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT id, name, created_at, created_by, updated_by, updated_at, action, deleted
	FROM roles
	`

	Result, err := utils.QueryExecuter(db, "role_result", query, nil)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no roles found")
	}

	return Result, nil
}

func CreateRole(requestBody []byte) (string, error) {
	var role map[string]interface{}
	if err := json.Unmarshal(requestBody, &role); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if role["name"] == nil {
		return "", errors.New("role name is required")
	}

	role["created_at"] = utils.CurrentTime
	role["updated_at"] = utils.CurrentTime
	role["actions"] = "created"
	role["deleted"] = false

	query := `
	INSERT INTO roles (id, name, created_at, created_by, updated_by, updated_at, action, deleted)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id
	`

	params := []interface{}{
		role["id"],
		role["name"],
		role["created_at"],
		role["created_by"],
		role["updated_by"],
		role["updated_at"],
		role["actions"],
		role["deleted"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "role_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 && len(Result[0]) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		} else {
			return "", errors.New("failed to retrieve the role ID")
		}
	}

	return "", errors.New("no result returned from the query")
}

func UpdateRole(roleID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingRole, err := GetRole(roleID)
	if err != nil {
		return nil, err
	}

	var updatedRole map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedRole); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	updatedRole["id"] = existingRole["id"]
	updatedRole["created_at"] = existingRole["created_at"]
	updatedRole["created_by"] = existingRole["created_by"]
	updatedRole["updated_at"] = utils.CurrentTime

	if updatedRole["updated_by"] == nil {
		return nil, errors.New("the 'UpdatedBy' field is required for updating")
	}

	if updatedRole["name"] == nil {
		return nil, errors.New("the 'Name' field is required for updating")
	}

	query := `
	UPDATE roles
	SET name = $1, updated_by = $2, updated_at = $3, action = $4, deleted = $5
	WHERE id = $6
	RETURNING *
	`

	params := []interface{}{
		updatedRole["name"],
		updatedRole["updated_by"],
		updatedRole["updated_at"],
		updatedRole["actions"],
		updatedRole["deleted"],
		roleID,
	}

	Result, err := utils.QueryExecuter(db, "role_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update the role")
	}

	return Result[0], nil
}

func SoftDeleteRole(roleID string, requestBody []byte) (map[string]interface{}, error) {
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
	UPDATE roles
	SET deleted = $1, updated_by = $2, updated_at = $3
	WHERE id = $4
	RETURNING id, name, created_at, created_by, updated_by, updated_at, action, deleted
	`

	params := []interface{}{
		softDeleteRequest.Deleted,
		softDeleteRequest.UpdatedBy,
		softDeleteRequest.UpdatedAt,
		roleID,
	}

	Result, err := utils.QueryExecuter(db, "role_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete the role")
	}

	return Result[0], nil
}
