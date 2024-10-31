package db

// import (
// 	"database/sql"
// 	"github.com/DATA-DOG/go-sqlmock"
// )

// func SetupMockDB(mock sqlmock.Sqlmock, db *sql.DB) error {

// 	mock.ExpectBegin()

// 	mock.ExpectExec("INSERT INTO roles \\(id, name, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)").
// 		WithArgs(1, "Admin", "2024-08-31 12:00:00", "2024-08-31 12:00:00").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO branches \\(id, name, branch_code, governate, longitude, latitude, address, flex_value, branch_company\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9\\)").
// 		WithArgs(1, "Main Branch", 1001, "Cairo", 31.2357, 30.0444, "123 Main St", "FLEX100", "Tech Corp").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO teams \\(id, name, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)").
// 		WithArgs(1, "Development Team", "2024-08-31 12:00:00", "2024-08-31 12:00:00").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO team_role \\(id, team_id, role_id, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)").
// 		WithArgs(1, 1, 1, "2024-08-31 12:00:00", "2024-08-31 12:00:00").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO users \\(id, password, last_login, is_superuser, username, first_name, email, is_active, date_joined, prefix, position\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8, \\$9, \\$10, \\$11\\)").
// 		WithArgs(1, "hashedpassword", "2024-08-31 12:00:00", true, "johndoe", "John Doe", "john.doe@example.com", true, "2024-08-31 12:00:00", "Mr.", "Developer").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO user_branch \\(id, user_id, branch_id\\) VALUES \\(\\$1, \\$2, \\$3\\)").
// 		WithArgs(1, 1, 1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO user_role \\(id, user_id, role_id\\) VALUES \\(\\$1, \\$2, \\$3\\)").
// 		WithArgs(1, 1, 1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO users_teams \\(id, user_id, team_id, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)").
// 		WithArgs(1, 1, 1, "2024-08-31 12:00:00", "2024-08-31 12:00:00").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	mock.ExpectExec("INSERT INTO user_team_role \\(id, user_id, team_id, role_id, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6\\)").
// 		WithArgs(1, 1, 1, 1, "2024-08-31 12:00:00", "2024-08-31 12:00:00").
// 		WillReturnResult(sqlmock.NewResult(1, 1))

// 	// Expect commit transaction
// 	mock.ExpectCommit()

// 	return nil
// }
