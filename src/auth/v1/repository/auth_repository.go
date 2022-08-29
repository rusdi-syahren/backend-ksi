package repository

import (
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/domain"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/dto"
	"github.com/rusdi-syahren/backend-ksi/src/shared"
)

// AuthRepository interface
type AuthRepository interface {
	// login patient
	LoginByPhone(*dto.LoginByPhoneRequest) shared.Output
	LoginByPhoneOtp(*dto.LoginByPhoneOtpRequest) shared.Output
	CheckUserDelete(*dto.LoginByPhoneRequest) int
	CheckUserExist(*dto.LoginByPhoneRequest) shared.Output
	LoadActiveSecPatient(string) shared.Output
	CountSms(string) int
	GetSmsLog(string) shared.Output
	SavePatientOtpSignIn(*domain.SecPatientSignInOtp) shared.Output
	SaveSmsLogs(*domain.SmsLog) shared.Output
	SaveSmsLogMessages(*domain.SmsLogMessage) shared.Output
	UpdatePatientOtpSignIn(*domain.SecPatientSignInOtp) shared.Output

	// login by otp
	FindBySecPatientSignInOtp(*dto.LoginByPhoneOtpRequest) shared.Output
	GetSecUserByID(string) shared.Output
}
