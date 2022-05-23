package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"gitlab.com/k1476/scaffolding/src/auth/v1/domain"
	"gitlab.com/k1476/scaffolding/src/auth/v1/dto"
	"gitlab.com/k1476/scaffolding/src/shared"
)

// AuthRepositoryGorm struct
type AuthRepositoryGorm struct {
	db *gorm.DB
}

// NewAuthRepositoryGorm function
func NewAuthRepositoryGorm(db *gorm.DB) *AuthRepositoryGorm {
	return &AuthRepositoryGorm{db: db}
}

// SignUpByPhone function
func (r *AuthRepositoryGorm) SignUpByPhone(params *dto.SignUpByPhoneRequest) shared.Output {

	var signUp domain.SignUpByPhone

	err := r.db.Save(&signUp).Error
	if err != nil {
		return shared.Output{Error: err, Result: signUp}
	}

	return shared.Output{Result: signUp}
}

// UpdateAuth function
func (r *AuthRepositoryGorm) GetSales(salesID int) shared.Output {
	var Sales domain.SignUpByPhone

	err := r.db.Raw(`SELECT * FROM user_management
	where id = ? `, salesID).Scan(&Sales).Error
	if err != nil {
		err = errors.New("data not found")
		return shared.Output{Error: err}
	}

	return shared.Output{Result: Sales}
}
