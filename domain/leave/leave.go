package leave

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/time-off-system/models"
)

const (
	Approved = "APPROVED"
)

// Usecase ...
type Usecase interface {
	RequestLeave(tx *gorm.DB, c echo.Context) (interface{}, error)
	ActionLeave(tx *gorm.DB, c echo.Context) (interface{}, error)
	GetQuotaLeave(tx *gorm.DB, c echo.Context) (interface{}, error)
	ListHistory(tx *gorm.DB, c echo.Context) (interface{}, error)
}

// Repository ...
type Repository interface {
	InsertTimeOff(*gorm.DB, *models.TimeOff) (*models.TimeOff, error)
	Get(*gorm.DB, int64) (*models.TimeOff, error)
	Update(*gorm.DB, *models.TimeOff) (*models.TimeOff, error)
	List(db *gorm.DB, employeeID string) (*[]models.TimeOff, error)
}
