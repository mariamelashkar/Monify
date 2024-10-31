package models

import (
	"collection/utils"
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	Id          string     `orm:"column(id);pk;size(70)"`
	Password    string     `orm:"column(password);size(128);"`
	LastLogin   time.Time  `orm:"column(last_login);type(datetime2);"`
	IsSuperuser bool       `orm:"column(is_superuser);"`
	Username    string     `orm:"column(username);size(150);unique"`
	Name        string     `orm:"column(name);size(150);"`
	Email       string     `orm:"column(email);size(254);"`
	IsActive    bool       `orm:"column(is_active);"`
	DateJoined  time.Time  `orm:"column(date_joined);type(datetime2);"`
	Prefix      string     `orm:"column(prefix);size(50);"`
	Position    string     `orm:"column(position);size(100);"`
	RoleId      *Role      `orm:"rel(fk);column(role_id)"`
	CreatedAt   time.Time  `orm:"column(created_at);auto_now_add;type(datetime)"`
	CreatedBy   string     `orm:"column(created_by):size(70)"`
	UpdatedBy   *string    `orm:"column(updated_by);size(70)"`
	UpdatedAt   *time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	Actions     string     `orm:"column(action);size(200);default(created)"`
	Deleted     bool       `orm:"column(deleted);default(false)"`
}

func (u *User) TableName() string {
	return "users"
}

// // ---------- CRUD Operations for user Model ----------//
func GetUserById(userID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
SELECT 
    u.id,
    u.password, 
    u.last_login, 
    u.is_superuser, 
    u.username, 
    u.name, 
    u.email, 
    u.is_active, 
    u.date_joined, 
    u.prefix, 
    u.position, 
    u.created_at, 
    u.created_by, 
    u.updated_at, 
    u.updated_by, 
    u.action, 
    u.deleted,
    ur.name AS role_name, 
    ur.created_at AS role_created_at,
    ur.created_by AS role_created_by,
    ur.updated_at AS role_updated_at,
    ur.action AS role_action,
    ur.deleted AS role_deleted
FROM users u
FULL OUTER JOIN user_roles ur ON u.role_id = ur.id
WHERE u.id = $1
  AND u.deleted = false
  AND ur.deleted = false
  `
  params := []interface{}{userID}

  Result, err := utils.QueryExecuter(db, "user_result", query, params)
  if err != nil {
	  return nil, err
  }

  if len(Result) == 0 {
	  return nil, errors.New("no user found with the given ID")
  }
  userData := map[string]interface{}{
	"id":           Result[0]["user_id"],
	"password":     Result[0]["password"],
	"last_login":   Result[0]["last_login"],
	"is_superuser": Result[0]["is_superuser"],
	"username":     Result[0]["username"],
	"name":         Result[0]["user_name"],
	"email":        Result[0]["email"],
	"is_active":    Result[0]["is_active"],
	"date_joined":  Result[0]["date_joined"],
	"prefix":       Result[0]["prefix"],
	"position":     Result[0]["position"],
	"created_at":   Result[0]["user_created_at"],
	"created_by":   Result[0]["user_created_by"],
	"updated_at":   Result[0]["user_updated_at"],
	"updated_by":   Result[0]["user_updated_by"],
	"action":       Result[0]["user_action"],
	"deleted":      Result[0]["user_deleted"],
	"role": map[string]interface{}{ 
		"id":         Result[0]["role_id"],
		"name":       Result[0]["role_name"],
		"created_at": Result[0]["role_created_at"],
		"created_by": Result[0]["role_created_by"],
		"updated_at": Result[0]["role_updated_at"],
		"action":     Result[0]["role_action"],
		"deleted":    Result[0]["role_deleted"],
	},
}

return userData, nil
}
func GetAllUsers() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
SELECT 
    u.id AS user_id,
    u.password, 
    u.last_login, 
    u.is_superuser, 
    u.username, 
    u.name AS user_name, 
    u.email, 
    u.role_id,
    u.is_active, 
    u.date_joined, 
    u.prefix, 
    u.position, 
    u.created_at AS user_created_at, 
    u.created_by AS user_created_by, 
    u.updated_at AS user_updated_at, 
    u.updated_by AS user_updated_by, 
    u.action AS user_action, 
    u.deleted AS user_deleted,
    ur.id AS role_id,
    ur.name AS role_name, 
    ur.created_at AS role_created_at,
    ur.created_by AS role_created_by,
    ur.updated_at AS role_updated_at,
    ur.action AS role_action,
    ur.deleted AS role_deleted
FROM users u
LEFT JOIN user_roles ur ON u.role_id = ur.id
WHERE u.deleted = false AND (ur.deleted = false OR ur.deleted IS NULL)
`

	Result, err := utils.QueryExecuter(db, "user_result", query, nil)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("no users found")
	}

	var users []map[string]interface{}
	for _, row := range Result {
		userData := map[string]interface{}{
			"id":           row["user_id"],
			"password":     row["password"],
			"last_login":   row["last_login"],
			"is_superuser": row["is_superuser"],
			"username":     row["username"],
			"name":         row["user_name"],
			"email":        row["email"],
			"is_active":    row["is_active"],
			"date_joined":  row["date_joined"],
			"prefix":       row["prefix"],
			"position":     row["position"],
			"created_at":   row["user_created_at"],
			"created_by":   row["user_created_by"],
			"updated_at":   row["user_updated_at"],
			"updated_by":   row["user_updated_by"],
			"action":       row["user_action"],
			"deleted":      row["user_deleted"],
			"role": map[string]interface{}{ 
				"id":         row["role_id"],
				"name":       row["role_name"],
				"created_at": row["role_created_at"],
				"created_by": row["role_created_by"],
				"updated_at": row["role_updated_at"],
				"action":     row["role_action"],
				"deleted":    row["role_deleted"],
			},
		}
		users = append(users, userData)
	}

	return users, nil
}

