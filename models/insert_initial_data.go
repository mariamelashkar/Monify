package models

import (
	"collection/utils"
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
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - حدائق القبه - مشروعي",
			CustomerId:             "61250e0562c2fe92d21deb9b",
			CustomerName:           "شريف رافت سيد احمد ابراهيم",
			CustomerHomeAddress:    "10ش سليم الديب عين شمس",
			NationalId:             "29009290107217",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01113331765",
			BusinessName:           "شركة المجد",
			BusinessAddress:        "شركة المجد",
			BusinessPhoneNumber:    "",
			Representative:         "",
			RepresentativeName:     "امل محمد عبدالحميد عوض - اخصائية",
			Principal:              24767,
			Installments:           41208,
			TotalLate:              12152,
			LateInsCount:           57,
			LateLoansCount:         6,
			Company:                "",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
		{
			Id:                     "3",
			SnapDate:               time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - حدائق القبه - مشروعي",
			CustomerId:             "6125fe641908b15457bbcd0e",
			CustomerName:           "احمد محمد احمد محمد بيومي",
			CustomerHomeAddress:    "18ش الجوهري الزيتون القاهره",
			NationalId:             "29407290100358",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01095960918",
			BusinessName:           "مندوب بشركة هرفست",
			BusinessAddress:        "مندوب بشركة هرفست",
			BusinessPhoneNumber:    "",
			Representative:         "",
			RepresentativeName:     "بيشوى مجدى عوض الله توفيق - رئيس مجموعه",
			Principal:              9499,
			Installments:           16776,
			TotalLate:              5396,
			LateInsCount:           22,
			LateLoansCount:         2,
			Company:                "",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
		{
			Id:                     "4",
			SnapDate:               time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - حدائق القبه - مشروعي",
			CustomerId:             "612639dd62c2fe92d21deb9e",
			CustomerName:           "فاديه عيد عبدالكريم عبدالرحمن",
			CustomerHomeAddress:    "79ش حسني الشرابيه-القاهره",
			NationalId:             "26904020103781",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01279036762",
			BusinessName:           "معاش",
			BusinessAddress:        "معاش",
			BusinessPhoneNumber:    "",
			Representative:         "",
			RepresentativeName:     "ساره طارق ادوارد شفيق - اخصائية",
			Principal:              4390,
			Installments:           6528,
			TotalLate:              1904,
			LateInsCount:           7,
			LateLoansCount:         1,
			Company:                "",
			CreatedBy:              createdBy,
			UpdatedBy:              updatedBy,
		},
		{
			Id:                     "5",
			SnapDate:               time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			BranchId:               "611657be16d4647c21bec3a4",
			BranchName:             "القاهرة - حدائق القبه - مشروعي",
			CustomerId:             "61263b4062c2fe92d21deb9f",
			CustomerName:           "هاجر مجدي محمود محمد النجار",
			CustomerHomeAddress:    "79ش حسني الشرابيه-القاهره",
			NationalId:             "29709110104586",
			CustomerAddressLatLong: "",
			HomePhoneNumber:        "",
			MobilePhoneNumber:      "01227610722",
			BusinessName:           "بائعة محل ملابس",
			BusinessAddress:        "بائعة بمحل ملابس",
			BusinessPhoneNumber:    "",
			Representative:         "",
			RepresentativeName:     "ساره طارق ادوارد شفيق - اخصائية",
			Principal:              6010,
			Installments:           8952,
			TotalLate:              2984,
			LateInsCount:           8,
			LateLoansCount:         1,
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
			CheckDate:             time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			LateAmount:            1500.75,
			InsCount:              5,
			Lid:                   "loan1",
			Lsid:                  "lsid1",
			BranchId:              "611657be16d4647c21bec3a4",
			BranchName:            "القاهرة - حدائق القبه - مشروعي",
			CustomerId:            "61250e0562c2fe92d21deb9b",
			CustomerName:          "شريف رافت سيد احمد ابراهيم",
			Key:                   "key1",
			RepresentativeId:      "rep1",
			LoanOfficer:           "المدير المالي",
			Principal:             5000.00,
			IssueDate:             1642051200,
			LoanKey:               "loankey1",
			TotalInstallmentSum:   6500.00,
			CustomerHomeAddress:   "10ش سليم الديب عين شمس",
			NationalId:            "29009290107217",
			CustomerAddressLatLong: "",
			HomePhoneNumber:       "",
			MobilePhoneNumber:     "01113331765",
			BusinessName:          "شركة المجد",
			BusinessAddress:       "شركة المجد",
			BusinessPhoneNumber:   "",
			CurrentRep:            "current_rep1",
			CurrentRepName:        "امل محمد عبدالحميد عوض - اخصائية",
			CreatedBy:             createdBy,
			UpdatedBy:             updatedBy,
		},
		{
			Id:                     "5",
			CheckDate:             time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			LateAmount:            1200.50,
			InsCount:              3,
			Lid:                   "loan2",
			Lsid:                  "lsid2",
			BranchId:              "611657be16d4647c21bec3a4",
			BranchName:            "القاهرة - حدائق القبه - مشروعي",
			CustomerId:            "6125fe641908b15457bbcd0e",
			CustomerName:          "احمد محمد احمد محمد بيومي",
			Key:                   "key2",
			RepresentativeId:      "rep2",
			LoanOfficer:           "المدير المالي",
			Principal:             3000.00,
			IssueDate:             1642051200,
			LoanKey:               "loankey2",
			TotalInstallmentSum:   4200.00,
			CustomerHomeAddress:   "18ش الجوهري الزيتون القاهره",
			NationalId:            "29407290100358",
			CustomerAddressLatLong: "",
			HomePhoneNumber:       "",
			MobilePhoneNumber:     "01095960918",
			BusinessName:          "مندوب بشركة هرفست",
			BusinessAddress:       "مندوب بشركة هرفست",
			BusinessPhoneNumber:   "",
			CurrentRep:            "current_rep2",
			CurrentRepName:        "بيشوى مجدى عوض الله توفيق - رئيس مجموعه",
			CreatedBy:             createdBy,
			UpdatedBy:             updatedBy,
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
			Name:      "Managerial",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "2",
			Name:      "Collection",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
			Actions:   "created",
			Deleted:   false,
		},
	}
	for _, roles := range intialroles {
		_, err := o.Insert(&roles)
		if err != nil {
			log.Printf("Failed to insert customer snap: %v\n", err)
		} else {
			log.Printf("Successfully inserted customer snap: %s\n", roles.Id)
		}
	}

	initialUsers := []User{
		{
			Id:          "1",
			Password:    "hashed_password_1",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "abanoub_george",
			Name:        "ابانوب جورج حليم نسيم - Test",
			Email:       "abanoub@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Developer",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "2",
			Password:    "hashed_password_2R",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "ahmed_gomaa",
			Name:        "احمد جمعة عبدالتواب مفتاح - عضو رقابه",
			Email:       "ahmed@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    createdBy,
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   "system",
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "3",
			Password:    "hashed_password_3",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "mostafa_desouky",
			Name:        "مصطفى سمير دسوقى محمد دسوقى - صراف",
			Email:       "mostafa@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Cashier",
			RoleId:      &intialroles[1], 
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "4",
			Password:    "hashed_password_4",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "sabry_halim",
			Name:        "صبرى حليم عدلى فهيم- صراف",
			Email:       "sabry@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Cashier",
			RoleId:      &intialroles[1],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "56",
			Password:    "hashed_password_5",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "walid_salah",
			Name:        "وليد صالح محمد بدوى",
			Email:       "walid@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Team Member",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "66",
			Password:    "hashed_password_6",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "irene_sobhy",
			Name:        "ايريني صبحي اسحاق متياس - مراجع sme",
			Email:       "irene@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Ms.",
			Position:    "Reviewer",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "71",
			Password:    "hashed_password_7",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "hany_nadi",
			Name:        "هاني نادي نادر رياض",
			Email:       "hany@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Team Member",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "84",
			Password:    "hashed_password_8",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "mohamed_abdelrahman",
			Name:        "محمد عبدالرحيم عبدالسميع عبيد - عضو رقابة",
			Email:       "mohamed@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Reviewer",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "944",
			Password:    "hashed_password_9",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "abdulrahman_mohamed",
			Name:        "عبدالرحمن محمد عبدالله حامد - مدخل بيانات",
			Email:       "abdulrahman@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Mr.",
			Position:    "Data Entry",
			RoleId:      &intialroles[1],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:          "100",
			Password:    "hashed_password_10",
			LastLogin:   time.Time{},
			IsSuperuser: false,
			Username:    "samar_ahmed",
			Name:        "سمر احمد عاطف محمود",
			Email:       "samar@example.com",
			IsActive:    true,
			DateJoined:  time.Now(),
			Prefix:      "Ms.",
			Position:    "Team Member",
			RoleId:      &intialroles[0],
			CreatedAt:   time.Now(),
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			UpdatedAt:   nil,
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
			Name:      "collection manager",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "22",
			Level:     2,
			Name:      "area supervisor",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "3",
			Level:     3,
			Name:      "teamleader",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "52",
			Level:     4,
			Name:      "supervisor",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
			Actions:   "created",
			Deleted:   false,
		},
		{
			Id:        "27",
			Level:     5,
			Name:      "collection officer",
			CreatedAt: time.Now(),
			CreatedBy: createdBy,
			UpdatedBy: updatedBy,
			UpdatedAt: nil,
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
	initialUserHierarchy := []UserHierarchy{
		{
			Id:        "55",
			UserId:      &initialUsers[0],
			HierarchyId: &initialHierarchies[0],
			CreatedAt:   utils.CurrentTime,
			UpdatedAt:   nil,
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:        "21",
			UserId:      &initialUsers[1],
			HierarchyId: &initialHierarchies[1],
			CreatedAt:   utils.CurrentTime,
			UpdatedAt:   nil,
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			Actions:     "created",
			Deleted:     false,
		},
		{
			Id:        "31",
			UserId:      &initialUsers[2],
			HierarchyId: &initialHierarchies[0],
			CreatedAt:   utils.CurrentTime,
			UpdatedAt:   nil,
			CreatedBy:   createdBy,
			UpdatedBy:   updatedBy,
			Actions:     "created",
			Deleted:     false,
		},
	}

	for _, userHierarchy := range initialUserHierarchy {
		_, err := o.Insert(&userHierarchy)
		if err != nil {
			log.Printf("Failed to insert user hierarchy: %v\n", err)
		} else {
			log.Printf("Successfully inserted user hierarchy for user: %v\n", &userHierarchy.UserId)
		}
	}

	initialTeams := []Team{
		{Id: "1", Name: "القاهرة - حدائق القبه - مشروعي", Region: "القاهرة", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "2", Name: "المنيا - المنيا - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "3", Name: "المنيا - بنى مزار - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "4", Name: "المنيا - ابوقرقاص - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "5", Name: "المنيا - سمالوط - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "6", Name: "المنيا - ملوى - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "7", Name: "المنيا - العدوة - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "8", Name: "المنيا - مطاى - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "9", Name: "المنيا - مغاغة - مشروعي", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "10", Name: "المنيا - الشيخ فضل - مشروعى", Region: "المنيا", CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
	}

	for _, team := range initialTeams {
		_, err := o.Insert(&team)
		if err != nil {
			log.Printf("Failed to insert team: %v\n", err)
		} else {
			log.Printf("Successfully inserted team: %s\n", team.Id)
		}
	}
	initialUserTeams := []UserTeam{
		{Id: "1", UserId: &initialUsers[0], TeamId: &initialTeams[0], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "2", UserId: &initialUsers[1], TeamId: &initialTeams[1], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "3", UserId: &initialUsers[2], TeamId: &initialTeams[2], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "4", UserId: &initialUsers[3], TeamId: &initialTeams[3], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "5", UserId: &initialUsers[4], TeamId: &initialTeams[4], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "6", UserId: &initialUsers[5], TeamId: &initialTeams[5], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "7", UserId: &initialUsers[6], TeamId: &initialTeams[6], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "8", UserId: &initialUsers[7], TeamId: &initialTeams[7], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "9", UserId: &initialUsers[8], TeamId: &initialTeams[8], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
		{Id: "10", UserId: &initialUsers[9], TeamId: &initialTeams[9], CreatedAt: time.Now(), CreatedBy: createdBy, Actions: "created", Deleted: false},
	}

	for _, userTeam := range initialUserTeams {
		_, err := o.Insert(&userTeam)
		if err != nil {
			log.Printf("Failed to insert user-team relationship: %v\n", err)
		} else {
			log.Printf("Successfully inserted user-team relationship for user: %s and team: %s\n", userTeam.UserId.Username, userTeam.TeamId.Name)
		}
	}

	initialUserTeamHierarchy := []UserTeamHierarchy{
		{
			Id:           "1",
			UserId:       &initialUsers[0], 
			TeamId:       &initialTeams[0], 
			HierarchyId:  &initialHierarchies[0], 
			CreatedAt:    time.Now(),
			CreatedBy:    createdBy,
			UpdatedBy:    updatedBy,
			UpdatedAt:    nil,
			Actions:      "created",
			Deleted:      false,
		},
		{
			Id:           "2",
			UserId:       &initialUsers[1],
			TeamId:       &initialTeams[1],
			HierarchyId:  &initialHierarchies[1],
			CreatedAt:    time.Now(),
			CreatedBy:    createdBy,
			UpdatedBy:    updatedBy,
			UpdatedAt:    nil,
			Actions:      "created",
			Deleted:      false,
		},
		{
			Id:           "3",
			UserId:       &initialUsers[2],
			TeamId:       &initialTeams[2],
			HierarchyId:  &initialHierarchies[0],
			CreatedAt:    time.Now(),
			CreatedBy:    createdBy,
			UpdatedBy:    updatedBy,
			UpdatedAt:    nil,
			Actions:      "created",
			Deleted:      false,
		},
	}

	for _, userTeamHierarchy := range initialUserTeamHierarchy {
		_, err := o.Insert(&userTeamHierarchy)
		if err != nil {
			log.Printf("Failed to insert user-team hierarchy: %v\n", err)
		} else {
			log.Printf("Successfully inserted user-team hierarchy for user: %s, team: %s, hierarchy: %s\n", userTeamHierarchy.UserId.Username, userTeamHierarchy.TeamId.Name, userTeamHierarchy.HierarchyId.Name)
		}
	}

}
