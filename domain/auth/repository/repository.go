package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/time-off-system/domain/auth"
	"github.com/time-off-system/models"
)

type repoHandler struct {
}

// NewAuthRepository ...
func NewAuthRepository() auth.Repository {
	return &repoHandler{}
}

func (r *repoHandler) GetCredential(db *gorm.DB, param *models.Auth) (*models.Auth, error) {
	auth := &models.Auth{}

	if err := db.Where("user_name = ? AND password = ?", param.UserName, param.Password).First(auth).Error; err != nil {
		return nil, fmt.Errorf("%s for user_name %s", err, param.UserName)
	}

	return auth, nil
}
