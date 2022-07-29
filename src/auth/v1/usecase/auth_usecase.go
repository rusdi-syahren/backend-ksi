package usecase

import (
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/shared"
)

// AuthUsecase interface
type AuthUsecase interface {
	// Sigin Patient
	SignInByPhonePassword(*dto.LoginByPhoneRequest) shared.OutputV1
	SignInByPhoneOtp(*dto.LoginByPhoneOtpRequest) shared.Output
}
