package models

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"

// 	"github.com/jaswdr/faker/v2"
// )

// func randomTimeBetween(start, end time.Time) time.Time {
// 	diff := end.Sub(start)
// 	randomDuration := time.Duration(rand.Int63n(int64(diff)))
// 	return start.Add(randomDuration)
// }

// func randomFloat64(min, max float64) float64 {
// 	return min + rand.Float64()*(max-min)
// }

// func randomStringWithSize(n int) string {
// 	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	result := make([]byte, n)
// 	for i := range result {
// 		result[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(result)
// }

// func GenerateFakeActionResult() *ActionResult {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &ActionResult{
// 		Id:              fake.IntBetween(1, 1000),
// 		Action:          GenerateFakeFollowerAction(),
// 		Sentimental:     fake.IntBetween(0, 100),
// 		Desc:            fake.Lorem().Paragraph(1),
// 		Result:          GenerateFakeResultChoose(),
// 		CollectedAmount: randomFloat64(0, 10000),
// 		NextFollowDate:  randomTimeBetween(past, now),
// 		AddDate:         randomTimeBetween(past, now),
// 		UpdateAt:        randomTimeBetween(past, now),
// 		UpdateUser:      GenerateFakeUser(),
// 	}
// }

// func GenerateFakeActions() *Actions {
//     fake := faker.New()

//     return &Actions{
//         Id:         fake.IntBetween(1, 1000),
//         ActionName: fake.Lorem().Word(),
//         ActionType: fake.Lorem().Word(), // Generates a random word for ActionType
//     }
// }

// func GenerateFakeBranch() *Branch {
//     fake := faker.New()

//     return &Branch{
// 		Id:         fake.IntBetween(1, 1000),
//         BranchCode:    fake.IntBetween(1000, 9999),
//         Governate:     fake.Address().City(),
//         Longitude:     fake.Address().Longitude(),
//         Latitude:      fake.Address().Latitude(),
//         Address:       fake.Address().Address(),
//         FlexValue:     fake.Lorem().Word(),
//         BranchCompany: fake.Company().Name(),
//     }
// }


// func GenerateFakeCurrentFollower() *CurrentFollower {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &CurrentFollower{
// 		Id:         fake.IntBetween(1, 1000),
// 		Ticket:     GenerateFakeTicketTable(),
// 		FollowUser: GenerateFakeFollowers(),
// 		AddDate:    randomTimeBetween(past, now),
// 		AddUser:    GenerateFakeUser(),
// 	}
// }

// func GenerateFakeFollowerAction() *FollowerAction {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &FollowerAction{
// 		Id:         fake.IntBetween(1, 1000),
// 		Ticket:     GenerateFakeTicketTable(),
// 		Follower:   GenerateFakeFollowers(),
// 		Action:     GenerateFakeActions(),
// 		Lat:        fmt.Sprintf("%f", fake.Address().Latitude()),
// 		Lang:       fmt.Sprintf("%f", fake.Address().Longitude()),
// 		AddDate:    randomTimeBetween(past, now),
// 		UpdateDate: randomTimeBetween(past, now),
// 		Date:       randomTimeBetween(past, now),
// 	}
// }

// func GenerateFakeFollowers() *Followers {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &Followers{
// 		Id:         fake.IntBetween(1, 1000),
// 		User:       GenerateFakeUser(),
// 		Position:   fake.IntBetween(0, 10),
// 		State:      fake.IntBetween(0, 10),
// 		AddAt:      randomTimeBetween(past, now),
// 		AddUser:    GenerateFakeUser(),
// 		UpdatedAt:  randomTimeBetween(past, now),
// 		UpdateUser: GenerateFakeUser(),
// 		DeviceId:   fake.Internet().MacAddress(),
// 	}
// }

// func truncateString(str string, maxLength int) string {
//     if len(str) > maxLength {
//         return str[:maxLength]
//     }
//     return str
// }

// func GenerateFakeLateCustomersSnap() *LateCustomersSnap {
//     fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

