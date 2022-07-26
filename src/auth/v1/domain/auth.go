package domain

import "time"

// SignUpByPhone struct
type SignUpByPhone struct {
	MobilePhone     string `json:"mobilePhone" schema:"mobilePhone"`
	SecUserSignUpId string `json:"secUserSignUpId" schema:"secUserSignUpId"`
	ExpiredDatetime string `json:"expiredDatetime" schema:"expiredDatetime"`
	SmsRateData     string `json:"smsRateData" schema:"smsRateData"`
}

// TableName function
func (sbp SignUpByPhone) TableName() string {
	return "form_template"
}

// SecUsers struct
type SecUsers struct {
	SecUserId                string `json:"secUserId" schema:"sec_user_id"`
	HptHospitalID            string `json:"hptHospitalId" schema:"hpt_hospital_id"`
	Email                    string `json:"email" schema:"email"`
	EmailVerifStatus         bool   `json:"emailVerifStatus" schema:"email_verif_status"`
	MobilePhone              string `json:"mobilePhone" schema:"mobile_phone"`
	MobilePhoneVerifStatus   bool   `json:"mobilePhoneVerifStatus" schema:"mobile_phone_verif_status"`
	UserTypeCode             string `json:"userTypeCode" schema:"user_type_code"`
	Password                 string `json:"password" schema:"password"`
	Fullname                 string `json:"full_name" schema:"full_name"`
	GenderCode               string `json:"genderCode" schema:"gender_code"`
	LastSuccessLoginDatetime string `json:"lastSuccessLoginDatetime" schema:"last_success_login_datetime"`
	LoginFailCount           int    `json:"loginFailCount" schema:"login_fail_count"`
	MustChangePassword       bool   `json:"mustChangePassword" schema:"must_change_password"`
	AccountExpired           bool   `json:"accountExpired" schema:"account_expired"`
	CredentialsExpired       bool   `json:"credentialsExpired" schema:"credentials_expired"`
	IsActive                 bool   `json:"isActive" schema:"is_active"`
	IsDeleted                bool   `json:"isDeleted" schema:"is_deleted"`
	CreatedBy                string `json:"createdBy" schema:"created_by"`
	CreatedOn                string `json:"createdOn" schema:"created_on"`
	UpdatedBy                string `json:"updatedBy" schema:"updated_by"`
	UpdatedOn                string `json:"updatedOn" schema:"updated_by"`
	TimeZoneID               string `json:"timeZoneId" schema:"time_zone_id"`
}

// TableName function
func (sbu SecUsers) TableName() string {
	return "security.sec_users"
}

// SecUsers struct
type SecPatientSignInOtp struct {
	SecPatientSignInOtpID string    `json:"secPatientSignInOtpId" schema:"sec_patient_sign_in_otp_id"`
	SecUserID             string    `json:"secUserId" schema:"sec_user_id"`
	MobilePhone           string    `json:"mobilePhone" schema:"mobile_phone"`
	DeviceTypeCode        string    `json:"deviceTypeCode" schema:"device_type_code"`
	DeviceCode            string    `json:"deviceCode" schema:"device_code"`
	Otp                   string    `json:"otp" schema:"otp"`
	ExpiredDatetime       time.Time `json:"expiredDatetime" schema:"expired_datetime"`
	RetryCounter          int       `json:"retryCounter" schema:"retry_counter"`
	IsActive              bool      `json:"isActive" schema:"is_active"`
	IsDeleted             bool      `json:"isDeleted" schema:"is_deleted"`
	CreatedBy             string    `json:"createdBy" schema:"created_by"`
	CreatedOn             time.Time `json:"createdOn" schema:"created_on"`
	UpdatedBy             string    `json:"updatedBy" schema:"updated_by"`
	UpdatedOn             time.Time `json:"updatedOn" schema:"updated_on"`
}

// TableName function
func (sps SecPatientSignInOtp) TableName() string {
	return "security.sec_patient_sign_in_otps"
}

// SecUsers struct
type SmsLog struct {
	SmsLogID     string    `gorm:"column:sms_log_id" json:"smsLogId"`
	SmsReffID    string    `gorm:"column:sms_reff_id"  json:"smsReffId"`
	SmsTypeCode  string    `gorm:"column:sms_type_code" json:"smsTypeCode"`
	MobilePhone  string    `gorm:"column:mobile_phone" json:"mobilePhone"`
	SmsContent   string    `gorm:"column:sms_content" json:"smsContent"`
	SendingCount int       `gorm:"column:sending_count" json:"sendingCount"`
	SmsStatus    bool      `gorm:"column:sms_status" json:"smsStatus"`
	CreatedBy    string    `gorm:"column:created_by" json:"createdBy"`
	CreatedOn    time.Time `gorm:"column:created_on" json:"createdOn"`
	UpdatedBy    string    `gorm:"column:updated_by" json:"updatedBy" `
	UpdatedOn    time.Time `gorm:"column:updated_on" json:"updatedOn" `
}

// TableName function
func (sl SmsLog) TableName() string {
	return "sms.sms_logs"
}
