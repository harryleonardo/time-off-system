package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/time-off-system/domain/leave"
	"github.com/time-off-system/models"
	"github.com/time-off-system/shared/common"
	"github.com/time-off-system/shared/log"
)

var (
	logger = log.NewLog()
)

type repository struct{}

// NewLeaveRepository ...
func NewLeaveRepository() leave.Repository {
	return &repository{}
}

func (r *repository) InsertTimeOff(db *gorm.DB, timeOff *models.TimeOff) (*models.TimeOff, error) {
	if err := db.Create(timeOff).Error; err != nil {
		return nil, errors.New("failed to insert timeOff because " + err.Error())
	}

	return timeOff, nil
}

func (r *repository) Get(db *gorm.DB, timeOffID int64) (*models.TimeOff, error) {
	timeOffModel := models.TimeOff{}

	err := db.
		Preload("Employee").
		Where("id = ?", timeOffID).
		First(&timeOffModel).Error

	if err != nil {
		return &timeOffModel, err
	}

	return &timeOffModel, nil
}

func (r *repository) Update(db *gorm.DB, timeOffModel *models.TimeOff) (*models.TimeOff, error) {
	if err := db.Model(timeOffModel).
		UpdateColumns(models.TimeOff{
			Status:      timeOffModel.Status,
			ProcessedBy: timeOffModel.ProcessedBy,
		}).Error; err != nil {
		return timeOffModel, err
	}

	return timeOffModel, nil
}

func (r *repository) List(db *gorm.DB, employeeID string) (*[]models.TimeOff, error) {
	timeOffListModel := []models.TimeOff{}
	var err error

	queryBuilder := db.
		Joins("LEFT JOIN employees ON time_offs.employee_id = employees.id")

	if employeeID != common.StringDefault {
		queryBuilder = queryBuilder.Where("employee_id = ?", employeeID)
	}

	err = queryBuilder.
		Preload("Employee").
		Find(&timeOffListModel).Error

	if err != nil {
		return &timeOffListModel, err
	}

	return &timeOffListModel, nil
}
