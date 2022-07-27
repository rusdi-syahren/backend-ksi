package usecase

import (
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/shared"
)

// AuthUsecase interface
type AuthUsecase interface {
	SignUpByPhone(*dto.SignUpByPhoneRequest) shared.Output
	// login Patient
	LoginByPhonePassword(*dto.LoginByPhoneRequest) shared.OutputV1
	LoginByPhoneOtp(*dto.LoginByPhoneOtpRequest) shared.Output
}
