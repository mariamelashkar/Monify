package models

import (
	"collection/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type LateCustomersSnap struct {
	Id                     string     `orm:"column(id);pk;size(70)" json:"id"`
	SnapDate               time.Time  `orm:"column(snap_date);type(date)" json:"snap_date"`
	BranchId               string     `orm:"column(branch_id);size(24)" json:"branch_id"`
	BranchName             string     `orm:"column(branch_name);size(255)" json:"branch_name"`
	CustomerId             string     `orm:"column(customer_id);size(24);unique" json:"customer_id"`
	CustomerName           string     `orm:"column(customer_name);size(255)" json:"customer_name"`
	CustomerHomeAddress    string     `orm:"column(customer_home_address);size(255)" json:"customer_home_address"`
	NationalId             string     `orm:"column(national_id);size(255)" json:"national_id"`
	CustomerAddressLatLong string     `orm:"column(customer_address_lat_long);size(255)" json:"customer_address_lat_long"`
	HomePhoneNumber        string     `orm:"column(home_phone_number);size(20)" json:"home_phone_number"`
	MobilePhoneNumber      string     `orm:"column(mobile_phone_number);size(20)" json:"mobile_phone_number"`
	BusinessName           string     `orm:"column(business_name);size(255)" json:"business_name"`
	BusinessAddress        string     `orm:"column(business_address);size(255)" json:"business_address"`
	BusinessPhoneNumber    string     `orm:"column(business_phone_number);size(20)" json:"business_phone_number"`
	Representative         string     `orm:"column(representative);size(24)" json:"representative"`
	RepresentativeName     string     `orm:"column(representative_name);size(255)" json:"representative_name"`
	Principal              float64    `orm:"column(principal);type(float)" json:"principal"`
	Installments           float64    `orm:"column(installments);type(float)" json:"installments"`
	TotalLate              float64    `orm:"column(total_late);type(float)" json:"total_late"`
	LateInsCount           int        `orm:"column(late_ins_count);type(int)" json:"late_ins_count"`
	LateLoansCount         int        `orm:"column(late_loans_count);type(int)" json:"late_loans_count"`
	Company                string     `orm:"column(company);size(150);null" json:"company"`
	CreatedAt              time.Time  `orm:"column(created_at);type(datetime);auto_now_add" json:"created_at"`
	CreatedBy              string     `orm:"column(created_by);size(70)" json:"created_by"`
	UpdatedBy              *string    `orm:"column(updated_by)" json:"updated_by"`
	UpdatedAt              *time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
	Actions                string     `orm:"column(action);size(200);default(created)" json:"action"`
	Deleted                bool       `orm:"column(deleted);default(false)" json:"deleted"`
}

func (L *LateCustomersSnap) TableName() string {
	return "collection_late_customer_snap"
}

// ---------- CRUD Operations for late_customer_snap Model ----------//

type LateCustomersSnapFilter struct {
	BranchId           *string    `json:"branch_id,omitempty"`
	CustomerId         *string    `json:"customer_id,omitempty"`
	CustomerName       *string    `json:"customer_name,omitempty"`
	NationalId         *string    `json:"national_id,omitempty"`
	MobilePhoneNumber  *string    `json:"mobile_phone_number"`
	RepresentativeName *string    `json:"representative_name"`
	TotalLate          *float64   `json:"total_late"`
	LateInsCount       *int       `json:"late_ins_count"`
	LateLoansCount     *int       `json:"late_loans_count"`
}

func GetLateCustomersSnap(CustomerID string) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `
  SELECT 
    id, snap_date, branch_id, branch_name, customer_id, customer_name, customer_home_address,
    national_id, customer_address_lat_long, home_phone_number, mobile_phone_number, business_name,
    business_address, business_phone_number, representative, representative_name, principal,
    installments, total_late, late_ins_count, late_loans_count, created_at, created_by, updated_at,
    updated_by, action, deleted, company
FROM collection_late_customer_snap
WHERE customer_id = $1 AND deleted = false
    `

	params := []interface{}{CustomerID}

	Result, err := utils.QueryExecuter(db, "late_customer_result", query, params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Result of the id %s", CustomerID)
	if len(Result) == 0 {
		return nil, errors.New("no late customer snap found with the given ID")
	}

	return Result[0], nil
}

