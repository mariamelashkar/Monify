package filters

import (
	"collection/utils"
	"database/sql"
)

func GetUsersAtNextLevel(db *sql.DB, nextLevel int) ([]map[string]interface{}, error) {
	query := `
SELECT u.id, u.name, h.level
FROM users u
JOIN user_hierarchy uh ON u.id = uh.user_id
JOIN hierarchy h ON uh.hierarchy_id = h.id
WHERE h.level = $1
AND u.deleted = false
AND h.deleted = false;

    `

	params := []interface{}{nextLevel}

	results, err := utils.QueryExecuter(db, "getNextLevelUsers", query, params)
	if err != nil {
		return nil, err
	}

	return results, nil
}
