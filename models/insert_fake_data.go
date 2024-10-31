package models

// import (
// 	"fmt"
// 	"log"
// 	"github.com/beego/beego/v2/client/orm"
// )

// func InsertFakeData() {
// 	o := orm.NewOrm()
// 	tx, err := o.Begin()
// 	if err != nil {
// 		log.Fatalf("Failed to start transaction: %v", err)
// 	}

// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 			log.Fatalf("Transaction failed and rolled back: %v", r)
// 		} else {
// 			err = tx.Commit()
// 			if err != nil {
// 				log.Fatalf("Failed to commit transaction: %v", err)
// 			} else {
// 				fmt.Println("Transaction committed successfully!")
// 			}
// 		}
// 	}()

// 	// Branch
// 	fakeBranch := GenerateFakeBranch()
// 	if _, err := tx.Insert(fakeBranch); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake branch: %v", err))
// 	}

// 	// User
// 	fakeUser := GenerateFakeUser()
// 	if _, err := tx.Insert(fakeUser); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake user: %v", err))
// 	}

// 	// Followers
// 	fakeFollowers := GenerateFakeFollowers()
// 	fakeFollowers.User = fakeUser
// 	if _, err := tx.Insert(fakeFollowers); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake followers: %v", err))
// 	}

// 	// TicketTable
// 	fakeTicketTable := GenerateFakeTicketTable()
// 	fakeTicketTable.FollowUser = fakeFollowers
// 	if _, err := tx.Insert(fakeTicketTable); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake ticket table: %v", err))
// 	}

// 	// CurrentFollower
// 	fakeCurrentFollower := GenerateFakeCurrentFollower()
// 	fakeCurrentFollower.Ticket = fakeTicketTable
// 	fakeCurrentFollower.FollowUser = fakeFollowers
// 	if _, err := tx.Insert(fakeCurrentFollower); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake current follower: %v", err))
// 	}

// 	// FollowerAction
// 	fakeFollowerAction := GenerateFakeFollowerAction()
// 	fakeFollowerAction.Ticket = fakeTicketTable
// 	fakeFollowerAction.Follower = fakeFollowers
// 	if _, err := tx.Insert(fakeFollowerAction); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake follower action: %v", err))
// 	}

// 	// ActionResult
// 	fakeActionResult := GenerateFakeActionResult()
// 	fakeActionResult.Action = fakeFollowerAction
// 	if _, err := tx.Insert(fakeActionResult); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake action result: %v", err))
// 	}

// 	// UserBranch
// 	fakeUserBranch := GenerateFakeUserBranch()
// 	fakeUserBranch.User = fakeUser
// 	fakeUserBranch.Branch = fakeBranch
// 	if _, err := tx.Insert(fakeUserBranch); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake user branch: %v", err))
// 	}

// 	// VisitCustomerTicket
// 	fakeVisitCustomerTicket := GenerateFakeVisitCustomerTicket()
// 	fakeVisitCustomerTicket.Ticket = fakeTicketTable
// 	fakeVisitCustomerTicket.Follower = fakeFollowers
// 	if _, err := tx.Insert(fakeVisitCustomerTicket); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake visit customer ticket: %v", err))
// 	}

// 	// LateCustomersSnap
// 	fakeLateCustomersSnap := GenerateFakeLateCustomersSnap()
// 	if _, err := tx.Insert(fakeLateCustomersSnap); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake late customers snap: %v", err))
// 	}

// 	// LateLoansSnap
// 	fakeLateLoansSnap := GenerateFakeLateLoansSnap()
// 	if _, err := tx.Insert(fakeLateLoansSnap); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake late loans snap: %v", err))
// 	}

// 	// ResultChoose
// 	fakeResultChoose := GenerateFakeResultChoose()
// 	if _, err := tx.Insert(fakeResultChoose); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake result choose: %v", err))
// 	}

// 	// StreamLocation
// 	fakeStreamLocation := GenerateFakeStreamLocation()
// 	if _, err := tx.Insert(fakeStreamLocation); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake stream location: %v", err))
// 	}

// 	// Actions
// 	fakeActions := GenerateFakeActions()
// 	if _, err := tx.Insert(fakeActions); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake actions: %v", err))
// 	}

// 	//Teams
// 	fakeTeam := GenerateFakeTeam()
// 	if _, err := tx.Insert(fakeTeam); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake team: %v", err))
// 	}
// 	// Roles
// 	fakeRole := GenerateFakeRole()
// 	if _, err := tx.Insert(fakeRole); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake role: %v", err))
// 	}

// 	//UserRole
// 	fakeUserRole := GenerateFakeUserRole()
// 	fakeUserRole.UserId = fakeUser
// 	fakeUserRole.RoleId = fakeRole
// 	if _, err := tx.Insert(fakeUserRole); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake user role: %v", err))
// 	}

// 	// UserTeams
// 	fakeUserTeam := GenerateFakeUserTeam()
// 	fakeUserTeam.UserId = fakeUser
// 	fakeUserTeam.TeamId = fakeTeam
// 	if _, err := tx.Insert(fakeUserTeam); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake user team: %v", err))
// 	}
// 	// TeamRole
// 	fakeTeamRole := GenerateFakeTeamRole()
// 	fakeTeamRole.TeamId = fakeTeam
// 	fakeTeamRole.RoleId = fakeRole
// 	if _, err := tx.Insert(fakeTeamRole); err != nil {
// 		panic(fmt.Sprintf("Failed to insert fake team role: %v", err))
// 	}
// 	// UserTeamRoles
// // 	fakeUserTeamRole := GenerateFakeUserTeamRole()
// // 	fakeUserTeamRole.UserId = fakeUser
// // 	fakeUserTeamRole.TeamId = fakeTeam
// // 	fakeUserTeamRole.RoleId = fakeRole
// // 	if _, err := tx.Insert(fakeUserTeamRole); err != nil {
// // 		panic(fmt.Sprintf("Failed to insert fake user team role: %v", err))
// // 	}
// // 	fmt.Println("All records inserted successfully!")
// }
