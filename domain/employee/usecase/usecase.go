package usecase

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/time-off-system/domain/employee"
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
	repository employee.Repository
}

// NewEmployeeUsecase ...
func NewEmployeeUsecase(repository employee.Repository) employee.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u usecase) Create(tx *gorm.DB, c echo.Context) (interface{}, error) {
	ac := c.(*context.ApplicationContext)
	dto := &vo.CreateRequestDTO{}
	if err := ac.Bind(dto); err != nil {
		logger.Error("failed to bind due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN002, "failed to bind CreateRequestDTO", http.StatusBadRequest, err)
	}

	if err := ac.Validate(dto); err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN003, "validation failed", http.StatusBadRequest, err)
	}

	// - insert into profile table first;
	profileModel := models.Profile{
		FullName:       dto.FullName,
		Email:          dto.Email,
		IdentityNumber: dto.IdentityNumber,
		Address:        dto.Address,
		DateOfBirth:    dto.DateOfBirth,
		PhoneNumber:    dto.PhoneNumber,
		Gender:         dto.Gender,
		MaritalStatus:  dto.MaritalStatus,
		Religion:       dto.Religion,
	}

	profilResponse, err := u.repository.InsertProfile(tx, &profileModel)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEEMP001, "validation failed", http.StatusBadRequest, err)
	}

	// - insert into employee table;
	employeeModel := models.Employee{
		CompanyID:        dto.CompanyID,
		ProfileID:        profilResponse.ID,
		OrganizationName: dto.OrganizationName,
		Position:         dto.Position,
		Level:            dto.Level,
		Status:           dto.Status,
		Branch:           dto.Branch,
		JoinDate:         dto.JoinDate,
	}

	_, err = u.repository.InsertEmployee(tx, &employeeModel)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEEMP002, "validation failed", http.StatusBadRequest, err)
	}

	return dto, nil
}

func (u usecase) Detail(tx *gorm.DB, c echo.Context) (interface{}, error) {
	id := c.QueryParam("employee_id")

	employeeResponse, err := u.repository.Get(tx, id)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEEMP003, "validation failed", http.StatusBadRequest, err)
	}

	return employeeResponse, nil
}

func (u usecase) List(tx *gorm.DB, c echo.Context) (interface{}, error) {
	companyID := c.QueryParam("company_id")
	queryParam := c.QueryParam("query")
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")
	orderBy := c.QueryParam("order_by")
	descendingParam := c.QueryParam("descending")

	params := employee.ListQueryParam{
		CompanyID:  companyID,
		Query:      queryParam,
		Page:       pageParam,
		Limit:      limitParam,
		Offset:     offsetParam,
		OrderBy:    orderBy,
		Descending: descendingParam,
	}

	employeeResponse, err := u.repository.List(tx, params)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEEMP004, "validation failed", http.StatusBadRequest, err)
	}

	return employeeResponse, nil
}
