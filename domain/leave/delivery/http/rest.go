package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/time-off-system/domain/leave"
	"github.com/time-off-system/shared/context"
	"github.com/time-off-system/shared/log"
)

var (
	logger = log.NewLog()
)

type leaveHandler struct {
	usecase leave.Usecase
}

// LeaveHandler ...
func LeaveHandler(e *echo.Echo, usecase leave.Usecase) {
	handler := leaveHandler{
		usecase: usecase,
	}

	e.POST("/api/request-leave", handler.RequestLeave)
	e.POST("/api/request-leave/action", handler.RequestLeaveAction)
	e.GET("/api/request-leave/quota", handler.LeaveQuota)
	e.GET("/api/request-leave/history", handler.History)
}

func (handler leaveHandler) RequestLeave(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.RequestLeave(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}

func (handler leaveHandler) RequestLeaveAction(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.ActionLeave(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}

func (handler leaveHandler) LeaveQuota(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.GetQuotaLeave(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}

func (handler leaveHandler) History(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.ListHistory(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}
