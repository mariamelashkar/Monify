package routers

import (
	"monify/controllers/crud"
	"monify/controllers/hierarchy"
	filter "monify/controllers/late_customers_assign"
	beego "github.com/beego/beego/v2/server/web"
)

// // func init() {
// // 	ns := beego.NewNamespace("/api",
// // 		beego.NSNamespace("/webapp",
// // 			beego.NSNamespace("/v1",
// // 				beego.NSRouter("/allocation_page_start", &controllers.AllocationController{}, "get:AllocationPage"),
// // 				beego.NSRouter("/branch_data", &controllers.AllocationController{}, "post:BranchData"),
// // 				beego.NSRouter("/reallocation", &controllers.ReallocationController{}, "get:ReallocationPage"),
// // 				beego.NSRouter("/reallocation_branch", &controllers.AllocationController{}, "post:ReallocationBranch"),
// // 				beego.NSRouter("/details_page", &controllers.AllocationController{}, "post:DetailsPage"),

// // 			),
// // 		),
// // 	)
// // 	beego.AddNamespace(ns)

// // 	beego.Router("/create/user", &hierarchy.HierarchyController{}, "get:CreateUser")

// // }

// // func init() {
// //     beego.Router("/user", &hierarchy.UserController{}, "post:CreateUser")
// // 	beego.Router("/team", &hierarchy.TeamController{}, "post:CreateTeam")
// //     // beego.Router("/role", &hierarchy.RoleController{}, "post:CreateRole")
// // 	beego.Router("/hierarchy", &hierarchy.HierarchyController{}, "post:CreateHierarchy")
// // 	beego.Router("/assign/hierarchy", &hierarchy.UserHierarchyController{}, "post:AssignHierarchyToUser")
// // 	beego.Router("/team", &hierarchy.TeamController{}, "post:CreateTeam")
// // 	beego.Router("/assign/team", &hierarchy.UserTeamController{}, "post:AssignUserToTeam")
// // 	beego.Router("/assign/customer", &filter.CollectionController{}, "post:AssignSelectedCustomersToUser")

// // }
func init() {
	/*customer routes */
	beego.Router("/late-customer-snap/:id", &controllers.LateCustomersSnapController{}, "get:GetLateCustomerSnap")
	beego.Router("/get/all/late-customer-snaps", &controllers.LateCustomersSnapController{}, "post:GetAllLateCustomersSnap")
	beego.Router("/late-customer-snaps", &controllers.LateCustomersSnapController{}, "post:CreateLateCustomersSnap")
	beego.Router("/late-customer-snaps/:id", &controllers.LateCustomersSnapController{}, "patch:SoftDeleteLateCustomer")
	beego.Router("/late-customer-snaps/:id", &controllers.LateCustomersSnapController{}, "put:UpdateLateCustomersSnap")

	/*user routes */

	beego.Router("/user", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/get/user/hierarchy/:id", &controllers.UserHierarchyController{}, "get:GetUserHierarchy")
	beego.Router("/get/user/:id", &controllers.UserController{}, "get:GetUser")
	beego.Router("/get/users", &controllers.UserController{}, "get:GetAllUsers")
	beego.Router("/user/team/:id", &controllers.UserTeamsController{}, "get:GetUserTeam")
	beego.Router("/user/teams", &controllers.UserTeamsController{}, "get:GetAllUserTeams")

	/* hierarchy */
	beego.Router("/insert/hierarchy", &hierarchy.HierarchyController{}, "post:CreateHierarchy")
	beego.Router("/filter/users", &filter.CollectionController{}, "post:GetUsersAtNextHierarchyLevel")
	beego.Router("/user-team-hierarchy/:id", &controllers.UserTeamHierarchyController{}, "get:GetUserTeamHierarchy")
	beego.Router("/user-team-hierarchies", &controllers.UserTeamHierarchyController{}, "get:GetAllUserTeamHierarchies")
	beego.Router("/hierarchy/delete", &hierarchy.HierarchyController{}, "delete:DeleteHierarchy")
	beego.Router("/get/hierarchies", &controllers.HierarchyController{}, "get:GetAllHierarchies")

}
