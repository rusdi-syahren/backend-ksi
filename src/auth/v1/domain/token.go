package domain

import "time"

type SecLongToken struct {
	SecLongTokenId  string    `gorm:"column:SEC_LONG_TOKEN_ID" json:"secLongTokenId"`
	SecUserId       string    `gorm:"column:SEC_USER_ID"  json:"secUserId"`
	HptHospitalId   string    `gorm:"column:HPT_HOSPITAL_ID" json:"hptHospitalId"`
	UserTypeCode    string    `gorm:"column:USER_TYPE_CODE" json:"userTypeCode"`
	DeviceType      string    `gorm:"column:DEVICE_TYPE" json:"deviceType"`
	DeviceCode      string    `gorm:"column:DEVICE_CODE" json:"deviceCode"`
	LongToken       string    `gorm:"column:LONG_TOKEN" json:"longToken"`
	ExpiredDatetime time.Time `gorm:"column:EXPIRED_DATETIME" json:"expiredDatetime"`
}

// TableName function
func (slt SecLongToken) TableName() string {
	return "SECURITY.SEC_LONG_TOKENS"
}
