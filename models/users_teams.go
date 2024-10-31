package models

import (
	"collection/utils"
	"encoding/json"
	"errors"
	"time"
)

type UserTeam struct {
	Id        string     `orm:"column(id);pk;size(70)"`
	UserId    *User      `orm:"rel(fk);column(user_id)"`
	TeamId    *Team      `orm:"rel(fk);column(team_id)"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);auto_now_add"`
	CreatedBy string     `orm:"column(created_by);size(70)"`
	UpdatedBy *string    `orm:"column(updated_by)"`
	UpdatedAt *time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	Actions   string     `orm:"column(action);size(200);default(created)"`
	Deleted   bool       `orm:"column(deleted);default(false)"`
}

func (u *UserTeam) TableName() string {
	return "users_teams"
}

// // ---------- CRUD Operations for user_teams Model ----------//
func GetUserTeamById(userTeamID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
		SELECT 
			ut.id AS user_team_id,
			ut.created_at AS user_team_created_at,
			ut.created_by AS user_team_created_by,
			ut.updated_at AS user_team_updated_at,
			ut.updated_by AS user_team_updated_by,
			ut.action AS user_team_action,
			ut.deleted AS user_team_deleted,
			
			u.id AS user_id,
			u.last_login AS user_last_login,
			u.is_superuser AS user_is_superuser,
			u.username AS user_name,
			u.name AS user_full_name,
			u.email AS user_email,
			u.is_active AS user_is_active,
			u.date_joined AS user_date_joined,
			u.prefix AS user_prefix,
			u.position AS user_position,
			u.role_id AS user_role_id,
			u.created_at AS user_created_at,
			u.created_by AS user_created_by,
			u.updated_by AS user_updated_by,
			u.updated_at AS user_updated_at,
			u.action AS user_action,
			u.deleted AS user_deleted,
			
			t.id AS team_id,
			t.name AS team_name,
			t.region AS team_region,
			t.created_at AS team_created_at,
			t.created_by AS team_created_by,
			t.updated_by AS team_updated_by,
			t.updated_at AS team_updated_at,
			t.action AS team_action,
			t.deleted AS team_deleted
		FROM users_teams ut
		LEFT JOIN users u ON ut.user_id = u.id
		LEFT JOIN teams t ON ut.team_id = t.id
		WHERE u.id = $1 AND ut.deleted = false;
		`

	params := []interface{}{userTeamID}

	Result, err := utils.QueryExecuter(db, "user_team_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("no user team found with the given ID")
	}

	userTeamData := map[string]interface{}{
		"id":         Result[0]["user_team_id"],
		"created_at": Result[0]["user_team_created_at"],
		"created_by": Result[0]["user_team_created_by"],
		"updated_at": Result[0]["user_team_updated_at"],
		"updated_by": Result[0]["user_team_updated_by"],
		"action":     Result[0]["user_team_action"],
		"deleted":    Result[0]["user_team_deleted"],
		"user": map[string]interface{}{
			"id":           Result[0]["user_id"],
			"last_login":   Result[0]["user_last_login"],
			"is_superuser": Result[0]["user_is_superuser"],
			"username":     Result[0]["user_name"],
			"full_name":    Result[0]["user_full_name"],
			"email":        Result[0]["user_email"],
			"is_active":    Result[0]["user_is_active"],
			"date_joined":  Result[0]["user_date_joined"],
			"prefix":       Result[0]["user_prefix"],
			"position":     Result[0]["user_position"],
			"role_id":      Result[0]["user_role_id"],
			"created_at":   Result[0]["user_created_at"],
			"created_by":   Result[0]["user_created_by"],
			"updated_by":   Result[0]["user_updated_by"],
			"updated_at":   Result[0]["user_updated_at"],
			"action":       Result[0]["user_action"],
			"deleted":      Result[0]["user_deleted"],
		},
		"team": map[string]interface{}{
			"id":         Result[0]["team_id"],
			"name":       Result[0]["team_name"],
			"region":     Result[0]["team_region"],
			"created_at": Result[0]["team_created_at"],
			"created_by": Result[0]["team_created_by"],
			"updated_by": Result[0]["team_updated_by"],
			"updated_at": Result[0]["team_updated_at"],
			"action":     Result[0]["team_action"],
			"deleted":    Result[0]["team_deleted"],
		},
	}

	return userTeamData, nil

}
func GetAllUserTeams() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT 
		ut.id AS user_team_id,
		ut.created_at AS user_team_created_at,
		ut.created_by AS user_team_created_by,
		ut.updated_at AS user_team_updated_at,
		ut.updated_by AS user_team_updated_by,
		ut.action AS user_team_action,
		ut.deleted AS user_team_deleted,
		
		u.id AS user_id,
		u.password AS user_password,
		u.last_login AS user_last_login,
		u.is_superuser AS user_is_superuser,
		u.username AS user_name,
		u.name AS user_full_name,
		u.email AS user_email,
		u.is_active AS user_is_active,
		u.date_joined AS user_date_joined,
		u.prefix AS user_prefix,
		u.position AS user_position,
		u.role_id AS user_role_id,
		u.created_at AS user_created_at,
		u.created_by AS user_created_by,
		u.updated_by AS user_updated_by,
		u.updated_at AS user_updated_at,
		u.action AS user_action,
		u.deleted AS user_deleted,
		
		t.id AS team_id,
		t.name AS team_name,
		t.region AS team_region,
		t.created_at AS team_created_at,
		t.created_by AS team_created_by,
		t.updated_by AS team_updated_by,
		t.updated_at AS team_updated_at,
		t.action AS team_action,
		t.deleted AS team_deleted
	FROM users_teams ut
	LEFT JOIN users u ON ut.user_id = u.id
	LEFT JOIN teams t ON ut.team_id = t.id
	WHERE ut.deleted = false
	`

	Result, err := utils.QueryExecuter(db, "user_team_result", query, nil)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("no user teams found")
	}

	var userTeams []map[string]interface{}
	for _, row := range Result {
		userTeamData := map[string]interface{}{
			"id":              row["user_team_id"],
			"created_at":     row["user_team_created_at"],
			"created_by":     row["user_team_created_by"],
			"updated_at":     row["user_team_updated_at"],
			"updated_by":     row["user_team_updated_by"],
			"action":         row["user_team_action"],
			"deleted":        row["user_team_deleted"],
			"user": map[string]interface{}{
				"id":           row["user_id"],
				"password":     row["user_password"], // Sensitive data - consider if you want to include this
				"last_login":   row["user_last_login"],
				"is_superuser": row["user_is_superuser"],
				"username":     row["user_name"],
				"full_name":    row["user_full_name"],
				"email":        row["user_email"],
				"is_active":    row["user_is_active"],
				"date_joined":  row["user_date_joined"],
				"prefix":       row["user_prefix"],
				"position":     row["user_position"],
				"role_id":      row["user_role_id"],
				"created_at":   row["user_created_at"],
				"created_by":   row["user_created_by"],
				"updated_by":   row["user_updated_by"],
				"updated_at":   row["user_updated_at"],
				"action":       row["user_action"],
				"deleted":      row["user_deleted"],
			},
			"team": map[string]interface{}{
				"id":          row["team_id"],
				"name":        row["team_name"],
				"region":      row["team_region"],
				"created_at":  row["team_created_at"],
				"created_by":  row["team_created_by"],
				"updated_by":  row["team_updated_by"],
				"updated_at":  row["team_updated_at"],
				"action":      row["team_action"],
				"deleted":     row["team_deleted"],
			},
		}
		userTeams = append(userTeams, userTeamData)
	}

	return userTeams, nil
}


