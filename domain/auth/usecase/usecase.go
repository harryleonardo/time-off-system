package usecase

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/time-off-system/domain/auth"
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
	repository auth.Repository
}

// NewAuthUsecase ...
func NewAuthUsecase(repository auth.Repository) auth.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u usecase) Login(tx *gorm.DB, c echo.Context) (interface{}, error) {
	ac := c.(*context.ApplicationContext)
	dto := &vo.AuthenticationRequestDTO{}
	if err := ac.Bind(dto); err != nil {
		logger.Error("failed to bind due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN002, "failed to bind AuthenticationRequestDTO", http.StatusBadRequest, err)
	}

	if err := ac.Validate(dto); err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEVPGEN003, "validation failed", http.StatusBadRequest, err)
	}

	// - check repository data;
	authParam := &models.Auth{
		UserName: dto.UserName,
		Password: dto.Password,
	}

	authResponse, err := u.repository.GetCredential(tx, authParam)
	if err != nil {
		logger.Error("failed to validate due to : ", err)
		return nil, errors.WrapError(errors.ERRCODEAUTH001, "validation failed", http.StatusBadRequest, err)
	}

	return &vo.AuthenticationResponseDTO{
		UserName: authResponse.UserName,
		Role:     authResponse.Role,
	}, nil
}
