package dto

// SignUpByPhoneRequest struct
type SignUpByPhoneRequest struct {
	MobilePhone string `json:"mobilePhone"`
	Password    string `json:"pasword"`
	RePassword  string `json:"rePassword"`
	UserType    string `json:"userType"`
}
