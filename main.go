package main

import (
	"monify/models"
	"log"
	_ "monify/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/postgres"
	_ "github.com/lib/pq"
)

func main() {
	dbDriver, _ := beego.AppConfig.String("db_driver")
	dbUser, _ := beego.AppConfig.String("db_user")
	dbPassword, _ := beego.AppConfig.String("db_password")
	dbHost, _ := beego.AppConfig.String("db_host")
	dbPort, _ := beego.AppConfig.String("db_port")
	db, _ := beego.AppConfig.String("db_")
	orm.Debug = true

	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", dbDriver, "user="+dbUser+" password="+dbPassword+" host="+dbHost+" port="+dbPort+" db="+db+" sslmode=disable")

	orm.RegisterModel(
		new(models.Branch),
		new(models.LateCustomersSnap),
		new(models.LateLoansSnap),
		new(models.Role),
		new(models.User),
		new(models.Team),
		new(models.Actions),
		new(models.ActionResult),
		new(models.Hierarchy),
		new(models.CurrentFollower),
		new(models.FollowerAction),
		new(models.ResultChoose),
		new(models.StreamLocation),
		new(models.TeamHierarchy),
		new(models.UserTeam),
		new(models.VisitCustomerTicket),
		new(models.TicketTable),
		new(models.UserHierarchy),
		new(models.UserTeamHierarchy),
		new(models.ChangeLog),
	)

	err := models.InitializeChangeLogSystem()
	if err != nil {
		log.Fatalf("Error initializing change log system: %v", err)
	}
	orm.RunSyncdb("default", false, true)

	models.InsertInitialCustomerData()
	models.InsertInitialUsersData()
	models.InsertInitialLateLoansData()
	beego.Run()
}
