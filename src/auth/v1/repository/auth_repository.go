package repository

import (
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/shared"
)

// AuthRepository interface
type AuthRepository interface {
	SignUpByPhone(*dto.SignUpByPhoneRequest) shared.Output
}
