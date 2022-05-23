package usecase

import (
	"gitlab.com/k1476/scaffolding/src/auth/v1/domain"
	"gitlab.com/k1476/scaffolding/src/auth/v1/dto"
	"gitlab.com/k1476/scaffolding/src/auth/v1/repository"
	"gitlab.com/k1476/scaffolding/src/shared"
	"gitlab.com/k1476/scaffolding/src/shared/external"
)

// AuthUsecaseImpl struct
type AuthUsecaseImpl struct {
	AuthRepositoryWrite repository.AuthRepository
	ext                 external.Notif
}

// NewAuthUsecaseImpl function
func NewAuthUsecaseImpl(AuthRepositoryWrite repository.AuthRepository, ext external.Notif) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		AuthRepositoryWrite: AuthRepositoryWrite,
		ext:                 ext,
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
