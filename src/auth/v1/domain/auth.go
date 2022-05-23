package domain

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
