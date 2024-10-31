package middlewares

// import (
// 	"fmt"
// 	userQuery "legal_back/query/user"
// 	"legal_back/utils"
// 	"net/http"
// 	"strings"

// 	beego "github.com/beego/beego/v2/server/web/context"
// )

// // JWTAuth is a middleware function that validates JWT tokens
// func JWTAuth(c *beego.Context) {
// 	var authHeader string
// 	if strings.HasPrefix(c.Input.URL(), "/v1/login") {
// 		return
// 	}

// 	if strings.HasPrefix(c.Input.URL(), "/swagger") {
// 		return
// 	}

// 	if strings.HasPrefix(c.Input.URL(), "/v1/notification/ws") {
// 		authHeader = c.Input.Query("token")
// 	} else {
// 		// Handle standard HTTP requests
// 		authHeader = c.Input.Header("Authorization")
// 	}

// 	if authHeader == "" {
// 		c.Output.SetStatus(http.StatusUnauthorized)
// 		c.Output.JSON(map[string]string{"error": "Missing authorization header"}, true, true)
// 		return
// 	}

// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 	claims, err := utils.ValidateJWT(tokenString)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.Output.SetStatus(http.StatusUnauthorized)
// 		c.Output.JSON(map[string]string{"error": "Invalid token"}, true, true)
// 		return
// 	}

// 	// Fetch user data from database
// 	user, err := userQuery.GetUserById(claims.Id)
// 	if err != nil {
// 		c.Output.SetStatus(http.StatusInternalServerError)
// 		c.Output.JSON(map[string]string{"error": "Failed to fetch user data"}, true, true)
// 		return
// 	}

// 	// Fetch user's groups from database
// 	groups, err := userQuery.GetUserGroups(user.Id)
// 	if err != nil {
// 		c.Output.SetStatus(http.StatusInternalServerError)
// 		c.Output.JSON(map[string]string{"error": "Failed to fetch user groups"}, true, true)
// 		return
// 	}

// 	// Fetch user's branches from database
// 	branches, err := userQuery.GetBranchesByUserId(user.Id)
// 	if err != nil {
// 		c.Output.SetStatus(http.StatusInternalServerError)
// 		c.Output.JSON(map[string]string{"error": "Failed to fetch user branches"}, true, true)
// 		return
// 	}

// 	// Attach user, groups, and branches to context
// 	c.Input.SetData("user", user)
// 	c.Input.SetData("groups", groups)
// 	c.Input.SetData("branches", branches)
// }
