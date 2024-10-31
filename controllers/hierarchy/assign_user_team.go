package hierarchy

// import (
//     "collection/utils"
//     "encoding/json"
//     beego "github.com/beego/beego/v2/server/web"
//     "log"
//     "time"
// )

// type UserTeamController struct {
//     beego.Controller
// }

// func (c *UserTeamController) AssignUserToTeam(){
//     var req struct {
//         UserID int `json:"user_id"`
//         TeamID int `json:"team_id"`
//     }

//     if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "Invalid JSON input"}
//         c.ServeJSON()
//         return
//     }

//     if req.UserID == 0 || req.TeamID == 0 {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "User ID and Team ID are required"}
//         c.ServeJSON()
//         return
//     }

//     db, err := utils.GetDBConnection("collection")
//     if err != nil {
//         log.Printf("Database connection error: %v", err)
//         c.Ctx.Output.SetStatus(500)
//         c.Data["json"] = map[string]string{"error": "Failed to connect to the database"}
//         c.ServeJSON()
//         return
//     }
//     defer db.Close()

//     var exists bool
//     checkQuery := `
//         SELECT EXISTS (
//             SELECT 1 FROM user_teams WHERE user_id = $1 AND team_id = $2
//         )
//     `
//     err = db.QueryRow(checkQuery, req.UserID, req.TeamID).Scan(&exists)
//     if err != nil {
//         log.Printf("Database check error: %v", err)
//         c.Ctx.Output.SetStatus(500)
//         c.Data["json"] = map[string]string{"error": "Failed to check team assignment"}
//         c.ServeJSON()
//         return
//     }

//     if exists {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "User is already assigned to this team"}
//         c.ServeJSON()
//         return
//     }

//     insertQuery := `
//         INSERT INTO user_teams (user_id, team_id, created_at, updated_at)
//         VALUES ($1, $2, $3, $4)
//     `
//     _, err = db.Exec(insertQuery, req.UserID, req.TeamID, time.Now(), time.Now())
//     if err != nil {
//         log.Printf("Insert user-team error: %v", err)
//         c.Ctx.Output.SetStatus(500)
//         c.Data["json"] = map[string]string{"error": "Failed to assign user to team"}
//         c.ServeJSON()
//         return
//     }

//     c.Ctx.Output.SetStatus(201)
//     c.Data["json"] = map[string]string{
//         "message": "User successfully assigned to team",
//     }
//     c.ServeJSON()
// }
