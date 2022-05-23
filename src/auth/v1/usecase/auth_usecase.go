package usecase

import (
	"gitlab.com/k1476/scaffolding/src/auth/v1/dto"
	"gitlab.com/k1476/scaffolding/src/shared"
)

// AuthUsecase interface
type AuthUsecase interface {
	SignUpByPhone(*dto.SignUpByPhoneRequest) shared.Output
}