func GetAllLateCustomersSnap(filters *LateCustomersSnapFilter) ([]map[string]interface{}, error) {
    db, err := utils.GetDBConnection()
    if err != nil {
        return nil, err
    }
    defer db.Close()

    query := `
    SELECT 
        id, 
        snap_date, 
        branch_id, 
        branch_name, 
        customer_id, 
        customer_name, 
        customer_home_address, 
        national_id, 
        customer_address_lat_long, 
        home_phone_number, 
        mobile_phone_number, 
        business_name, 
        business_address, 
        business_phone_number, 
        representative, 
        representative_name, 
        principal, 
        installments, 
        total_late, 
        late_ins_count, 
        late_loans_count, 
        company, 
        created_at, 
        created_by, 
        updated_by, 
        updated_at, 
        action, 
        deleted 
    FROM 
        collection_late_customer_snap 
    WHERE 
        1=1
    `

    var params []interface{}
    parameterIndex := 1

    if filters.BranchId != nil {
        query += fmt.Sprintf(" AND branch_id = $%d", parameterIndex)
        params = append(params, *filters.BranchId)
        parameterIndex++
    }
    if filters.CustomerId != nil {
        query += fmt.Sprintf(" AND customer_id = $%d", parameterIndex)
        params = append(params, *filters.CustomerId)
        parameterIndex++
    }
    if filters.CustomerName != nil {
        query += fmt.Sprintf(" AND customer_name ILIKE $%d", parameterIndex)
        params = append(params, "%"+*filters.CustomerName+"%")
        parameterIndex++
    }
    if filters.NationalId != nil {
        query += fmt.Sprintf(" AND national_id = $%d", parameterIndex)
        params = append(params, *filters.NationalId)
        parameterIndex++
    }
    if filters.MobilePhoneNumber != nil {
        query += fmt.Sprintf(" AND mobile_phone_number = $%d", parameterIndex)
        params = append(params, *filters.MobilePhoneNumber)
        parameterIndex++
    }
    if filters.RepresentativeName != nil {
        query += fmt.Sprintf(" AND representative_name ILIKE $%d", parameterIndex)
        params = append(params, "%"+*filters.RepresentativeName+"%")
        parameterIndex++
    }
    if filters.TotalLate != nil {
        query += fmt.Sprintf(" AND total_late >= $%d", parameterIndex)
        params = append(params, *filters.TotalLate)
        parameterIndex++
    }
    if filters.LateInsCount != nil {
        query += fmt.Sprintf(" AND late_ins_count = $%d", parameterIndex)
        params = append(params, *filters.LateInsCount)
        parameterIndex++
    }
    if filters.LateLoansCount != nil {
        query += fmt.Sprintf(" AND late_loans_count = $%d", parameterIndex)
        params = append(params, *filters.LateLoansCount)
        parameterIndex++
    }

    Result, err := utils.QueryExecuter(db, "late_customer_result", query, params)
    if err != nil {
        return nil, err
    }
    if len(Result) == 0 {
        return nil, errors.New("no late customers found")
    }

    var mappedSnaps []map[string]interface{}
    for _, row := range Result {
        mappedFields := make(map[string]interface{})
        for key, value := range row {
            mappedFields[key] = value
        }
        mappedSnaps = append(mappedSnaps, mappedFields)
    }

    return mappedSnaps, nil
}



