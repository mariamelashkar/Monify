package models

import (
	// "collection/utils"
	// "encoding/json"
	// "errors"
	"time"
)

type TeamHierarchy struct {
	Id          string     `orm:"column(id);pk;size(70)"`
	TeamId      *Team      `orm:"rel(fk);column(team_id)"`
	HeirarchyId *Hierarchy `orm:"rel(fk);column(role_id)"`
	CreatedAt   time.Time  `orm:"column(created_at);type(datetime);auto_now_add"`
	CreatedBy   string     `orm:"column(created_by)"`
	UpdatedBy   *string    `orm:"column(updated_by);size(70)"`
	UpdatedAt   *time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
	Actions     string     `orm:"column(action);size(200);default(created)"`
	Deleted     bool       `orm:"column(deleted);default(false)"`
}

func (tr *TeamHierarchy) TableName() string {
	return "team_roles"
}

// // ---------- CRUD Operations for team_heirarchy Model ----------//
