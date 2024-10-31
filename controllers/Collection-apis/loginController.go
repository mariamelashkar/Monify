package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"collection/models"
// 	"collection/utils"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/beego/beego/v2/client/orm"
// 	beego "github.com/beego/beego/v2/server/web"
// ) 

// type LoginController struct {
// 	beego.Controller
// }

// // Login endpoint
// func (c *LoginController) Login() {
// 	var reqBody struct {
// 		Username string `json:"username"`
// 		Password string `json:"password"`
// 	}

// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqBody); err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusBadRequest)
// 		c.Data["json"] = map[string]string{"error": "Invalid request body"}
// 		c.ServeJSON()
// 		return
// 	}

// 	username := reqBody.Username
// 	password := reqBody.Password

// 	fmt.Println(username, password)

// 	if username == "" || password == "" {
// 		c.Ctx.Output.SetStatus(http.StatusBadRequest)
// 		c.Data["json"] = map[string]string{"error": "Username and password are required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	var user models.AuthUser
// 	o := orm.NewOrm()
// 	err := o.QueryTable("auth_user").Filter("username", username).One(&user)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
// 		c.Data["json"] = map[string]string{"error": "Invalid username"}
// 		c.ServeJSON()
// 		return
// 	}

// 	if !utils.CheckPasswordHash(password, user.Password) {
// 		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
// 		c.Data["json"] = map[string]string{"error": "Invalid credentials"}
// 		c.ServeJSON()
// 		return
// 	}

// 	token, err := utils.GenerateJWT(user.Username, user.Id)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
// 		c.Data["json"] = map[string]string{"error": "Could not generate token"}
// 		c.ServeJSON()
// 		return
// 	}

// 	c.Data["json"] = map[string]string{"token": token}
// 	c.ServeJSON()
// }

// func (c *LoginController) CreateUser() {
// 	var user models.AuthUser
// 	username := c.GetString("username")
// 	password := c.GetString("password")
// 	first_name := c.GetString("first_name")
// 	last_name := c.GetString("last_name")
// 	email := c.GetString("email")
// 	is_superuser_str := c.GetString("is_superuser")
// 	is_staff_str := c.GetString("is_staff")

// 	is_superuser, err := strconv.ParseBool(is_superuser_str)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusBadRequest)
// 		c.Data["json"] = map[string]string{"error": "Invalid value for is_superuser"}
// 		c.ServeJSON()
// 		return
// 	}

// 	is_staff, err := strconv.ParseBool(is_staff_str)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusBadRequest)
// 		c.Data["json"] = map[string]string{"error": "Invalid value for is_staff"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Hash the password
// 	hashedPassword, err := utils.HashPassword(password)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
// 		c.Data["json"] = map[string]string{"error": "Failed to hash password"}
// 		c.ServeJSON()
// 		return
// 	}
// 	user.Password = string(hashedPassword)

// 	// Set default values
// 	user.Username = username
// 	user.FirstName = first_name
// 	user.LastName = last_name
// 	user.Email = email
// 	user.IsStaff = is_staff
// 	user.IsSuperuser = is_superuser
// 	now := time.Now()
// 	user.DateJoined = &now
// 	user.IsActive = true
// 	user.IsStaff = false

// 	// Insert user into the database
// 	o := orm.NewOrm()
// 	_, err = o.Insert(&user)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
// 		c.Data["json"] = map[string]string{"error": "Failed to create user"}
// 		c.ServeJSON()
// 		return
// 	}

// 	// Return the created user
// 	c.Data["json"] = user
// 	c.ServeJSON()
// }

// // Protected endpoint
// func (c *LoginController) UserInfo() {
// 	user := c.Ctx.Input.GetData("user").(models.AuthUser)
// 	groups := c.Ctx.Input.GetData("groups").([]models.AuthGroup)
// 	branches := c.Ctx.Input.GetData("branches").([]models.Branches)

// 	// Construct JSON response
// 	responseData := map[string]interface{}{
// 		"user":     user,
// 		"groups":   groups,
// 		"branches": branches,
// 	}

// 	c.Data["json"] = responseData
// 	c.ServeJSON()
// }

// // Logout endpoint
// func (c *LoginController) Logout() {
// 	authHeader := c.Ctx.Input.Header("Authorization")
// 	if authHeader == "" {
// 		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
// 		c.Data["json"] = map[string]string{"error": "Missing authorization header"}
// 		c.ServeJSON()
// 		return
// 	}

// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 	utils.AddToBlacklist(tokenString)

// 	c.Data["json"] = map[string]string{"message": "Successfully logged out"}
// 	c.ServeJSON()
// }

// // RefreshToken endpoint
// func (c *LoginController) RefreshToken() {
// 	authHeader := c.Ctx.Input.Header("Authorization")
// 	if authHeader == "" {
// 		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
// 		c.Data["json"] = map[string]string{"error": "Missing authorization header"}
// 		c.ServeJSON()
// 		return
// 	}

// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 	newToken, err := utils.RefreshJWT(tokenString)
// 	if err != nil {
// 		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
// 		c.Data["json"] = map[string]string{"error": "Invalid or expired token"}
// 		c.ServeJSON()
// 		return
// 	}

// 	c.Data["json"] = map[string]string{"token": newToken}
// 	c.ServeJSON()
// }
