package db

// import (
// 	_ "collection/routers"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// )

// func TestInsertAndFetchData(t *testing.T) {
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("Failed to create mock database: %v", err)
// 	}
// 	defer mockDB.Close()

// 	err = SetupMockDB(mock, mockDB)
// 	if err != nil {
// 		t.Fatalf("Failed to set up mock database: %v", err)
// 	}

// 	err = InsertData(mockDB)
// 	if err != nil {
// 		t.Fatalf("Failed to insert data: %v", err)
// 	}

// 	var id int
// 	var name string
// 	err = mockDB.QueryRow("SELECT id, name FROM roles WHERE id = $1", 1).Scan(&id, &name)
// 	if err != nil {
// 		t.Fatalf("Failed to query data: %v", err)
// 	}

// 	if id != 1 || name != "Admin" {
// 		t.Fatalf("Expected to find role 'Admin' with id 1, but found id %d, name %s", id, name)
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Fatalf("There were unfulfilled expectations: %v", err)
// 	}
// }