func CreateLateCustomersSnap(requestBody []byte) (string, error) {
	var lateCustomerSnap map[string]interface{}
	if err := json.Unmarshal(requestBody, &lateCustomerSnap); err != nil {
		return "", errors.New("invalid JSON input")
	}

	if lateCustomerSnap["customer_id"] == nil || lateCustomerSnap["branch_id"] == nil {
		return "", errors.New("customer ID and branch ID are required")
	}

	lateCustomerSnap["created_at"] = utils.CurrentTime
	lateCustomerSnap["updated_at"] = utils.CurrentTime
	lateCustomerSnap["actions"] = "created"
	lateCustomerSnap["deleted"] = false

	query := `
        INSERT INTO collection_late_customer_snap (
            id,snap_date, branch_id, branch_name, customer_id, customer_name, customer_home_address, national_id,
            customer_address_lat_long, home_phone_number, mobile_phone_number, business_name, business_address, 
            business_phone_number, representative, representative_name, principal, installments, total_late,
            late_ins_count, late_loans_count, created_at, created_by, updated_at, updated_by, action, deleted, company
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, 
                $21, $22, $23, $24, $25, $26, $27,$28)
        RETURNING id
    `

	params := []interface{}{
		lateCustomerSnap["id"],
		lateCustomerSnap["snap_date"],
		lateCustomerSnap["branch_id"],
		lateCustomerSnap["branch_name"],
		lateCustomerSnap["customer_id"],
		lateCustomerSnap["customer_name"],
		lateCustomerSnap["customer_home_address"],
		lateCustomerSnap["national_id"],
		lateCustomerSnap["customer_address_lat_long"],
		lateCustomerSnap["home_phone_number"],
		lateCustomerSnap["mobile_phone_number"],
		lateCustomerSnap["business_name"],
		lateCustomerSnap["business_address"],
		lateCustomerSnap["business_phone_number"],
		lateCustomerSnap["representative"],
		lateCustomerSnap["representative_name"],
		lateCustomerSnap["principal"],
		lateCustomerSnap["installments"],
		lateCustomerSnap["total_late"],
		lateCustomerSnap["late_ins_count"],
		lateCustomerSnap["late_loans_count"],
		lateCustomerSnap["created_at"],
		lateCustomerSnap["created_by"],
		lateCustomerSnap["updated_at"],
		lateCustomerSnap["updated_by"],
		lateCustomerSnap["actions"],
		lateCustomerSnap["deleted"],
		lateCustomerSnap["company"],
	}

	db, err := utils.GetDBConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	Result, err := utils.QueryExecuter(db, "late_customer_result", query, params)
	if err != nil {
		return "", err
	}

	if len(Result) > 0 && len(Result[0]) > 0 {
		if id, ok := Result[0]["id"].(string); ok {
			return id, nil
		} else {
			return "", errors.New("failed to retrieve the late customer snap ID")
		}
	}

	return "", errors.New("no result returned from the query")
}

