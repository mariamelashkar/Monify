package models

// import (
// 	"testing"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGenerateFakeUser(t *testing.T) {
// 	user := GenerateFakeUser()

// 	assert.NotNil(t, user)
// 	assert.NotEmpty(t, user.Username)
// 	assert.NotEmpty(t, user.Email)
// 	assert.NotEmpty(t, user.Password)
// 	assert.NotZero(t, user.Id)
// }

// func TestGenerateFakeBranch(t *testing.T) {
// 	branch := GenerateFakeBranch()

// 	assert.NotNil(t, branch)
// 	assert.NotEmpty(t, branch.Id)
// 	assert.NotEmpty(t, branch.Name)
// 	assert.NotZero(t, branch.BranchCode)
// }

// func TestGenerateFakeFollowers(t *testing.T) {
// 	follower := GenerateFakeFollowers()

// 	assert.NotNil(t, follower)
// 	assert.NotZero(t, follower.Id)
// 	assert.NotNil(t, follower.User)
// 	assert.NotZero(t, follower.Position)
// }

// func TestGenerateFakeVisitCustomerTicket(t *testing.T) {
// 	visitTicket := GenerateFakeVisitCustomerTicket()

// 	assert.NotNil(t, visitTicket)
// 	assert.NotZero(t, visitTicket.Id)
// 	assert.NotNil(t, visitTicket.Ticket)
// 	assert.NotEmpty(t, visitTicket.Lat)
// 	assert.NotEmpty(t, visitTicket.Lang)
// 	assert.NotNil(t, visitTicket.Follower)
// }

// func TestGenerateFakeLateCustomersSnap(t *testing.T) {
// 	lateCustomerSnap := GenerateFakeLateCustomersSnap()

// 	assert.NotNil(t, lateCustomerSnap)
// 	assert.NotZero(t, lateCustomerSnap.Id)
// 	assert.NotEmpty(t, lateCustomerSnap.CustomerName)
// 	assert.NotEmpty(t, lateCustomerSnap.BranchName)
// 	assert.NotZero(t, lateCustomerSnap.Principal)
// }

// func TestGenerateFakeTicketTable(t *testing.T) {
// 	ticketTable := GenerateFakeTicketTable()

// 	assert.NotNil(t, ticketTable)
// 	assert.NotZero(t, ticketTable.Id)
// 	assert.NotNil(t, ticketTable.CustomerTicket)
// 	assert.NotNil(t, ticketTable.FollowUser)
// 	assert.NotEmpty(t, ticketTable.UpdateState)
// }
