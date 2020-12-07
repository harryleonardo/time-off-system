package context

import (
	"github.com/fgrosse/goldi"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/time-off-system/shared/config"
	"github.com/time-off-system/shared/log"
)

var (
	guid   = xid.New()
	logger = log.NewLog()
)

// ApplicationContext ...
type ApplicationContext struct {
	echo.Context
	Container    *goldi.Container
	MysqlSession *gorm.DB
	Config       config.ImmutableConfig
}

func (c *ApplicationContext) GetRequestID() string {
	rid := c.Request().Header.Get("X-Request-Id")

	if rid == "" {
		rid = guid.String()
	}

	return rid
}

func (c *ApplicationContext) WithError(message, code string, status int, err error) error {
	switch original := errors.Cause(err).(type) {
	default:
		logger.Errorf("%+v", original)
	}

	return c.JSON(status, &ErrorResponse{
		Message:   message,
		ErrorCode: code,
		Status:    status,
	})
}

func (c *ApplicationContext) WithEmptySuccess(message string, status int) error {
	return c.WithSuccess(message, status, make(map[string]string))
}

func (c *ApplicationContext) WithSuccess(message string, status int, data interface{}) error {
	return c.JSON(status, &SuccessResponse{
		Message: message,
		Data:    data,
		Status:  status,
	})
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

type ErrorResponse struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    int         `json:"status"`
	ErrorCode string      `json:"error_code"`
}
