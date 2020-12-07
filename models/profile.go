package models

import "time"

type (
	// Profile ...
	Profile struct {
		ID             int64     `db:"id" gorm:"primary_key" json:"id"`
		FullName       string    `db:"full_name" gorm:"column:full_name" json:"full_name"`
		Email          string    `db:"email" gorm:"column:email" json:"email"`
		IdentityNumber string    `db:"identity_number" gorm:"column:identity_number" json:"identity_number"`
		Address        string    `db:"address" gorm:"column:address" json:"address"`
		DateOfBirth    string    `db:"date_of_birth" gorm:"column:date_of_birth" json:"date_of_birth"`
		PhoneNumber    string    `db:"phone_number" gorm:"column:phone_number" json:"phone_number"`
		Gender         string    `db:"gender" gorm:"column:gender" json:"gender"`
		MaritalStatus  string    `db:"marital_status" gorm:"column:marital_status" json:"marital_status"`
		Religion       string    `db:"religion" gorm:"column:religion" json:"religion"`
		CreatedAt      time.Time `db:"created_at" gorm:"column:created_at" json:"created_at"`
		UpdatedAt      time.Time `db:"updated_at" gorm:"column:updated_at" json:"updated_at"`
	}
)
