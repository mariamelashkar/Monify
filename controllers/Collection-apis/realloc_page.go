package controllers

// import (
// 	"collection/utils"
// 	"log"
// 	"sync"

// 	_ "github.com/lib/pq"
// )

// func (c *ReallocationController) ReallocationPage() {
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

// 	wg.Add(4)
// 	go ExecuteGoQuery(DB, "branchesReallocate", getBranchesReallocateQuery(), results, errChan, done, nil, &wg)
// 	go ExecuteGoQuery(DB, "allAvailableFollower", getAllAvailableFollowerQuery(), results, errChan, done, nil, &wg)
// 	go ExecuteGoQuery(DB, "underWorkAllBranches", getUnderWorkAllBranchesQuery(), results, errChan, done, nil, &wg)
// 	go ExecuteGoQuery(DB, "underWorkBranches", getUnderWorkBranchesQuery(), results, errChan, done, nil, &wg)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 		close(errChan)
// 		close(done)
// 	}()

// 	finalResults := make(map[string]interface{})
// 	for range [count4]struct{}{} {
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
