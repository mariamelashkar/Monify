
package controllers

import (
	_"github.com/lib/pq"

)
func getFollowerHavePortfolioAvaQuery() string {
	return `
		WITH ava_user_id AS (
			SELECT DISTINCT ct.follow_user_id AS follower_id
			FROM collectionm_tickettable ct 
			LEFT JOIN collection_late_customer_snap cs ON ct.customer_ticket_id = cs.id 
			WHERE ct.follow_state IN (0,1) AND cs.branch_id = $1
		)
		SELECT DISTINCT cf.id, "User_lts_data"."name" 
		FROM ava_user_id ai 
		LEFT JOIN collectionm_followers cf ON ai.follower_id = cf.id 
		LEFT JOIN "User_lts_data" ON "User_lts_data".id = cf.user_id_id`
}

func getAllBranchCustomersQuery() string {
	return `
		WITH ava_customer AS( 
			SELECT ct.id, ct.follow_user_id, 
			(SELECT u."name" 
			 FROM collectionm_followers uf 
			 LEFT JOIN "User_lts_data" u ON uf.user_id_id = u.id 
			 WHERE uf.id = ct.follow_user_id) AS follow_user_name, 
			ct.add_date, cs.id AS cs_id, cs.snap_date, cs.branch_name, cs.customer_name, cs.principal, cs.installments, cs.total_late, cs.late_ins_count, cs.late_loans_count,
			cs1.snap_date AS date_of_now, cs1.principal AS principal_now, cs1.installments AS ins_now, cs1.total_late AS late_now, cs1.late_ins_count AS ins_late_count_now, cs1.late_loans_count AS late_loans_now
			FROM collectionm_tickettable ct 
			LEFT JOIN collection_late_customer_snap cs ON ct.customer_ticket_id = cs.id 
			LEFT JOIN collection_late_customer_snap cs1 ON cs.customer_id = cs1.customer_id AND cs1.snap_date = current_date
			WHERE ct.follow_state IN (0,1) AND cs.branch_id = $1
		)
		SELECT * FROM ava_customer`
}