func CreateUserTeam(requestBody []byte) (string, error) {
	var userTeam map[string]interface{}
	if err := json.Unmarshal(requestBody, &userTeam); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if userTeam["user_id"] == nil || userTeam["team_id"] == nil {
		return "", errors.New("user_id and team_id are required")
	}

	userTeam["created_at"] = utils.CurrentTime
	userTeam["updated_at"] = utils.CurrentTime
	userTeam["action"] = "created"
	userTeam["deleted"] = false

	query := `
		INSERT INTO user_teams (
			id, user_id, team_id, created_at, created_by, updated_at, updated_by, action, deleted
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	params := []interface{}{
		userTeam["id"], userTeam["user_id"], userTeam["team_id"],
		userTeam["created_at"], userTeam["created_by"],
		userTeam["updated_at"], userTeam["updated_by"],
		userTeam["action"], userTeam["deleted"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "user_team_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		}
	}

	return "", errors.New("failed to create user team")
}
func UpdateUserTeam(userTeamID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingUserTeam, err := GetUserTeamById(userTeamID)
	if err != nil {
		return nil, errors.New("no user team found with the given ID")
	}

	var updatedUserTeam map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedUserTeam); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	updatedUserTeam["id"] = existingUserTeam["id"]
	updatedUserTeam["created_at"] = existingUserTeam["created_at"]
	updatedUserTeam["created_by"] = existingUserTeam["created_by"]
	updatedUserTeam["updated_at"] = utils.CurrentTime

	query := `
		UPDATE user_teams
		SET user_id = $1, team_id = $2, updated_at = $3, updated_by = $4, action = $5, deleted = $6
		WHERE id = $7
		RETURNING *
	`

	params := []interface{}{
		updatedUserTeam["user_id"], updatedUserTeam["team_id"],
		updatedUserTeam["updated_at"], updatedUserTeam["updated_by"],
		updatedUserTeam["action"], updatedUserTeam["deleted"], userTeamID,
	}

	Result, err := utils.QueryExecuter(db, "user_team_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update user team")
	}

	return Result[0], nil
}
func SoftDeleteUserTeam(userTeamID string, requestBody []byte) (map[string]interface{}, error) {
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
		UPDATE user_teams
		SET deleted = $1, updated_by = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, user_id, team_id, created_at, updated_at, deleted
	`

	params := []interface{}{
		softDeleteRequest.Deleted, softDeleteRequest.UpdatedBy,
		softDeleteRequest.UpdatedAt, userTeamID,
	}

	Result, err := utils.QueryExecuter(db, "user_team_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete user team")
	}

	return Result[0], nil
}
