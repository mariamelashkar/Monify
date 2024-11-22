package utils
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)
func QueryExecuter(db *sql.DB, queryName string, query string,params []interface{}) ([]map[string]interface{}, error) {

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	
	scanValues := make([]interface{}, len(columns))
	for i := range scanValues {
		scanValues[i] = new(sql.NullString)
	}

	var results []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(scanValues...)
		if err != nil {
			return nil, err
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			ns := *(scanValues[i].(*sql.NullString))
			if ns.Valid {
				// 	//debug line

				// fmt.Printf("Column: %s, Type: %T, Value: %v\n", col, ns.String, ns.String)

				rowMap[col] = ns.String
			} else {
				rowMap[col] = nil
			}
		}

		results = append(results, rowMap)
	}
	// //debug line
	// fmt.Println("results:",results)

	result := make(map[string][]map[string]interface{})
	result[queryName] = results

	return results, nil
}

