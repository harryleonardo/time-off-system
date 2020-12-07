package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/time-off-system/domain/employee"
	"github.com/time-off-system/shared/context"
	"github.com/time-off-system/shared/log"
)

var (
	logger = log.NewLog()
)

type employeeHandler struct {
	usecase employee.Usecase
}

// EmployeeHandler ...
func EmployeeHandler(e *echo.Echo, usecase employee.Usecase) {
	handler := employeeHandler{
		usecase: usecase,
	}

	e.POST("/api/employee", handler.Create)
	e.GET("/api/employee", handler.Detail)
	e.POST("/api/employee/list", handler.List)
}

func (handler employeeHandler) Create(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.Create(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}

func (handler employeeHandler) Detail(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.Detail(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}

func (handler employeeHandler) List(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.List(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}
