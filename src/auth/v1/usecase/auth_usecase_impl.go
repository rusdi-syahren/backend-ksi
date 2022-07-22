package usecase

import (
	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/auth/v1/repository"
	"github.com/Klinisia/backend-ksi/src/shared"
)

// AuthUsecaseImpl struct
type AuthUsecaseImpl struct {
	AuthRepositoryWrite repository.AuthRepository
}

// NewAuthUsecaseImpl function
func NewAuthUsecaseImpl(AuthRepositoryWrite repository.AuthRepository) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		AuthRepositoryWrite: AuthRepositoryWrite,
	}
}

// SignUpByPhone function
func (u *AuthUsecaseImpl) SignUpByPhone(filter *dto.SignUpByPhoneRequest) shared.Output {

	// assign Auth process
	signUp := u.AuthRepositoryWrite.SignUpByPhone(filter)
	if signUp.Error != nil {
		return shared.Output{Error: signUp.Error}
	}

	response := signUp.Result.(domain.SignUpByPhone)

	return shared.Output{Result: response}
}
