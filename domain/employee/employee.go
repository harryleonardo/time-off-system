package employee

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/time-off-system/models"
)

type (
	// ListQueryParam ...
	ListQueryParam struct {
		CompanyID  string `query:"company_id"`
		Query      string `query:"query"`
		Page       string `query:"page"`
		Limit      string `query:"limit"`
		Offset     string `query:"offset"`
		OrderBy    string `query:"order_by"`
		Descending string `query:"descending"`
	}
)

// Usecase ...
type Usecase interface {
	Create(tx *gorm.DB, c echo.Context) (interface{}, error)
	Detail(tx *gorm.DB, c echo.Context) (interface{}, error)
	List(tx *gorm.DB, c echo.Context) (interface{}, error)
}

// Repository ...
type Repository interface {
	InsertProfile(*gorm.DB, *models.Profile) (*models.Profile, error)
	InsertEmployee(*gorm.DB, *models.Employee) (*models.Employee, error)
	Get(*gorm.DB, string) (*models.Employee, error)
	List(*gorm.DB, ListQueryParam) (*[]models.Employee, error)
	UpdateByEmpID(*gorm.DB, *models.Employee) (*models.Employee, error)
}
