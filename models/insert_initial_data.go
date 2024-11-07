package models

// NOTE: All data in this file is fictional and used solely for testing or development purposes.
//       The names, addresses, and other details are fake and do not correspond to any real individuals
//  

import (
	"log"
	"time"
	"github.com/beego/beego/v2/client/orm"
)

func InsertInitialCustomerData() {
	o := orm.NewOrm()
	createdBy := "2"
	updatedBy := &createdBy
	initialLateCustomersSnap := []LateCustomersSnap{
		{
			Id:                     "2",
			SnapDate:               time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			BranchId:               "611657bea4",
			BranchName:             "القاهرة - المعادي",
			CustomerId:             "61250e0562c2fe92d21deb9c",
			CustomerName:           "ياسمين علي محمود حسن",
			CustomerHomeAddress:    "15ش التحرير - الدقي",
			NationalId:             "29009290107218",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01234567890",
			BusinessName:           "شركة النور",
			BusinessAddress:        "شركة النور",
			BusinessPhoneNumber:    "",
			Representative:         "",
			RepresentativeName:     "هشام عبدالرحمن - موظف",
			Principal:              24000,
			Installments:           40000,
			TotalLate:              12000,
			LateInsCount:           50,
			LateLoansCount:         5,
			Company:                "",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
		{
			Id:                     "3",
			SnapDate:               time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - حلوان",
			CustomerId:             "6125fe641908b15457bbcd0f",
			CustomerName:           "سعيد حسن محمد",
			CustomerHomeAddress:    "شارع 23 - مدينة نصر",
			NationalId:             "29407290100359",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01112345678",
			BusinessName:           "مطعم الفيروز",
			BusinessAddress:        "مطعم الفيروز",
			BusinessPhoneNumber:    "",
			Representative:         "",
			RepresentativeName:     "منى عادل - مسؤولة مبيعات",
			Principal:              9000,
			Installments:           15000,
			TotalLate:              5000,
			LateInsCount:           20,
			LateLoansCount:         2,
			Company:                "",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
	}

	for _, customerSnap := range initialLateCustomersSnap {
		_, err := o.Insert(&customerSnap)
		if err != nil {
			log.Printf("Failed to insert customer snap: %v\n", err)
		} else {
			log.Printf("Successfully inserted customer snap: %s\n", customerSnap.CustomerName)
		}
	}
}

func InsertInitialLateLoansData() {
	o := orm.NewOrm()
	createdBy := "2"
	updatedBy := &createdBy
	initialLateLoansSnap := []LateLoansSnap{
		{
			Id:                     "1",
			CheckDate:              time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			LateAmount:             1300.50,
			InsCount:               6,
			Lid:                    "loan10",
			Lsid:                   "lsid10",
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - المعادي",
			CustomerId:             "61250e0562c2fe92d21deb9c",
			CustomerName:           "ياسمين علي محمود حسن",
			Key:                    "key10",
			RepresentativeId:       "rep10",
			LoanOfficer:            "محمد عبد السلام",
			Principal:              5500.00,
			IssueDate:              1642051200,
			LoanKey:                "loankey10",
			TotalInstallmentSum:    7000.00,
			CustomerHomeAddress:    "15ش التحرير - الدقي",
			NationalId:             "29009290107218",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01234567890",
			BusinessName:           "شركة النور",
			BusinessAddress:        "شركة النور",
			BusinessPhoneNumber:    "",
			CurrentRep:             "current_rep10",
			CurrentRepName:         "هشام عبدالرحمن - موظف",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
		{
			Id:                     "2",
			CheckDate:              time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			LateAmount:             1400.75,
			InsCount:               4,
			Lid:                    "loan11",
			Lsid:                   "lsid11",
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - حلوان",
			CustomerId:             "6125fe641908b15457bbcd0f",
			CustomerName:           "سعيد حسن محمد",
			Key:                    "key11",
			RepresentativeId:       "rep11",
			LoanOfficer:            "أحمد سمير",
			Principal:              3200.00,
			IssueDate:              1642051200,
			LoanKey:                "loankey11",
			TotalInstallmentSum:    4500.00,
			CustomerHomeAddress:    "شارع 23 - مدينة نصر",
			NationalId:             "29407290100359",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01112345678",
			BusinessName:           "مطعم الفيروز",
			BusinessAddress:        "مطعم الفيروز",
			BusinessPhoneNumber:    "",
			CurrentRep:             "current_rep11",
			CurrentRepName:         "منى عادل - مسؤولة مبيعات",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
	}

	for _, loanSnap := range initialLateLoansSnap {
		_, err := o.Insert(&loanSnap)
		if err != nil {
			log.Printf("Failed to insert late loan snap: %v\n", err)
		} else {
			log.Printf("Successfully inserted late loan snap for customer: %s\n", loanSnap.CustomerName)
		}
	}
}

func InsertInitialUsersData() {
	o := orm.NewOrm()
	createdBy := "admin"
	updatedBy := &createdBy
	intialroles := []Role{
		{
			Id:        "1",
			Name:      "مدير",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "2",
			Name:      "مجموعة الرقابة",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			Actions:   "created",
			Deleted:   false,
		},
	}
	for _, roles := range intialroles {
		_, err := o.Insert(&roles)
		if err != nil {
			log.Printf("Failed to insert role: %v\n", err)
		} else {
			log.Printf("Successfully inserted role: %s\n", roles.Name)
		}
	}

	initialUsers := []User{
		{
			Id:          "1",
			Password:    "hashed_password_1",
			Username:    "omar_ahmed",
			Name:        "عمر أحمد عبد العزيز",
			Email:       "omar@example.com",
			IsActive:    true,
			Prefix:      "Mr.",
			Position:    "مدير",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "2",
			Password:    "hashed_password_2",
			Username:    "sara_hassan",
			Name:        "سارة حسن عبد الله",
			Email:       "sara@example.com",
			IsActive:    true,
			Prefix:      "Ms.",
			Position:    "مراجعة",
			RoleId:      &intialroles[1],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			Actions:     "created",
			Deleted:     false,
		},
	}

	for _, user := range initialUsers {
		_, err := o.Insert(&user)
		if err != nil {
			log.Printf("Failed to insert user: %v\n", err)
		} else {
			log.Printf("Successfully inserted user: %s\n", user.Username)
		}
	}

	initialHierarchies := []Hierarchy{
		{
			Id:        "11",
			Level:     1,
			Name:      "مدير عام",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "22",
			Level:     2,
			Name:      "مشرف منطقة",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			Actions:   "created",
			Deleted:   false,
		},
	}

	for _, hierarchy := range initialHierarchies {
		_, err := o.Insert(&hierarchy)
		if err != nil {
			log.Printf("Failed to insert hierarchy: %v\n", err)
		} else {
			log.Printf("Successfully inserted hierarchy: %s\n", hierarchy.Name)
		}
	}

	initialTeams := []Team{
		{Id: "1", Name: "القاهرة - مدينة نصر", Region: "القاهرة", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "2", Name: "الجيزة - الدقي", Region: "الجيزة", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
	}

	for _, team := range initialTeams {
		_, err := o.Insert(&team)
		if err != nil {
			log.Printf("Failed to insert team: %v\n", err)
		} else {
			log.Printf("Successfully inserted team: %s\n", team.Name)
		}
	}
}
