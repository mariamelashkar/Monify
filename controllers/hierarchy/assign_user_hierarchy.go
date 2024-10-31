package hierarchy

// import (
//     "collection/utils"
//     "encoding/json"
//     beego "github.com/beego/beego/v2/server/web"
//     "log"
//     "time"
// )

// type UserHierarchyController struct {
//     beego.Controller
// }

// func (c *UserHierarchyController) AssignHierarchyToUser() {
//     var req struct {
//         UserID      int `json:"user_id"`
//         HierarchyID int `json:"hierarchy_id"`
//     }

//     if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "Invalid JSON input"}
//         c.ServeJSON()
//         return
//     }

//     if req.UserID == 0 || req.HierarchyID == 0 {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "User ID and Hierarchy ID are required"}
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
//             SELECT 1 FROM user_hierarchy WHERE user_id = $1 AND hierarchy_id = $2
//         )
//     `
//     err = db.QueryRow(checkQuery, req.UserID, req.HierarchyID).Scan(&exists)
//     if err != nil {
//         log.Printf("Database check error: %v", err)
//         c.Ctx.Output.SetStatus(500)
//         c.Data["json"] = map[string]string{"error": "Failed to check hierarchy assignment"}
//         c.ServeJSON()
//         return
//     }

//     if exists {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "User is already assigned to this hierarchy"}
//         c.ServeJSON()
//         return
//     }

//     insertQuery := `
//         INSERT INTO user_hierarchy (user_id, hierarchy_id, created_at, updated_at)
//         VALUES ($1, $2, $3, $4)
//     `
//     _, err = db.Exec(insertQuery, req.UserID, req.HierarchyID, time.Now(), time.Now())
//     if err != nil {
//         log.Printf("Insert user-hierarchy error: %v", err)
//         c.Ctx.Output.SetStatus(500)
//         c.Data["json"] = map[string]string{"error": "Failed to assign user to hierarchy"}
//         c.ServeJSON()
//         return
//     }

//     c.Ctx.Output.SetStatus(201)
//     c.Data["json"] = map[string]string{
//         "message": "User successfully assigned to hierarchy",
//     }
//     c.ServeJSON()
// }
