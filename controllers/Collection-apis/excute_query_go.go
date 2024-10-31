package controllers

import (
	"database/sql"
	"fmt"
	"sync"
	_ "github.com/go-sql-driver/mysql"
)

func ExecuteGoQuery(db *sql.DB, tableName, query string, results chan<- map[string]interface{}, errChan chan<- string, done chan<- string, params []interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	stmt, err := db.Prepare(query)
	if err != nil {
		errChan <- fmt.Sprintf("Error preparing query for table %s: %v", tableName, err)
		done <- tableName
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(params...)
	if err != nil {
		errChan <- fmt.Sprintf("Error executing query for table %s: %v", tableName, err)
		done <- tableName
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		errChan <- fmt.Sprintf("Error getting columns for table %s: %v", tableName, err)
		done <- tableName
		return
	}

	scanValues := make([]interface{}, len(columns))
	for i := range scanValues {
		scanValues[i] = new(string)
	}

	var resultsData []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(scanValues...)
		if err != nil {
			errChan <- fmt.Sprintf("Error scanning rows for table %s: %v", tableName, err)
			done <- tableName
			return
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = *(scanValues[i]).(*string)
		}
		resultsData = append(resultsData, rowMap)
	}
	results <- map[string]interface{}{
		tableName: resultsData,
	}

	done <- tableName
}
