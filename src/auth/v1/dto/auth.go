package dto

import "time"

type LoginByPhoneRequest struct {
	DeviceType  string `json:"deviceType"`
	DeviceCode  string `json:"deviceCode"`
	MobilePhone string `json:"mobilePhone"`
	Password    string `json:"password"`
}

type LoginByPhoneResponse struct {
	SecPatientSignInOtpId string  `json:"secPatientSignInOtpId"`
	ExpiredDatetime       string  `json:"expiredDatetime"`
	SmsRateData           SmsRate `json:"smsRateData"`
}

type SmsRate struct {
	MobilePhone              string `json:"mobilePhone"`
	AllowSendSms             string `json:"allowSendSms"`
	ErrorMessage             string `json:"errorMessage"`
	IntervalBetweenSmsSecond string `json:"intervalBetweenSmsSecond"`
	IntervalGroupSecond      string `json:"intervalGroupSecond"`
	SmsCount                 string `json:"smsCount"`
	LastSmsDateTime          string `json:"lastSmsDateTime"`
	AllowSendInSecond        string `json:"allowSendInSecond"`
	AllowSendLocalDateTime   string `json:"allowSendLocalDateTime"`
}

type LoginByPhoneOtpRequest struct {
	SecPatientSignInOtpId string `json:"secPatientSignInOtpId"`
	Otp                   string `json:"otp"`
}

type LoginByPhoneOtpResponse struct {
	SecUserId          string       `json:"secUserId"`
	HospitalId         string       `json:"hospitalId"`
	TokenId            string       `json:"tokenId"`
	FullName           string       `json:"fullName"`
	MustChangePassword bool         `json:"mustChangePassword"`
	JwtToken           JwtTokenData `json:"jwtToken"`
	DeviceType         string       `json:"deviceType"`
	DeviceCode         string       `json:"deviceCode"`
	UserType           string       `json:"userType"`
	Role               string       `json:"role"`
	Features           []Features   `json:"features"` // []
	Menus              []Menus      `json:"menus"`    // []
	TimeZoneId         string       `json:"timeZoneId"`
}

type Token struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}

type JwtTokenData struct {
	LongToken  Token `json:"longToken"`
	ShortToken Token `json:"shortToken"`
}

type Features struct {
	SecFeatureId string `json:"secFeatureId"`
	FeatureCode  string `json:"featureCode"`
	FeatureName  string `json:"featureName"`
	Create       bool   `json:"create"`
	Read         bool   `json:"read"`
	Update       bool   `json:"update"`
	Delete       bool   `json:"delete"`
}

type Menus struct {
	SecMenuId       string `json:"secMenuId"`
	MenuCode        string `json:"menuCode"`
	MenuName        string `json:"menuName"`
	MenuDescription string `json:"menuDescription"`
	Create          bool   `json:"create"`
	Read            bool   `json:"read"`
	Update          bool   `json:"update"`
	Delete          bool   `json:"delete"`
}

type SmsRateData struct {
	MobilePhone              string `json:"mobilePhone"`
	AllowSendSms             bool   `json:"allowSendSms"`
	ErrorMessage             string `json:"errorMessage"`
	IntervalBetweenSmsSecond int    `json:"intervalBetweenSmsSecond"`
	IntervalGroupSecond      int    `json:"intervalGroupSecond"`
	SmsCount                 int    `json:"smsCount"`
	LastSmsDateTime          string `json:"lastSmsDateTime"`
	AllowSendInSecond        int    `json:"allowSendInSecond"`
	AllowSendLocalDateTime   string `json:"allowSendLocalDateTime"`
}

type PatientLoginPasswordResp struct {
	SecPatientSignInOtpId string      `json:"secPatientSignInOtpId"`
	ExpiredDatetime       string      `json:"expiredDatetime"`
	SmsRateData           SmsRateData `json:"smsRateData"`
}