//     return &LateCustomersSnap{
//         Id:                     fake.IntBetween(1, 1000),
//         SnapDate:               randomTimeBetween(past, now),
//         BranchId:               truncateString(randomStringWithSize(20), 20),
//         BranchName:             truncateString(fake.Company().Name(), 20),
//         CustomerId:             truncateString(randomStringWithSize(20), 20),
//         CustomerName:           truncateString(fake.Person().Name(), 20),
//         CustomerHomeAddress:    truncateString(fake.Address().Address(), 20),
//         NationalId:             truncateString(randomStringWithSize(14), 14),
//         CustomerAddressLatLong: fmt.Sprintf("%f,%f", fake.Address().Latitude(), fake.Address().Longitude()),
//         HomePhoneNumber:        truncateString(fake.Phone().Number(), 20),
//         MobilePhoneNumber:      truncateString(fake.Phone().Number(), 20),
//         BusinessName:           truncateString(fake.Company().Name(), 20),
//         BusinessAddress:        truncateString(fake.Address().Address(), 20),
//         BusinessPhoneNumber:    truncateString(fake.Phone().Number(), 20),
//         Representative:         truncateString(randomStringWithSize(20), 20),
//         RepresentativeName:     truncateString(fake.Person().Name(), 20),
//         Principal:              randomFloat64(1000, 50000),
//         Installments:           randomFloat64(100, 10000),
//         TotalLate:              randomFloat64(0, 5000),
//         LateInsCount:           fake.IntBetween(0, 10),
//         LateLoansCount:         fake.IntBetween(0, 10),
//     }
// }
// func GenerateFakeLateLoansSnap() *LateLoansSnap {
//     fake := faker.New()
//     now := time.Now()
//     past := now.AddDate(-1, 0, 0)

//     return &LateLoansSnap{
//         Id:                     fake.IntBetween(1, 1000),
//         CheckDate:              randomTimeBetween(past, now),
//         LateAmount:             randomFloat64(100, 10000),
//         InsCount:               fake.IntBetween(1, 10),
//         Lid:                    truncateString(randomStringWithSize(255), 20), 
//         Lsid:                   truncateString(randomStringWithSize(255), 20), 
//         BranchId:               truncateString(randomStringWithSize(255), 20), 
//         BranchName:             truncateString(fake.Company().Name(), 20), 
//         CustomerId:             truncateString(randomStringWithSize(255), 20),
//         CustomerName:           truncateString(fake.Person().Name(), 20), 
//         Key:                    truncateString(randomStringWithSize(255), 20), 
//         RepresentativeId:       truncateString(randomStringWithSize(255), 20), 
//         LoanOfficer:            truncateString(fake.Person().Name(), 20), 
//         Principal:              randomFloat64(1000, 50000),
//         IssueDate:              fake.Int64Between(past.Unix(), now.Unix()),
//         LoanKey:                truncateString(randomStringWithSize(255), 20),
//         TotalInstallmentSum:    randomFloat64(1000, 50000),
//         CustomerHomeAddress:    truncateString(fake.Address().Address(), 20), 
//         NationalId:             truncateString(randomStringWithSize(255), 20), 
//         CustomerAddressLatLong: fmt.Sprintf("%f,%f", fake.Address().Latitude(), fake.Address().Longitude()),
//         HomePhoneNumber:        truncateString(fake.Phone().Number(), 20), 
//         MobilePhoneNumber:      truncateString(fake.Phone().Number(), 20), 
//         BusinessName:           truncateString(fake.Company().Name(), 20), 
//         BusinessAddress:        truncateString(fake.Address().Address(), 20), 
//         BusinessPhoneNumber:    truncateString(fake.Phone().Number(), 20), 
//         CurrentRep:             truncateString(randomStringWithSize(255), 20), 
//         CurrentRepName:         truncateString(fake.Person().Name(), 20),
//     }
// }


// func GenerateFakeResultChoose() *ResultChoose {
// 	fake := faker.New()

// 	return &ResultChoose{
// 		Id:         fake.IntBetween(1, 1000),
// 		ResultName: fake.Lorem().Sentence(3),
// 		Category:   fake.IntBetween(0, 10),
// 	}
// }

// func GenerateFakeStreamLocation() *StreamLocation {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &StreamLocation{
// 		Id:       fake.IntBetween(1, 1000),
// 		Date:     randomTimeBetween(past, now),
// 		DeviceId: fake.Internet().MacAddress(),
// 		Lat:      fmt.Sprintf("%f", fake.Address().Latitude()),
// 		Lang:     fmt.Sprintf("%f", fake.Address().Longitude()),
// 	}
// }

// func GenerateFakeTicketTable() *TicketTable {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &TicketTable{
// 		Id:                  fake.IntBetween(1, 1000),
// 		CustomerTicket:      GenerateFakeLateCustomersSnap(),
// 		ClosePrincipal:      randomFloat64(1000, 50000),
// 		CloseInstallments:   randomFloat64(100, 10000),
// 		CloseTotalLate:      randomFloat64(0, 5000),
// 		CloseLateInsCount:   fake.IntBetween(0, 10),
// 		CloseLateLoansCount: fake.IntBetween(0, 10),
// 		FollowUser:          GenerateFakeFollowers(),
// 		FollowState:         fake.IntBetween(0, 10),
// 		AddUser:             GenerateFakeFollowers(),
// 		AddDate:             randomTimeBetween(past, now),
// 		UpdateUser:          GenerateFakeUser(),
// 		UpdatedAt:           randomTimeBetween(past, now),
// 		UpdateState:         fake.Lorem().Sentence(5),
// 	}
// }

