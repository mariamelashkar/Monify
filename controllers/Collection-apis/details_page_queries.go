package controllers

import (
	_"database/sql"
)

func getCusDetailsChartQuery() string {
	return `
		SELECT lcs.snap_date, lcs.total_late
		FROM ticket_table ct
		LEFT JOIN collection_late_customer_snap lc ON lc.id = ct.customer_ticket_id
		LEFT JOIN collection_late_customer_snap lcs ON lcs.customer_id = lc.customer_id
		WHERE ct.id = $1`
}

func getCusDetailsActionsQuery() string {
	return `
		SELECT * 
		FROM follower_action fa 
		LEFT JOIN action_result ar ON ar.action_id = fa.id
		LEFT JOIN actions ON actions.id = fa.action_id
		LEFT JOIN result_choose ON result_choose.id = ar.result_id
		WHERE fa.ticket_id_id = $1`
}

func getCusDetailsMainDataQuery() string {
	return `
		SELECT lcs.*
		FROM ticket_table ct
		LEFT JOIN collection_late_customer_snap lc ON lc.id = ct.customer_ticket_id
		LEFT JOIN collection_late_customer_snap lcs ON lcs.customer_id = lc.customer_id AND lcs.snap_date = current_date
		WHERE ct.id = $1`
}

func getCusDetailsLoansQuery() string {
	return `
		SELECT ll.*
		FROM ticket_table ct
		LEFT JOIN collection_late_customer_snap lc ON lc.id = ct.customer_ticket_id
		LEFT JOIN collection_late_loans_snap ll ON ll.customer_id = lc.customer_id AND ll.check_date = current_date
		WHERE ct.id = $1`
}

func getCusDetailsFolResQuery() string {
	return `
		SELECT CAST(fa.update_date AS date) AS date, u.name, ar.desc, ar.collected_amount, ar.next_follow_date, ca.action_name, ca.action_type, result_choose.result_name  
		FROM public.follower_action fa
		LEFT JOIN action_result ar ON ar.action_id = fa.id
		LEFT JOIN actions ca ON ca.id = fa.action_id
		LEFT JOIN result_choose ON result_choose.id = ar.result_id
		LEFT JOIN followers cf ON cf.id = fa.follower_id
		LEFT JOIN "User_lts_data" u ON u.id = cf.add_user_id
		LEFT JOIN ticket_table ct ON ct.id = fa.ticket_id_id
		WHERE ct.id = $1`
}