
package controllers

import (
	_"github.com/lib/pq"

)

func getBranchesReallocateQuery() string {
	return `SELECT DISTINCT cs.branch_id, cs.branch_name FROM ticket_table ct LEFT JOIN collection_late_customer_snap cs ON ct.customer_ticket_id = cs.id`
}

func getAllAvailableFollowerQuery() string {
	return `SELECT uf.id, u."name" FROM followers uf LEFT JOIN "User_lts_data" u ON uf.user_id_id = u.id`
}

func getUnderWorkAllBranchesQuery() string {
	return `
	WITH CTE AS (
		SELECT
			ct.follow_state,
			CASE
				WHEN follow_state = 0 THEN 0
				WHEN follow_state = 1 THEN 0
				WHEN follow_state = 2 THEN 1
				WHEN follow_state = 3 THEN 2
			END AS state_agg,
			SUM(cs.total_late) AS lates_amount,
			SUM(cs.late_ins_count) AS late_ins_count,
			SUM(cs.late_loans_count) AS loan_count
		FROM
			ticket_table ct
			LEFT JOIN collection_late_customer_snap cs ON ct.customer_ticket_id = cs.id
		WHERE
			ct.follow_state IN (0, 1)
		GROUP BY
			ct.follow_state
	)
	SELECT
		CASE
			WHEN follow_state = 0 THEN 'تم التوزيع في انتظار المتابعة'
			WHEN follow_state = 1 THEN 'تحت العمل'
			WHEN follow_state = 2 THEN 'تم الاغلاق و عدم وجود نتيجة'
			WHEN follow_state = 3 THEN 'تم الاقفال و اغلاق المديونية'
			ELSE 'unknown'
		END AS case_state,
		lates_amount,
		late_ins_count,
		loan_count,
		CAST(ROUND((lates_amount / SUM(lates_amount) OVER (PARTITION BY state_agg))::numeric, 2) * 100 AS float) AS percent_of_amount
	FROM
		CTE`
}

func getUnderWorkBranchesQuery() string {
	return `
	WITH CTE AS (
		SELECT
			cs.branch_name,
			ct.follow_state,
			SUM(cs.total_late) AS lates_amount,
			SUM(cs.late_ins_count) AS late_ins_count,
			SUM(cs.late_loans_count) AS loan_count
		FROM
			ticket_table ct
			LEFT JOIN collection_late_customer_snap cs ON ct.customer_ticket_id = cs.id
		WHERE
			ct.follow_state IN (0, 1)
		GROUP BY
			cs.branch_name,
			ct.follow_state
	)
	SELECT
		branch_name,
		follow_state,
		CASE
			WHEN follow_state = 0 THEN 'تم التوزيع في انتظار المتابعة'
			WHEN follow_state = 1 THEN 'تحت العمل'
			WHEN follow_state = 2 THEN 'تم الاغلاق و عدم وجود نتيجة'
			WHEN follow_state = 3 THEN 'تم الاقفال و اغلاق المديونية'
			ELSE 'unknown'
		END AS case_state,
		lates_amount,
		late_ins_count,
		loan_count,
		CAST(ROUND((lates_amount / SUM(lates_amount) OVER (PARTITION BY branch_name))::numeric * 100, 2) AS float) AS percent_ava
	FROM
		CTE
	ORDER BY
		branch_name,
		follow_state`
}