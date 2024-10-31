package controllers
import (
    "encoding/json"
    "log"
)

func PrettyPrintData(Data[]map[string]interface{}) (string, error){
    jsonData, err := json.MarshalIndent(Data, "", "  ")
    if err != nil {
        log.Printf("Error converting customersData to JSON: %v", err)
        return "", err
    }
    log.Printf("Customers Data: %s\n", string(jsonData))
    return string(jsonData), nil
}
