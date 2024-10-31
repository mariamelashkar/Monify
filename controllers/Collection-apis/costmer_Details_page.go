package controllers

// import (
// 	"collection/utils"
// 	"log"
// 	"sync"
// )

// func (c *CustomerController) DetailsPage() {
// 	ticketId := c.GetString("ticketid")
// 	if ticketId == "" {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = map[string]string{"error": "ticketid is required"}
// 		c.ServeJSON()
// 		return
// 	}

// 	DB, err := utils.GetDBConnection()
// 	if err != nil {
// 		log.Println("Error connecting to database:", err)
// 		c.Ctx.Output.SetStatus(500)
// 		c.Data["json"] = map[string]string{"error": "Error connecting to database"}
// 		c.ServeJSON()
// 		return
// 	}
// 	defer DB.Close()

// 	results := make(chan map[string]interface{})
// 	errChan := make(chan string)
// 	done := make(chan string)
// 	var wg sync.WaitGroup

// 	wg.Add(5)
// 	go ExecuteGoQuery(DB, "chartlates", getCusDetailsChartQuery(), results, errChan, done, []interface{}{ticketId}, &wg)
// 	go ExecuteGoQuery(DB, "customer_action", getCusDetailsActionsQuery(), results, errChan, done, []interface{}{ticketId}, &wg)
// 	go ExecuteGoQuery(DB, "customer_main", getCusDetailsMainDataQuery(), results, errChan, done, []interface{}{ticketId}, &wg)
// 	go ExecuteGoQuery(DB, "customer_loans", getCusDetailsLoansQuery(), results, errChan, done, []interface{}{ticketId}, &wg)
// 	go ExecuteGoQuery(DB, "follow_result", getCusDetailsFolResQuery(), results, errChan, done, []interface{}{ticketId}, &wg)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 		close(errChan)
// 		close(done)
// 	}()

// 	finalResults := make(map[string]interface{})
// 	for range [count5]struct{}{} {
// 		select {
// 		case res := <-results:
// 			for key, value := range res {
// 				finalResults[key] = value
// 			}
// 		case errMsg := <-errChan:
// 			log.Println(errMsg)
// 			c.Ctx.Output.SetStatus(500)
// 			c.Data["json"] = map[string]string{"error": errMsg}
// 			c.ServeJSON()
// 			return
// 		case <-done:
// 		}
// 	}

// 	c.Data["json"] = finalResults
// 	c.ServeJSON()
// }
