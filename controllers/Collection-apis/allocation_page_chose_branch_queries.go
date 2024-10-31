package controllers

import (
)


func getBranchAggQuery() string {
	return `
		WITH late_branch AS (
			SELECT * FROM collection_late_loans_snap WHERE check_date = current_date
		), all_branches AS (
			SELECT branch_id, branch_name, SUM(lateamount) AS late_amounts, SUM(inscount) AS inscount,
			CAST(CAST((SUM(lateamount) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS float) AS percentage_late_amount
			FROM late_branch GROUP BY branch_id, branch_name
		)
		SELECT * FROM all_branches WHERE branch_id = $1`
}

func getEmpBranchAggQuery() string {
	return `
		WITH late_branch AS (
			SELECT * FROM collection_late_loans_snap WHERE check_date = current_date AND branch_id = $1
		)
		SELECT representative_id, loan_officer, SUM(lateamount) AS late_amounts, SUM(inscount) AS inscount, COUNT(DISTINCT lsid) AS loans,
		CAST(CAST((SUM(lateamount) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS float) AS percentage_late_amount
		FROM late_branch GROUP BY representative_id, loan_officer ORDER BY percentage_late_amount DESC`
}

func getEmpBranchAggAfterQuery() string {
	return `
		WITH late_branch AS (
			SELECT * FROM collection_late_loans_snap WHERE check_date = current_date AND branch_id = $1
			AND customer_id NOT IN (
				SELECT cs.customer_id FROM ticket_table ti 
				LEFT JOIN collection_late_customer_snap cs ON ti.customer_ticket_id = cs.id 
				WHERE ti.follow_state NOT IN (3, 4)
			)
		)
		SELECT representative_id, loan_officer, SUM(lateamount) AS late_amounts, SUM(inscount) AS inscount, COUNT(DISTINCT lsid) AS loans,
		CAST(CAST((SUM(lateamount) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS float) AS percentage_late_amount
		FROM late_branch GROUP BY representative_id, loan_officer ORDER BY percentage_late_amount DESC`
}

func getEmpBranchQuery() string {
	return `
		WITH late_branch AS (
			SELECT * FROM collection_late_loans_snap WHERE check_date = current_date AND branch_id = $1
			AND customer_id NOT IN (
				SELECT cs.customer_id FROM ticket_table ti 
				LEFT JOIN collection_late_customer_snap cs ON ti.customer_ticket_id = cs.id 
				WHERE ti.follow_state NOT IN (3, 4)
			)
		)
		SELECT representative_id, loan_officer, SUM(lateamount) AS late_amounts,
		CAST(CAST((SUM(lateamount) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS float) AS percentage_late_amount
		FROM late_branch GROUP BY representative_id, loan_officer`
}

func getBranchCustomerQuery() string {
	return `
		WITH late_branch AS (
			SELECT * FROM collection_late_loans_snap WHERE check_date = current_date AND branch_id = $1
			AND customer_id NOT IN (
				SELECT cs.customer_id FROM ticket_table ti 
				LEFT JOIN collection_late_customer_snap cs ON ti.customer_ticket_id = cs.id 
				WHERE ti.follow_state NOT IN (3, 4)
			)
		)
		SELECT branch_id, branch_name, customer_id, customer_name, "key", representative_id, loan_officer, MAX(customer_home_address) AS address,
		SUM(COUNT(DISTINCT lsid)) OVER (PARTITION BY customer_id) AS loans_count,
		SUM(SUM(lateamount)) OVER (PARTITION BY customer_id) AS late_amounts,
		SUM(SUM(inscount)) OVER (PARTITION BY customer_id) AS inscount,
		CAST(CAST((SUM(SUM(lateamount)) OVER (PARTITION BY customer_id) / SUM(SUM(lateamount)) OVER ()) * 100 AS DECIMAL(10, 2)) AS float) AS percentage_late_amount
		FROM late_branch GROUP BY branch_id, branch_name, customer_id, customer_name, "key", representative_id, loan_officer`
}
