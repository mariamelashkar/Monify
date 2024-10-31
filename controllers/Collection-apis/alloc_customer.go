package controllers

// import (
// 	utils "collection/utils"
// 	"log"
// 	"time"
// )

// type JsonResponse struct {
// 	State bool   `json:"state"`
// 	Error string `json:"error,omitempty"`
// }

// func (c *AllocationController) AllocateCustomers() {
// 	userId := c.GetString("userId")
// 	if userId == "" {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = JsonResponse{State: false, Error: "userId is required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	collectionEmp := c.GetString("employee")
// 	if collectionEmp == "" {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = JsonResponse{State: false, Error: "employee is required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	customerArray := c.GetStrings("customers[]")
// 	if len(customerArray) == 0 {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = JsonResponse{State: false, Error: "customers[] is required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	DB, err := utils.GetDBConnection()
// 	if err != nil {
// 		log.Println("Error connecting to database:", err)
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = JsonResponse{State: false, Error: "Error connecting to database"}
// 		c.ServeJSON()
// 		return
// 	}
// 	defer DB.Close()

// 	tx, err := DB.Begin()
// 	if err != nil {
// 		log.Println("Error starting transaction:", err)
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = JsonResponse{State: false, Error: "Error starting transaction"}
// 		c.ServeJSON()
// 		return
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		} else {
// 			tx.Commit()
// 		}
// 	}()

// 	for _, customerId := range customerArray {
// 		var customerID int
// 		err := tx.QueryRow("SELECT id FROM collection_late_loans_snap WHERE snap_date = $1 AND customer_id = $2", time.Now().Format("2006-01-02"), customerId).Scan(&customerID)
// 		if err != nil {
// 			log.Println("Error fetching customer:", err)
// 			c.Ctx.Output.SetStatus(500)
// 			c.Data["json"] = JsonResponse{State: false, Error: "Error fetching customer"}
// 			c.ServeJSON()
// 			return
// 		}

// 		_, err = tx.Exec("INSERT INTO ticketTable (customer_ticket, follow_user, add_user) VALUES ($1, $2, $3)", customerID, collectionEmp, userId)
// 		if err != nil {
// 			log.Println("Error saving ticketTable:", err)
// 			c.Ctx.Output.SetStatus(500)
// 			c.Data["json"] = JsonResponse{State: false, Error: "Error saving ticketTable"}
// 			c.ServeJSON()
// 			return
// 		}

// 		var ticketId int
// 		err = tx.QueryRow("SELECT id FROM ticketTable WHERE customer_ticket = $1 AND follow_user = $2 AND add_user = $3 ORDER BY id DESC LIMIT 1", customerID, collectionEmp, userId).Scan(&ticketId)
// 		if err != nil {
// 			log.Println("Error fetching ticket ID:", err)
// 			c.Ctx.Output.SetStatus(500)
// 			c.Data["json"] = JsonResponse{State: false, Error: "Error fetching ticket ID"}
// 			c.ServeJSON()
// 			return
// 		}

// 		_, err = tx.Exec("INSERT INTO currentFollower (ticket_id, follow_user, add_user) VALUES ($1, $2, $3)", ticketId, collectionEmp, userId)
// 		if err != nil {
// 			log.Println("Error saving currentFollower:", err)
// 			c.Ctx.Output.SetStatus(500)
// 			c.Data["json"] = JsonResponse{State: false, Error: "Error saving currentFollower"}
// 			c.ServeJSON()
// 			return
// 		}
// 	}

// 	c.Data["json"] = JsonResponse{State: true}
// 	c.ServeJSON()
// }
