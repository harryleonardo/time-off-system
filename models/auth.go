package models

import "time"

type (
	// Auth ...
	Auth struct {
		ID         int64     `db:"id" gorm:"primary_key" json:"id"`
		EmployeeID int64     `db:"employee_id" gorm:"column:employee_id" json:"employee_id"`
		UserName   string    `db:"user_name" gorm:"column:user_name" json:"user_name"`
		Password   string    `db:"password" gorm:"column:password" json:"password"`
		Role       string    `db:"role" gorm:"column:role" json:"role"`
		CreatedAt  time.Time `db:"created_at" gorm:"column:created_at" json:"created_at"`
		UpdatedAt  time.Time `db:"updated_at" gorm:"column:updated_at" json:"updated_at"`
	}
)
