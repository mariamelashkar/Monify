package controllers

import (
	"database/sql"
	"sync"
)

func BranchesAva(DB *sql.DB, results chan<- map[string]interface{}, errChan chan<- string, done chan<- string, wg *sync.WaitGroup) {
	query := `SELECT DISTINCT collection_late_loans_snap.branch_id, collection_late_loans_snap.branch_name FROM collection_late_loans_snap`
	ExecuteGoQuery(DB, "select_branch", query, results, errChan, done, nil, wg)
}

func AllBranchesLates(DB *sql.DB, results chan<- map[string]interface{}, errChan chan<- string, done chan<- string, wg *sync.WaitGroup) {
	query := `WITH late_branch AS (
		SELECT * FROM collection_late_loans_snap WHERE check_date = CURRENT_DATE 
	)
	SELECT branch_id, branch_name, SUM(lateamount) AS late_amounts, SUM(inscount) AS inscount, COUNT(DISTINCT lsid) AS loan_count,
		CAST(CAST((SUM(lateamount) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS FLOAT) AS percentage_late_amount 
	FROM late_branch 
	GROUP BY branch_id, branch_name 
	ORDER BY percentage_late_amount DESC`
	ExecuteGoQuery(DB, "all_branches_late", query, results, errChan, done, nil, wg)
}

func FollowersAll(DB *sql.DB, results chan<- map[string]interface{}, errChan chan<- string, done chan<- string, wg *sync.WaitGroup) {
	query := `SELECT cf.id, u."name"
			  FROM followers cf
			  LEFT JOIN "User_lts_data" u ON u.id = cf.user_id_id 
			  WHERE cf.state = 0`
			  ExecuteGoQuery(DB, "followers_all", query, results, errChan, done, nil, wg)
}

func StartCustomerSample(DB *sql.DB, results chan<- map[string]interface{}, errChan chan<- string, done chan<- string, wg *sync.WaitGroup) {
	query := `WITH late_branch AS (
		SELECT * FROM collection_late_loans_snap WHERE check_date = CURRENT_DATE 
		AND customer_id NOT IN (
			SELECT cs.customer_id 
			FROM ticket_table ti 
			LEFT JOIN collection_late_customer_snap cs ON ti.customer_ticket_id = cs.id 
			AND ti.follow_state NOT IN (3, 4)
		)
	)
	SELECT branch_id, branch_name, customer_id, customer_name, "key", representative_id, loan_officer, MAX(customer_home_address) AS address,
		SUM(COUNT(DISTINCT lsid)) OVER (PARTITION BY customer_id) AS loans_count,
		SUM(SUM(lateamount)) OVER (PARTITION BY customer_id) AS late_amounts,
		SUM(SUM(inscount)) OVER (PARTITION BY customer_id) AS inscount,
		CAST(CAST((SUM(SUM(lateamount)) OVER (PARTITION BY customer_id) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS FLOAT) AS percentage_late_amount 
	FROM late_branch 
	GROUP BY branch_name, customer_id, customer_name, "key", representative_id, loan_officer 
	LIMIT 25`
	ExecuteGoQuery(DB, "start_customer_sample", query, results, errChan, done, nil, wg)
}
