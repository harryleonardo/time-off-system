package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/time-off-system/domain/auth"
	"github.com/time-off-system/shared/context"
	"github.com/time-off-system/shared/log"
)

var (
	logger = log.NewLog()
)

type authHandler struct {
	usecase auth.Usecase
}

// AuthHandler ...
func AuthHandler(e *echo.Echo, usecase auth.Usecase) {
	handler := authHandler{
		usecase: usecase,
	}

	e.POST("/api/login", handler.Login)
}

func (handler authHandler) Login(e echo.Context) error {
	ac := e.(*context.ApplicationContext)
	session := ac.MysqlSession
	tx := session.Begin()

	res, err := handler.usecase.Login(tx, e)
	if err != nil {
		tx.Rollback()
		messageError := fmt.Sprintf("%s", err)
		logger.Error(messageError)
		return err
	}

	tx.Commit()
	return ac.WithSuccess("success", http.StatusOK, res)
}
