package usecase

import (
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/dto"
	"github.com/rusdi-syahren/backend-ksi/src/shared"
)

// AuthUsecase interface
type AuthUsecase interface {
	// Sigin Patient
	SignInByPhonePassword(*dto.LoginByPhoneRequest) shared.OutputV1
	SignInByPhoneOtp(*dto.LoginByPhoneOtpRequest) shared.OutputV1
}
