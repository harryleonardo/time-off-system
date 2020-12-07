package usecase

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/time-off-system/domain/employee"
	"github.com/time-off-system/domain/leave"
	"github.com/time-off-system/models"
	"github.com/time-off-system/shared/context"
	"github.com/time-off-system/shared/errors"
	"github.com/time-off-system/shared/log"
	"github.com/time-off-system/shared/vo"
)

var (
	logger = log.NewLog()
)

type usecase struct {
	repository   leave.Repository
	employeeRepo employee.Repository
}

// NewLeaveUsecase ...
func NewLeaveUsecase(repository leave.Repository, employeeRepo employee.Repository) leave.Usecase {
	return &usecase{
		repository:   repository,
		employeeRepo: employeeRepo,
	}
}

func (u usecase) RequestLeave(tx *gorm.DB, c echo.Context) (interface{}, error) {
	ac := c.(*context.ApplicationContext)
	dto := &vo.LeaveRequestDTO{}
	if err := ac.Bind(dto); err != nil {
		logger.Error("failed to bind due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN002, "failed to bind CreateRequestDTO", http.StatusBadRequest, err)
	}

	if err := ac.Validate(dto); err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN003, "validation failed", http.StatusBadRequest, err)
	}

	// - validate first related to quota;
	strEmpID := fmt.Sprintf("%d", dto.EmployeeID)
	employeeData, err := u.employeeRepo.Get(tx, strEmpID)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEEMP003, "validation failed", http.StatusBadRequest, err)
	}

	if dto.Total > employeeData.LeaveQuota {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEOFF001, "validation failed", http.StatusBadRequest, err)
	}

	// - insert leave request into time_off table;
	timeOffModel := models.TimeOff{
		EmployeeID: dto.EmployeeID,
		Type:       dto.Type,
		StartDate:  dto.StartDate,
		EndDate:    dto.EndDate,
		Status:     "REQUESTED",
		Total:      dto.Total,
		Notes:      dto.Notes,
	}

	timeOffResponse, err := u.repository.InsertTimeOff(tx, &timeOffModel)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEOFF002, "failed to insert", http.StatusBadRequest, err)
	}

	return timeOffResponse, nil
}

func (u usecase) ActionLeave(tx *gorm.DB, c echo.Context) (interface{}, error) {
	ac := c.(*context.ApplicationContext)
	dto := &vo.ActionRequestDTO{}
	if err := ac.Bind(dto); err != nil {
		logger.Error("failed to bind due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN002, "failed to bind ActionRequestDTO", http.StatusBadRequest, err)
	}

	if err := ac.Validate(dto); err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN003, "validation failed", http.StatusBadRequest, err)
	}

	// - get time off detail first;
	timeOffDetail, err := u.repository.Get(tx, dto.LeaveID)
	if err != nil {
		logger.Error("failed to get due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEOFF003, "failed to get time off", http.StatusBadRequest, err)
	}

	// - do action;
	if dto.Action == leave.Approved {
		timeOffDetail.Status = leave.Approved
		timeOffDetail.ProcessedBy = dto.ProcessedBy

		quota := timeOffDetail.Employee.LeaveQuota - timeOffDetail.Total

		// - update the leave quota for that employee
		employeeParam := models.Employee{
			ID:         timeOffDetail.EmployeeID,
			LeaveQuota: quota,
		}

		// - query update
		_, err := u.employeeRepo.UpdateByEmpID(tx, &employeeParam)
		if err != nil {
			logger.Error("failed to update employee data due to : ", err)
			return nil, errors.WrapError(errors.ERRCODEOFF004, "failed to update employee data", http.StatusBadRequest, err)
		}

	} else {
		timeOffDetail.Status = dto.Action
	}

	// - update the time off table
	_, err = u.repository.Update(tx, timeOffDetail)
	if err != nil {
		logger.Error("failed to update employee data due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEOFF005, "failed to update employee data", http.StatusBadRequest, err)
	}

	return dto, nil
}

func (u usecase) GetQuotaLeave(tx *gorm.DB, c echo.Context) (interface{}, error) {
	id := c.QueryParam("employee_id")

	employeeResponse, err := u.employeeRepo.Get(tx, id)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEEMP003, "validation failed", http.StatusBadRequest, err)
	}

	return vo.QuotaResponseDTO{
		QuotaLeft: employeeResponse.LeaveQuota,
	}, nil
}

func (u usecase) ListHistory(tx *gorm.DB, c echo.Context) (interface{}, error) {
	id := c.QueryParam("employee_id")
	logger.Info("ID : ", id)

	timeOffListResponse, err := u.repository.List(tx, id)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEOFF006, "failed to listing", http.StatusInternalServerError, err)
	}

	logger.Info("TimeOffListResponse : ", timeOffListResponse)

	return timeOffListResponse, nil
}
