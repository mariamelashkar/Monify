package utils

import (
    "database/sql"
    "fmt"
    "log"

    beego "github.com/beego/beego/v2/server/web"
)

func GetDBConnection() (*sql.DB, error) {
    dbDriver, _ := beego.AppConfig.String("db_driver")
    dbUser, _ := beego.AppConfig.String("db_user")
    dbPassword, _ := beego.AppConfig.String("db_password")
    dbHost, _ := beego.AppConfig.String("db_host")
    dbPort, _ := beego.AppConfig.Int("db_port")
	dbName, _ := beego.AppConfig.String("db_name")

    log.Printf("Connecting to database: %s at host: %s:%d as user: %s", dbName, dbHost, dbPort, dbUser)

    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
        dbUser, dbPassword, dbName, dbHost, dbPort)

    db, err := sql.Open(dbDriver, connStr)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
