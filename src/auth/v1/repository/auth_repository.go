package repository

import (
	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/shared"
)

// AuthRepository interface
type AuthRepository interface {
	SignUpByPhone(*dto.SignUpByPhoneRequest) shared.Output
	LoginByPhone(*dto.LoginByPhoneRequest) shared.Output
	LoginByPhoneOtp(*dto.LoginByPhoneOtpRequest) shared.Output
	CheckUserDelete(*dto.LoginByPhoneRequest) bool
	CheckUserExist(*dto.LoginByPhoneRequest) shared.Output
	LoadActiveSecPatient(string) shared.Output
	CountSms(string) int
	GetSmsLog(string) shared.Output
	SavePatientOtpSignIn(*domain.SecPatientSignInOtp) shared.Output
	SaveSmsLogs(*domain.SmsLog) shared.Output
	SaveSmsLogMessages(*domain.SmsLogMessage) shared.Output
	UpdatePatientOtpSignIn(*domain.SecPatientSignInOtp) shared.Output
}
