package models

import (
	"time"

	// "github.com/beego/beego/v2/client/orm"
)

type TicketTable struct {
    Id         string        `orm:"column(id);pk;size(70)"`
	CustomerTicket      *LateCustomersSnap `orm:"column(customer_ticket);rel(fk)"`
	ClosePrincipal      float64            `orm:"column(close_principal);null"`
	CloseInstallments   float64            `orm:"column(close_installments);null"`
	CloseTotalLate      float64            `orm:"column(close_total_late);null"`
	CloseLateInsCount   int                `orm:"column(close_late_ins_count);null"`
	CloseLateLoansCount int                `orm:"column(close_late_loans_count);null"`
	//FollowUser          *Followers         `orm:"column(follow_user);rel(fk);null"`
	FollowState         int                `orm:"column(follow_state);type(int);default(0)"`
//	AddUser             *Followers         `orm:"column(add_user);rel(fk)"`
	AddDate             time.Time          `orm:"column(add_date);type(date);auto_now_add"`
	UpdateUser          *User       `orm:"column(update_user);rel(fk);null"`
	UpdatedAt           time.Time          `orm:"column(updated_at);type(timestamp);auto_now_add;null"`
	UpdateState         string             `orm:"column(update_state);type(text);null"`
}

func (T *TicketTable) TableName() string {
	return "ticket_table"
}

// func init() {
// 	orm.RegisterModel(new(TicketTable))
// }
