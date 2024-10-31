package models

import (
	"collection/utils"
	//"encoding/json"
	"errors"
	"time"
	"fmt"
)
type LateLoansSnap struct {
	Id                     string     `orm:"column(id);pk;size(70)" json:"id"`
	CheckDate              time.Time  `orm:"column(check_date);type(date)" json:"check_date"`
	LateAmount             float64    `orm:"column(lateamount);type(float)" json:"late_amount"`
	InsCount               int        `orm:"column(inscount);type(int)" json:"ins_count"`
	Lid                    string     `orm:"column(lid);size(255)" json:"lid"`
	Lsid                   string     `orm:"column(lsid);size(255)" json:"lsid"`
	BranchId               string     `orm:"column(branch_id);size(255)" json:"branch_id"`
	BranchName             string     `orm:"column(branch_name);size(255)" json:"branch_name"`
	CustomerId             string     `orm:"column(customer_id);size(255)" json:"customer_id"`
	CustomerName           string     `orm:"column(customer_name);size(255)" json:"customer_name"`
	Key                    string     `orm:"column(key);size(255)" json:"key"`
	RepresentativeId       string     `orm:"column(representative_id);size(255)" json:"representative_id"`
	LoanOfficer            string     `orm:"column(loan_officer);size(255)" json:"loan_officer"`
	Principal              float64    `orm:"column(principal);type(float)" json:"principal"`
	IssueDate              int64      `orm:"column(issue_date);type(bigint)" json:"issue_date"`
	LoanKey                string     `orm:"column(loan_key);size(255)" json:"loan_key"`
	TotalInstallmentSum    float64    `orm:"column(total_installment_sum);type(float)" json:"total_installment_sum"`
	CustomerHomeAddress    string     `orm:"column(customer_home_address);size(255)" json:"customer_home_address"`
	NationalId             string     `orm:"column(national_id);size(255);null" json:"national_id"`
	CustomerAddressLatLong string     `orm:"column(customer_address_lat_long);size(255)" json:"customer_address_lat_long"`
	HomePhoneNumber        string     `orm:"column(home_phone_number);size(20)" json:"home_phone_number"`
	MobilePhoneNumber      string     `orm:"column(mobile_phone_number);size(20)" json:"mobile_phone_number"`
	BusinessName           string     `orm:"column(business_name);size(255)" json:"business_name"`
	BusinessAddress        string     `orm:"column(business_address);size(255)" json:"business_address"`
	BusinessPhoneNumber    string     `orm:"column(business_phone_number);size(20)" json:"business_phone_number"`
	CurrentRep             string     `orm:"column(current_rep);size(255)" json:"current_rep"`
	CurrentRepName         string     `orm:"column(current_rep_name);size(255)" json:"current_rep_name"`
	CreatedAt              time.Time  `orm:"column(created_at);type(datetime);auto_now_add" json:"created_at"`
	CreatedBy              string     `orm:"column(created_by);size(70)" json:"created_by"`
	UpdatedBy              *string    `orm:"column(updated_by)" json:"updated_by"`
	UpdatedAt              *time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
	Actions                string     `orm:"column(action);size(200);default(created)" json:"action"`
	Deleted                bool       `orm:"column(deleted);default(false)" json:"deleted"`
}


func (L *LateLoansSnap) TableName() string {
	return "collection_late_loans_snap"
}

// // ---------- CRUD Operations for late_loan_snap Model ----------//
func GetLateLoansSnap(CustomerID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
  SELECT 
    id, check_date, lateamount AS late_amount, inscount AS ins_count, lid, lsid,
    branch_id, branch_name, customer_id, customer_name, key, representative_id, 
    loan_officer, principal, issue_date, loan_key, total_installment_sum,
    customer_home_address, national_id, customer_address_lat_long, 
    home_phone_number, mobile_phone_number, business_name, business_address,
    business_phone_number, current_rep, current_rep_name, created_at, created_by,
    updated_at, updated_by, action, deleted
FROM collection_late_loans_snap
WHERE customer_id = $1
`

	params := []interface{}{CustomerID}

	result, err := utils.QueryExecuter(db, "late_loans_result", query, params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Result of the id %s", CustomerID)
	if len(result) == 0 {
		return nil, errors.New("no late loans snap found with the given ID")
	}

	return result[0], nil
}

func GetAllLateLoansSnap() ([]map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
	SELECT 
		id, check_date, lateamount AS late_amount, inscount AS ins_count, lid, lsid,
		branch_id, branch_name, customer_id, customer_name, key, representative_id, 
		loan_officer, principal, issue_date, loan_key, total_installment_sum,
		customer_home_address, national_id, customer_address_lat_long, 
		home_phone_number, mobile_phone_number, business_name, business_address,
		business_phone_number, current_rep, current_rep_name, created_at, created_by,
		updated_at, updated_by, action, deleted
	FROM collection_late_loans_snap
	`

	result, err := utils.QueryExecuter(db, "late_loans_result", query, nil)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no late loans found")
	}

	var mappedSnaps []map[string]interface{}
	for _, row := range result {
		mappedFields := make(map[string]interface{})

		for key, value := range row {
			mappedFields[key] = value
		}

		mappedSnaps = append(mappedSnaps, mappedFields)
	}

	return mappedSnaps, nil
}
