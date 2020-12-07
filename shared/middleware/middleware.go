package middleware

import (
	"github.com/labstack/echo"
	PkgErrors "github.com/pkg/errors"
	SharedErrors "github.com/time-off-system/shared/errors"
	"github.com/time-off-system/shared/log"
	SharedVO "github.com/time-off-system/shared/vo"
)

type (
	// Middleware ...
	Middleware interface {
		ErrorHandler() func(error, echo.Context)
	}

	middlewares struct{}
)

var (
	logger = log.NewLog()

	// ErrorHandler ...
	ErrorHandler = func(err error, c echo.Context) {
		code, message, status := "PYS-GEN-001", "Internal Server Error", 500

		switch original := PkgErrors.Cause(err).(type) {
		case *SharedErrors.ErrorWrapper:
			code = original.Code
			message = original.Message
			status = original.StatusCode
		default:
			logger.Errorf("%+v", err)
		}

		json := &SharedVO.ErrorResponse{
			Message:   message,
			ErrorCode: code,
			Status:    status,
		}

		c.JSON(status, json)
	}
)

func (m *middlewares) ErrorHandler() func(error, echo.Context) {
	return ErrorHandler
}

// NewMiddleware ...
func NewMiddleware() Middleware {
	return &middlewares{}
}
