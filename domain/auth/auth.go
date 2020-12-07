package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/time-off-system/models"
)

// Usecase ...
type Usecase interface {
	Login(tx *gorm.DB, c echo.Context) (interface{}, error)
}

// Repository ...
type Repository interface {
	GetCredential(*gorm.DB, *models.Auth) (*models.Auth, error)
}
