package filters

import (
	"log"
	"collection/utils"
	"fmt"
)

func (c *CollectionController) AssignSelectedCustomersToUser(customersData []map[string]interface{}, selectedUserID string) error {
    db, err := utils.GetDBConnection()
    if err != nil {
        log.Printf("Error connecting to the database: %v", err)
        return err
    }
    defer db.Close()

    var userCount int
    userCheckQuery := "SELECT COUNT(*) FROM users WHERE id = $1"
    err = db.QueryRow(userCheckQuery, selectedUserID).Scan(&userCount)
    if err != nil || userCount == 0 {
        log.Printf("User not found: %v", err)
        return fmt.Errorf("user not found")
    }

    for _, customer := range customersData {
        customerID := customer["customer_id"].(string)

        assignQuery := `
            INSERT INTO user_customer_assignment (user_id, customer_id, assigned_at)
            VALUES ($1, $2, NOW())
            ON CONFLICT (customer_id) DO UPDATE 
            SET user_id = EXCLUDED.user_id, assigned_at = NOW()
        `
        _, err := db.Exec(assignQuery, selectedUserID, customerID)
        if err != nil {
            log.Printf("Error assigning customer %s to user %s: %v", customerID, selectedUserID, err)
            return err
        }

        log.Printf("Customer %s successfully assigned to user %s", customerID, selectedUserID)
    }

    return nil
}
