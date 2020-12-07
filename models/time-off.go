package models

import "time"

type (
	// TimeOff ...
	TimeOff struct {
		ID          int64     `db:"id" gorm:"primary_key" json:"id"`
		EmployeeID  int64     `db:"employee_id" gorm:"column:employee_id" json:"employee_id"`
		Type        string    `db:"type" gorm:"column:type" json:"type"`
		StartDate   time.Time `db:"start_date" gorm:"column:start_date" json:"start_date"`
		EndDate     time.Time `db:"end_date" gorm:"column:end_date" json:"end_date"`
		Status      string    `db:"status" gorm:"column:status" json:"status"`
		Total       int64     `db:"total" gorm:"column:total" json:"total"`
		Notes       string    `db:"notes" gorm:"column:notes" json:"notes"`
		ProcessedBy int64     `db:"processed_by" gorm:"column:processed_by" json:"processed_by"`
		Employee    Employee  `json:"employee"`
		CreatedAt   time.Time `db:"created_at" gorm:"column:created_at" json:"created_at"`
		UpdatedAt   time.Time `db:"updated_at" gorm:"column:updated_at" json:"updated_at"`
	}
)
