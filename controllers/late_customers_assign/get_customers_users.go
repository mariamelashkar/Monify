package filters

// import (
// 	"log"
// )
// func (c *CollectionController) FilterAndSelectCustomers() ([]map[string]interface{}, error) {
//     customersData, err := c.FilterCustomersByAnyParam()
//     if err != nil {
//         return nil, err
//     }

//     if len(customersData) == 0 {
//         log.Println("No customers found to assign")
//         return nil, nil
//     }

//     return customersData, nil
// }

// func (c *CollectionController) GetAvailableUsersForAssignment(nextLevel int) ([]map[string]interface{}, error) {
//     usersAtNextLevel, err := c.GetUsersAtNextHierarchyLevel(nextLevel)
//     if err != nil {
//         return nil, err
//     }

//     if len(usersAtNextLevel) == 0 {
//         log.Println("No users found at the next hierarchy level")
//         return nil, nil
//     }

//     return usersAtNextLevel, nil
// }

// func (c *CollectionController) FilterAndAssignCustomers() {
//     customersData, err := c.FilterAndSelectCustomers()
//     if err != nil || len(customersData) == 0 {
//         log.Println("No customers available for assignment")
//         return
//     }

//     usersAtNextLevel, err := c.GetAvailableUsersForAssignment(1)  
//     if err != nil || len(usersAtNextLevel) == 0 {
//         log.Println("No users available for assignment")
//         return
//     }

//     selectedUserID := usersAtNextLevel[0]["id"].(string)

//     err = c.AssignSelectedCustomersToUser(customersData, selectedUserID)
//     if err != nil {
//         log.Println("Failed to assign customers")
//         return
//     }

//     log.Println("Customers successfully assigned to the selected user")
// }
