package middlewares

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"
// 	"sync"
// 	"collection/controllers"
// 	"collection/internal/redis"
// 	"github.com/beego/beego/v2/server/web/context"
// 	"collection/utils"
// 	"log"
// )

// type contextKey string

// const (
// 	UserKey      contextKey = "user"
// 	AllUserGroupsKey contextKey = "allUserGroups"
// 	AllUserBranchesKey contextKey = "allUserBranches"
// 	count = 3

// )

// func AuthFilter(ctx *context.Context) {
// 	if ctx.Request.URL.Path == "/login" {
// 		return
// 	}

// 	authHeader := ctx.Request.Header.Get("Authorization")
// 	if authHeader == "" {
// 		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
// 		ctx.Output.Body([]byte("Authorization header required"))
// 		return
// 	}

// 	parts := strings.Split(authHeader, "Bearer ")
// 	if len(parts) != 2 {
// 		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
// 		ctx.Output.Body([]byte("Invalid authorization header format"))
// 		return
// 	}

// 	tokenString := parts[1]
// 	claims, err := redis.ParseJWT(tokenString) 
// 	if err != nil {
// 		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
// 		ctx.Output.Body([]byte("Invalid token"))
// 		return
// 	}

// 	userID, ok := claims["user_id"].(string)
// 	if !ok {
// 		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
// 		ctx.Output.Body([]byte("Invalid token claims"))
// 		return
// 	}

// 	queryUsers := "SELECT * FROM user WHERE user_id = $1"
// 	queryUserGroups := "SELECT * FROM group WHERE user_id = $1"
// 	queryUserBranches := "SELECT * FROM branches WHERE user_id = $1"

// 	results := make(chan map[string]interface{})
// 	errChan := make(chan string)
// 	done := make(chan string)
// 	var wg sync.WaitGroup

// 	wg.Add(count)

// 	db, err := utils.GetDBConnection()
// 	if err != nil {
// 		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
// 		ctx.Output.Body([]byte(fmt.Sprintf("Error connecting to database: %v", err)))
// 		return
// 	}
// 	defer db.Close()

// 	go func() {
// 		controllers.ExecuteGoQuery(db, "users", queryUsers, results, errChan, done, []interface{}{userID}, &wg)
// 	}()

// 	go func() {
// 		controllers.ExecuteGoQuery(db, "userGroups", queryUserGroups, results, errChan, done, []interface{}{userID}, &wg)
// 	}()

// 	go func() {
// 	controllers.ExecuteGoQuery(db, "userBranches", queryUserBranches, results, errChan, done, []interface{}{userID}, &wg)
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 		close(errChan)
// 		close(done)
// 	}()

// 	finalResults := make(map[string]interface{})
// 	for i := 0; i < count; i++ {
// 		select {
// 		case res := <-results:
// 			for key, value := range res {
// 				finalResults[key] = value
// 			}
// 		case errMsg := <-errChan:
// 			log.Println(errMsg)
// 			ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
// 			ctx.Output.Body([]byte(errMsg))
// 			return
// 		case <-done:
// 		}
// 	}

// 	ctx.Input.SetData(string(UserKey), finalResults["user"])
// 	ctx.Input.SetData(string(AllUserGroupsKey), finalResults["userGroups"])
// 	ctx.Input.SetData(string(AllUserBranchesKey), finalResults["userBranches"])
// }