func UpdateLateCustomersSnap(lateCustomersSnapID string, requestBody []byte) (map[string]interface{}, error) {
	fmt.Println(lateCustomersSnapID)
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	existingLateCustomersSnap, err := GetLateCustomersSnap(lateCustomersSnapID)
	if err != nil {
		return nil, errors.New("no late customers snap found with the given ID")
	}
	fmt.Println(existingLateCustomersSnap)

	var updatedLateCustomersSnap map[string]interface{}
	if err := json.Unmarshal(requestBody, &updatedLateCustomersSnap); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	updatedLateCustomersSnap["id"] = existingLateCustomersSnap["id"]
	updatedLateCustomersSnap["created_at"] = existingLateCustomersSnap["created_at"]
	updatedLateCustomersSnap["created_by"] = existingLateCustomersSnap["created_by"]
	updatedLateCustomersSnap["updated_at"] = utils.CurrentTime

	if updatedLateCustomersSnap["updated_by"] == nil {
		return nil, errors.New("the 'UpdatedBy' field is required for updating")
	}

	if updatedLateCustomersSnap["action"] == "" {
		return nil, errors.New("the 'Actions' field is required for updating")
	}

	query := `
        UPDATE collection_late_customer_snap
        SET 
            snap_date = $1, 
            branch_id = $2, 
            branch_name = $3, 
            customer_id = $4, 
            customer_name = $5, 
            customer_home_address = $6, 
            national_id = $7, 
            customer_address_lat_long = $8, 
            home_phone_number = $9, 
            mobile_phone_number = $10, 
            business_name = $11, 
            business_address = $12, 
            business_phone_number = $13, 
            representative = $14, 
            representative_name = $15, 
            principal = $16, 
            installments = $17, 
            total_late = $18, 
            late_ins_count = $19, 
            late_loans_count = $20, 
            updated_at = $21, 
            updated_by = $22, 
            action = $23, 
            deleted = $24,
			company = $25
        WHERE id = $26
        RETURNING *
    `

	params := []interface{}{
		updatedLateCustomersSnap["snap_date"],
		updatedLateCustomersSnap["branch_id"],
		updatedLateCustomersSnap["branch_name"],
		updatedLateCustomersSnap["customer_id"],
		updatedLateCustomersSnap["customer_name"],
		updatedLateCustomersSnap["customer_home_address"],
		updatedLateCustomersSnap["national_id"],
		updatedLateCustomersSnap["customer_address_lat_long"],
		updatedLateCustomersSnap["home_phone_number"],
		updatedLateCustomersSnap["mobile_phone_number"],
		updatedLateCustomersSnap["business_name"],
		updatedLateCustomersSnap["business_address"],
		updatedLateCustomersSnap["business_phone_number"],
		updatedLateCustomersSnap["representative"],
		updatedLateCustomersSnap["representative_name"],
		updatedLateCustomersSnap["principal"],
		updatedLateCustomersSnap["installments"],
		updatedLateCustomersSnap["total_late"],
		updatedLateCustomersSnap["late_ins_count"],
		updatedLateCustomersSnap["late_loans_count"],
		updatedLateCustomersSnap["updated_at"],
		updatedLateCustomersSnap["updated_by"],
		updatedLateCustomersSnap["action"],
		updatedLateCustomersSnap["deleted"],
		updatedLateCustomersSnap["company"],
		lateCustomersSnapID,
	}

	Result, err := utils.QueryExecuter(db, "late_customer_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to update the late customers snap")
	}

	return Result[0], nil
}
func SoftDeleteLateCustomer(lateCustomerID string, requestBody []byte) (map[string]interface{}, error) {
	db, err := utils.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var softDeleteRequest struct {
		Deleted   bool       `json:"deleted"`
		UpdatedBy string     `json:"updated_by"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	if err := json.Unmarshal(requestBody, &softDeleteRequest); err != nil {
		return nil, errors.New("invalid JSON input")
	}

	softDeleteRequest.UpdatedAt = &utils.CurrentTime

	query := `
        UPDATE collection_late_customer_snap
        SET deleted = $1, updated_by = $2, updated_at = $3
        WHERE id = $4
        RETURNING id, snap_date, branch_id, branch_name, customer_id, customer_name,
                  customer_home_address, national_id, customer_address_lat_long,
                  home_phone_number, mobile_phone_number, business_name, business_address,
                  business_phone_number, representative, representative_name, principal,
                  installments, total_late, late_ins_count, late_loans_count, created_at,
                  created_by, updated_at, updated_by, action, deleted, company
    `

	params := []interface{}{
		softDeleteRequest.Deleted,
		softDeleteRequest.UpdatedBy,
		softDeleteRequest.UpdatedAt,
		lateCustomerID,
	}

	Result, err := utils.QueryExecuter(db, "late_customer_result", query, params)
	if err != nil {
		return nil, err
	}

	if len(Result) == 0 {
		return nil, errors.New("failed to soft delete the late customer snap")
	}

	return Result[0], nil
}
