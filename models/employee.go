package models

import "time"

type (
	// Employee ...
	Employee struct {
		ID               int64     `db:"id" gorm:"primary_key" json:"id"`
		CompanyID        int64     `db:"company_id" gorm:"column:company_id" json:"company_id"`
		ProfileID        int64     `db:"profile_id" gorm:"column:profile_id" json:"profile_id"`
		OrganizationName string    `db:"organization_name" gorm:"column:organization_name" json:"organization_name"`
		Position         string    `db:"position" gorm:"column:position" json:"position"`
		Level            string    `db:"level" gorm:"column:level" json:"level"`
		Status           string    `db:"status" gorm:"column:status" json:"status"`
		Branch           string    `db:"branch" gorm:"column:branch" json:"branch"`
		JoinDate         string    `db:"join_date" gorm:"column:join_date" json:"join_date"`
		ApprovalID       int64     `db:"approval_id" gorm:"column:approval_id" json:"approval_id"`
		LeaveQuota       int64     `db:"leave_quota" gorm:"column:leave_quota" json:"leave_quota"`
		Profile          Profile   `json:"profile"`
		Company          Company   `json:"company"`
		CreatedAt        time.Time `db:"created_at" gorm:"column:created_at" json:"created_at"`
		UpdatedAt        time.Time `db:"updated_at" gorm:"column:updated_at" json:"updated_at"`
	}
)
