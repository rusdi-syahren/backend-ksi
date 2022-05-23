package repository

import (
	"gitlab.com/k1476/scaffolding/src/auth/v1/dto"
	"gitlab.com/k1476/scaffolding/src/shared"
)

// AuthRepository interface
type AuthRepository interface {
	SignUpByPhone(*dto.SignUpByPhoneRequest) shared.Output
}