// func GenerateFakeUserBranch() *UserBranch {
// 	fake := faker.New()

// 	return &UserBranch{
// 		Id:     fake.IntBetween(1, 1000),
// 		User:   GenerateFakeUser(),
// 		Branch: GenerateFakeBranch(),
// 	}
// }

// func GenerateFakeUserRole() *UserRole{
// 	fake := faker.New()
// 	return &UserRole{
// 		Id:    fake.IntBetween(1, 1000),
// 		UserId:  GenerateFakeUser(),
// 		RoleId: GenerateFakeRole(),
// 	}
// }

// func GenerateFakeUser() *User {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	firstName := fake.Person().FirstName()
// 	lastName := fake.Person().LastName()
// 	username := fmt.Sprintf("%s.%s", firstName, lastName)

// 	prefixes := []string{"Mr.", "Mrs.", "Ms.", "Dr."}
// 	randomPrefix := prefixes[rand.Intn(len(prefixes))]

// 	return &User{
// 		Id:          int64(fake.IntBetween(1, 1000)),
// 		Password:    fake.Internet().Password(),
// 		LastLogin:   randomTimeBetween(past, now),
// 		IsSuperuser: fake.Bool(),
// 		Username:    username,
// 		Name:        firstName,
// 		Email:       fake.Internet().Email(),
// 		IsActive:    fake.Bool(),
// 		DateJoined:  randomTimeBetween(past, now),
// 		Prefix:      randomPrefix,
// 		Position:    fake.Company().JobTitle(),
// 	}
// }

// func GenerateFakeVisitCustomerTicket() *VisitCustomerTicket {
// 	fake := faker.New()
// 	now := time.Now()
// 	past := now.AddDate(-1, 0, 0)

// 	return &VisitCustomerTicket{
// 		Id:        fake.IntBetween(1, 1000),
// 		Ticket:    GenerateFakeTicketTable(),
// 		Lat:       fmt.Sprintf("%f", fake.Address().Latitude()),
// 		Lang:      fmt.Sprintf("%f", fake.Address().Longitude()),
// 		Timestamp: randomTimeBetween(past, now),
// 		Follower:  GenerateFakeFollowers(),
// 		Date:      randomTimeBetween(past, now),
// 		DeviceId:  fake.Internet().MacAddress(),
// 	}
// }

// func GenerateFakeTeamRole() *TeamRole {
// 	fake := faker.New()
// 	return &TeamRole{
// 		Id:        fake.IntBetween(1, 1000),
// 		TeamId:    GenerateFakeTeam(),
// 		RoleId:    GenerateFakeRole(),
// 		CreatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// 		UpdatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// 	}
// }
// // func GenerateFakeUserTeamRole() *UserTeamRole {
// // 	fake := faker.New()
// // 	return &UserTeamRole{
// // 		Id:        fake.IntBetween(1, 1000),
// // 		UserId:    GenerateFakeUser(),
// // 		TeamId:    GenerateFakeTeam(),
// // 		RoleId:    GenerateFakeRole(),
// // 		CreatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// // 		UpdatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// // 	}
// // }

// func GenerateFakeUserTeam() *UserTeam {
// 	fake := faker.New()
// 	return &UserTeam{
// 		Id:        fake.IntBetween(1, 1000),
// 		UserId:    GenerateFakeUser(),
// 		TeamId:    GenerateFakeTeam(),
// 		CreatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// 		UpdatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// 	}
// }

// func GenerateFakeTeam() *Team {
// 	fake := faker.New()
// 	return &Team{
// 		Id:        fake.IntBetween(1, 1000),
// 		Name:      fake.Company().Name(),
// 		CreatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// 		UpdatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
// 	}
// }
// func GenerateFakeRole() *Role {
//     fake := faker.New()
//     jobTitles := []string{"Engineer", "Manager", "Administrator", "Specialist", "Coordinator", "Analyst"}
//     randomJobTitle := jobTitles[rand.Intn(len(jobTitles))]

//     return &Role{
//         Id:        fake.IntBetween(1, 1000),
//         Name:      randomJobTitle,
//         CreatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
//         UpdatedAt: randomTimeBetween(time.Now().AddDate(-1, 0, 0), time.Now()),
//     }
// }


