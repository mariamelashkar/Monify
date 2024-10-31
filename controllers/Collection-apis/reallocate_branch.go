package controllers

// import (
// 	"collection/utils"
// 	"log"
// 	"sync"

// 	_ "github.com/lib/pq"
// )

// func (c *AllocationController) ReallocationBranch() {
// 	branchId := c.GetString("branchId")
// 	if branchId == "" {
// 		c.Ctx.Output.SetStatus(400)
// 		c.Data["json"] = map[string]string{"error": "branchId is required"}
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

// 	wg.Add(2)
// 	go ExecuteGoQuery(DB, "followersHavePortfolioAva", getFollowerHavePortfolioAvaQuery(), results, errChan, done, []interface{}{branchId}, &wg)
// 	go ExecuteGoQuery(DB, "allBranchCustomers", getAllBranchCustomersQuery(), results, errChan, done, []interface{}{branchId}, &wg)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 		close(errChan)
// 		close(done)
// 	}()

// 	finalResults := make(map[string]interface{})
// 	for range [count2]struct{}{} {
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

// 	c.Data["json"] = map[string]interface{}{
// 		"state":            true,
// 		"followers_ava":    finalResults["followersHavePortfolioAva"],
// 		"branch_customers": finalResults["allBranchCustomers"],
// 	}
// 	c.ServeJSON()
// }