func CreateUser(requestBody []byte) (string, error) {
	var user map[string]interface{}
	if err := json.Unmarshal(requestBody, &user); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if user["username"] == nil || user["email"] == nil {
		return "", errors.New("username and email are required")
	}

	user["created_at"] = utils.CurrentTime
	user["updated_at"] = utils.CurrentTime
	user["action"] = "created"
	user["deleted"] = false

	query := `
        INSERT INTO users (
            id, password, last_login, is_superuser, username, name, email, role_id, 
            is_active, date_joined, prefix, position, created_at, created_by, updated_at, 
            updated_by, action, deleted
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
        RETURNING id
    `

	params := []interface{}{
		user["id"], user["password"], user["last_login"], user["is_superuser"], user["username"],
		user["name"], user["email"], user["role_id"], user["is_active"], user["date_joined"],
		user["prefix"], user["position"], user["created_at"], user["created_by"],
		user["updated_at"], user["updated_by"], user["action"], user["deleted"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "user_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		}
	}

	return "", errors.New("failed to create user")
}
func UpdateUser(userID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingUser, err := GetUserById(userID)
	if err != nil {
		return nil, errors.New("no user found with the given ID")
	}

	var updatedUser map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedUser); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	updatedUser["id"] = existingUser["id"]
	updatedUser["created_at"] = existingUser["created_at"]
	updatedUser["created_by"] = existingUser["created_by"]
	updatedUser["updated_at"] = utils.CurrentTime

	query := `
        UPDATE users
        SET password = $1, last_login = $2, is_superuser = $3, username = $4, name = $5, 
            email = $6, role_id = $7, is_active = $8, prefix = $9, position = $10, updated_at = $11, 
            updated_by = $12, action = $13, deleted = $14
        WHERE id = $15
        RETURNING *
    `

	params := []interface{}{
		updatedUser["password"], updatedUser["last_login"], updatedUser["is_superuser"], updatedUser["username"],
		updatedUser["name"], updatedUser["email"], updatedUser["role_id"], updatedUser["is_active"],
		updatedUser["prefix"], updatedUser["position"], updatedUser["updated_at"], updatedUser["updated_by"],
		updatedUser["action"], updatedUser["deleted"], userID,
	}

	Result, err := utils.QueryExecuter(db, "user_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update user")
	}

	return Result[0], nil
}
func SoftDeleteUser(userID string, requestBody []byte) (map[string]interface{}, error) {
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
        UPDATE users
        SET deleted = $1, updated_by = $2, updated_at = $3
        WHERE id = $4
        RETURNING id, username, email, created_at, updated_at, deleted
    `

	params := []interface{}{
		softDeleteRequest.Deleted, softDeleteRequest.UpdatedBy, softDeleteRequest.UpdatedAt, userID,
	}

	Result, err := utils.QueryExecuter(db, "user_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete user")
	}

	return Result[0], nil
}
