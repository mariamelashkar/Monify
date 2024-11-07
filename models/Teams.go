package models

import (
	"monify/utils"
	"encoding/json"
	"errors"
	"time"
)

type Team struct {
	Id        string     `orm:"column(id);pk;size(70)" json:"id"`
	Name      string     `orm:"column(name);size(200)" json:"name"`
	Region    string     `orm:"column(region);size(255)" json:"region"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);auto_now_add" json:"created_at"`
	CreatedBy string     `orm:"column(created_by)" json:"created_by"`
	UpdatedBy *string    `orm:"column(updated_by);size(70)" json:"updated_by,omitempty"`
	UpdatedAt *time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at,omitempty"`
	Actions   string     `orm:"column(action);size(200);default(created)" json:"actions"`
	Deleted   bool       `orm:"column(deleted);default(false)" json:"deleted"`
}


func (t *Team) TableName() string {
	return "teams"
}


// // ---------- CRUD Operations for Teams Model ----------//
func GetTeam(teamID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
    SELECT id, name, region, created_at, created_by, updated_by, updated_at, action, deleted
    FROM teams
    WHERE id = $1
    `
	params := []interface{}{teamID}

	Result, err := utils.QueryExecuter(db, "team_result", query, params)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no team found with the given ID")
	}

	return Result[0], nil
}

func GetAllTeams() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
    SELECT id, name, region, created_at, created_by, updated_by, updated_at, action, deleted
    FROM teams
    `
	Result, err := utils.QueryExecuter(db, "team_result", query, nil)
	if err != nil {
		return nil, err
	}
	if len(Result) == 0 {
		return nil, errors.New("no teams found")
	}

	return Result, nil
}

func CreateTeam(requestBody []byte) (string, error) {
	var team map[string]interface{}
	if err := json.Unmarshal(requestBody, &team); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if team["name"] == nil || team["region"] == nil {
		return "", errors.New("name and region are required")
	}

	team["created_at"] = utils.CurrentTime
	team["updated_at"] = utils.CurrentTime
	team["action"] = "created"
	team["deleted"] = false

	query := `
    INSERT INTO teams (id, name, region, created_at, created_by, updated_by, updated_at, action, deleted)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id
    `

	params := []interface{}{
		team["id"],
		team["name"],
		team["region"],
		team["created_at"],
		team["created_by"],
		team["updated_by"],
		team["updated_at"],
		team["action"],
		team["deleted"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "team_results", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 && len(Result[0]) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		} else {
			return "", errors.New("failed to retrieve the team ID")
		}
	}

	return "", errors.New("no result returned from the query")
}

func UpdateTeam(teamID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingTeam, err := GetTeam(teamID)
	if err != nil {
		return nil, errors.New("no team found with the given ID")
	}

	var updatedTeam map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedTeam); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	updatedTeam["id"] = existingTeam["id"]
	updatedTeam["created_at"] = existingTeam["created_at"]
	updatedTeam["created_by"] = existingTeam["created_by"]
	updatedTeam["updated_at"] = utils.CurrentTime

	query := `
    UPDATE team
    SET name = $1, region = $2, updated_at = $3, updated_by = $4, action = $5, deleted = $6
    WHERE id = $7
    RETURNING *
    `

	params := []interface{}{
		updatedTeam["name"],
		updatedTeam["region"],
		updatedTeam["updated_at"],
		updatedTeam["updated_by"],
		updatedTeam["action"],
		updatedTeam["deleted"],
		teamID,
	}

	Result, err := utils.QueryExecuter(db, "team_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update the team")
	}

	return Result[0], nil
}

func SoftDeleteTeam(teamID string, requestBody []byte) (map[string]interface{}, error) {
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
    UPDATE teams
    SET deleted = $1, updated_by = $2, updated_at = $3
    WHERE id = $4
    RETURNING *
    `

	params := []interface{}{
		softDeleteRequest.Deleted,
		softDeleteRequest.UpdatedBy,
		softDeleteRequest.UpdatedAt,
		teamID,
	}

	Result, err := utils.QueryExecuter(db, "team_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete the team")
	}

	return Result[0], nil
}
