package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/time-off-system/domain/employee"
	"github.com/time-off-system/models"
	"github.com/time-off-system/shared/common"
	"github.com/time-off-system/shared/errors"
)

type repository struct {
}

func NewEmployeeRepository() employee.Repository {
	return &repository{}
}

func (r *repository) InsertProfile(db *gorm.DB, profile *models.Profile) (*models.Profile, error) {
	if err := db.Create(profile).Error; err != nil {
		return nil, errors.New("failed to insert profile because " + err.Error())
	}

	return profile, nil
}

func (r *repository) InsertEmployee(db *gorm.DB, employee *models.Employee) (*models.Employee, error) {
	if err := db.Create(employee).Error; err != nil {
		return nil, errors.New("failed to insert employee because " + err.Error())
	}

	return employee, nil
}

func (r *repository) Get(db *gorm.DB, employeeID string) (*models.Employee, error) {
	employeeModel := models.Employee{}

	err := db.
		Preload("Profile").
		Preload("Company").
		Where("id = ?", employeeID).
		First(&employeeModel).Error

	if err != nil {
		return &employeeModel, err
	}

	return &employeeModel, nil
}

func (r *repository) UpdateByEmpID(db *gorm.DB, employeeParam *models.Employee) (*models.Employee, error) {
	employeeModel := models.Employee{}

	err := db.Where("id = ?", employeeParam.ID).First(&employeeModel).Error
	if err != nil {
		return &employeeModel, err
	}

	if err := db.Model(employeeModel).
		UpdateColumns(models.Employee{
			LeaveQuota: employeeParam.LeaveQuota,
		}).Error; err != nil {
		return &employeeModel, err
	}

	return &employeeModel, nil
}

func (r *repository) List(db *gorm.DB, params employee.ListQueryParam) (*[]models.Employee, error) {
	employeeListModel := []models.Employee{}

	// - build query param;
	var err error

	queryBuilder := db.
		Joins("LEFT JOIN profiles ON employees.profile_id = profiles.id")

	// - query
	if params.Query != common.StringDefault {
		qParam := fmt.Sprintf("%s%s%s", "%", params.Query, "%")
		queryBuilder = queryBuilder.Where("profiles.full_name LIKE ? OR profiles.email LIKE ?", qParam, qParam)
	}

	// - limit
	if params.Limit != common.StringDefault {
		queryBuilder = queryBuilder.Limit(params.Limit)
	}

	// - page && offset
	if params.Page != common.StringDefault {
		offsetVal := "0"
		if params.Offset != common.StringDefault {
			offsetVal = params.Offset
		}
		queryBuilder = queryBuilder.Offset(offsetVal)
	}

	// - order by
	if params.OrderBy != common.StringDefault {
		queryBuilder = queryBuilder.Order("employee.created_at DESC")
	}

	err = queryBuilder.
		Preload("Profile").
		Preload("Company").
		Where("company_id = ?", params.CompanyID).
		Find(&employeeListModel).Error

	if err != nil {
		return &employeeListModel, err
	}

	return &employeeListModel, nil
}
