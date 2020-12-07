package models

import "time"

type (
	// Company ...
	Company struct {
		ID          int64     `db:"id" gorm:"primary_key" json:"id"`
		Name        string    `db:"name" gorm:"column:name" json:"name"`
		Address     string    `db:"address" gorm:"column:address" json:"address"`
		Majority    string    `db:"majority" gorm:"column:majority" json:"majority"`
		Established string    `db:"established" gorm:"column:established" json:"established"`
		CreatedAt   time.Time `db:"created_at" gorm:"column:created_at" json:"created_at"`
		UpdatedAt   time.Time `db:"updated_at" gorm:"column:updated_at" json:"updated_at"`
	}
)
