package db

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"os"
// "collection/models"
// 	_ "github.com/lib/pq"
// )

// func InsertData(db *sql.DB) error {
// 	type Data struct {
// 		Users        []models.User        `json:"users"`
// 		Teams        []models.Team        `json:"teams"`
// 		Roles        []models.Role        `json:"roles"`
// 		Branches     []models.Branch      `json:"branches"`
// 		UserRoles    []models.UserRole    `json:"user_roles"`
// 		UserTeams    []models.UserTeam    `json:"user_teams"`
// 		UserTeamRoles []models.UserTeamRole `json:"user_team_roles"`
// 		UserBranches []models.UserBranch  `json:"user_branches"`
// 	}

// 	file, err := os.ReadFile("db.json")
// 	if err != nil {
// 		return err
// 	}


// 	var data Data
// 	err = json.Unmarshal(file, &data)
// 	if err != nil {
// 		return err
// 	}

// 	for _, user := range data.Users {
// 		_, err := db.Exec("INSERT INTO users (id, password, last_login, is_superuser, username, first_name, email, is_active, date_joined, prefix, position) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
// 			user.Id, user.Password, user.LastLogin, user.IsSuperuser, user.Username, user.Name, user.Email, user.IsActive, user.DateJoined, user.Prefix, user.Position)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, team := range data.Teams {
// 		_, err := db.Exec("INSERT INTO teams (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)",
// 			team.Id, team.Name, team.CreatedAt, team.UpdatedAt)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, role := range data.Roles {
// 		_, err := db.Exec("INSERT INTO roles (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)",
// 			role.Id, role.Name, role.CreatedAt, role.UpdatedAt)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, branch := range data.Branches {
// 		_, err := db.Exec("INSERT INTO branches (id, name, branch_code, governate, longitude, latitude, address, flex_value, branch_company) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
// 			branch.Id, branch.Name, branch.BranchCode, branch.Governate, branch.Longitude, branch.Latitude, branch.Address, branch.FlexValue, branch.BranchCompany)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, userRole := range data.UserRoles {
// 		_, err := db.Exec("INSERT INTO user_roles (id, user_id, role_id) VALUES ($1, $2, $3)",
// 			userRole.Id, userRole.UserId, userRole.RoleId)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, userTeam := range data.UserTeams {
// 		_, err := db.Exec("INSERT INTO user_teams (id, user_id, team_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
// 			userTeam.Id, userTeam.UserId, userTeam.TeamId, userTeam.CreatedAt, userTeam.UpdatedAt)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, userTeamRole := range data.UserTeamRoles {
// 		_, err := db.Exec("INSERT INTO user_team_roles (id, user_id, team_id, role_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
// 			userTeamRole.Id, userTeamRole.UserId, userTeamRole.TeamId, userTeamRole.RoleId, userTeamRole.CreatedAt, userTeamRole.UpdatedAt)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, userBranch := range data.UserBranches {
// 		_, err := db.Exec("INSERT INTO user_branches (id, user_id, branch_id) VALUES ($1, $2, $3)",
// 			userBranch.Id, userBranch.UserId, userBranch.BranchId)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
