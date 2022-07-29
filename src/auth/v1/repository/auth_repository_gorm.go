package repository

import (
	"errors"
	"time"

	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/jinzhu/gorm"
)

// AuthRepositoryGorm struct
type AuthRepositoryGorm struct {
	db *gorm.DB
}

// NewAuthRepositoryGorm function
func NewAuthRepositoryGorm(db *gorm.DB) *AuthRepositoryGorm {
	return &AuthRepositoryGorm{db: db}
}

// SignUpByPhone function
func (r *AuthRepositoryGorm) LoginByPhone(params *dto.LoginByPhoneRequest) shared.Output {

	var Login domain.SignUpByPhone

	err := r.db.Raw(`SELECT * FROM user_management
	where phone = ? `, params.MobilePhone).Scan(&Login).Error
	if err != nil {
		err = errors.New("data not found")
		return shared.Output{Error: err}
	}

	return shared.Output{Result: Login}
}

// LoginByPhoneOtp function
func (r *AuthRepositoryGorm) LoginByPhoneOtp(params *dto.LoginByPhoneOtpRequest) shared.Output {

	var Login domain.SignUpByPhone

	err := r.db.Raw(`SELECT * FROM user_management
	where secPatientSignInOtpId = ? `, params.SecPatientSignInOtpId).Scan(&Login).Error
	if err != nil {
		err = errors.New("data not found")
		return shared.Output{Error: err}
	}

	return shared.Output{Result: Login}
}

// CheckUserDelete function
func (r *AuthRepositoryGorm) CheckUserDelete(params *dto.LoginByPhoneRequest) bool {

	var results = struct {
		Total uint
	}{}

	err := r.db.Raw(`SELECT count(*) total FROM security.sec_users
	where mobile_phone = ? and is_active = false and is_deleted = true`, params.MobilePhone).Scan(&results).Error
	if err != nil {
		return false
	}

	if results.Total > 0 {
		return true
	} else {
		return false
	}

}

// CheckUserExist function
func (r *AuthRepositoryGorm) CheckUserExist(params *dto.LoginByPhoneRequest) shared.Output {

	var secUser domain.SecUsers

	err := r.db.Raw(`SELECT  * FROM security.sec_users
	where mobile_phone = ? and user_type_code='patient' and is_active = true and is_deleted = false`, params.MobilePhone).Scan(&secUser).Error
	if err != nil {
		return shared.Output{Error: err}
	}

	if secUser.SecUserId == "" {
		return shared.Output{Error: errors.New("data not found")}
	}

	return shared.Output{Result: secUser}
}

func (r *AuthRepositoryGorm) LoadActiveSecPatient(secUserId string) shared.Output {

	var secUserSignOtp domain.SecPatientSignInOtp

	err := r.db.Raw(`SELECT  * FROM security.sec_patient_sign_in_otps
	where sec_user_id = ?  and is_active = true and is_deleted = false order by created_on desc limit 1 `, secUserId).Scan(&secUserSignOtp).Error
	if err != nil {
		return shared.Output{Error: err}
	}
	// spew.Dump(secUserSignOtp)

	return shared.Output{Result: secUserSignOtp}
}

func (r *AuthRepositoryGorm) CountSms(mobilePhone string) int {

	var results = struct {
		Total uint
	}{}
	r.db.Raw(`SELECT count(*) total FROM sms.sms_logs
	where mobile_phone = ? and  created_on > ?`, mobilePhone, time.Now().Local().Add(-time.Second*1200)).Scan(&results)

	return int(results.Total)
}

func (r *AuthRepositoryGorm) GetSmsLog(mobilePhone string) shared.Output {

	var smsLog domain.SmsLog

	r.db.Raw(`SELECT * FROM sms.sms_logs
	where mobile_phone = ?  order by created_on desc limit 1`, mobilePhone).Scan(&smsLog)

	return shared.Output{Result: smsLog}
}

func (r *AuthRepositoryGorm) SavePatientOtpSignIn(params *domain.SecPatientSignInOtp) shared.Output {

	err := r.db.Save(params).Error
	if err != nil {
		return shared.Output{Error: err, Result: params}
	}

	return shared.Output{Result: params}
}

func (r *AuthRepositoryGorm) UpdatePatientOtpSignIn(params *domain.SecPatientSignInOtp) shared.Output {

	var response domain.SecPatientSignInOtp

	var dbArgsCustomer []interface{}
	sqlCustomer := "update security.sec_patient_sign_in_otps set retry_counter = retry_counter + 1 , created_on = ?"

	dbArgsCustomer = append(dbArgsCustomer, time.Now().UTC())

	sqlCustomer += " where sec_patient_sign_in_otp_id = ? "
	dbArgsCustomer = append(dbArgsCustomer, params.SecPatientSignInOtpID)

	r.db.Raw(sqlCustomer, dbArgsCustomer...).Scan(&response)

	return shared.Output{Result: &response}
}

func (r *AuthRepositoryGorm) SaveSmsLogs(params *domain.SmsLog) shared.Output {

	err := r.db.Save(params).Error
	if err != nil {
		return shared.Output{Error: err, Result: params}
	}

	return shared.Output{Result: params}
}

func (r *AuthRepositoryGorm) SaveSmsLogMessages(params *domain.SmsLogMessage) shared.Output {

	err := r.db.Save(params).Error
	if err != nil {
		return shared.Output{Error: err, Result: params}
	}

	return shared.Output{Result: params}
}
