package external

import "encoding/xml"

type SmsRequest struct {
	DestinationNumber string `json:"destinationNumber" url:"destinationNumber"`
	Message           string `json:"message" url:"message"`
	SmsReffId         string `json:"smsReffId" url:"smsReffId"`
	SmsType           string `json:"smsType" url:"smsType"`
}

type SmsResponse struct {
	Status bool   `json:"status"`
	Rc     string `json:"rc"`
	Rm     string `json:"rm"`
}

type AcsSmsRequest struct {
	SmsBc AcsSmsReq `json:"smsbc" xml:"smsbc"`
}

type AcsSmsReq struct {
	Request AcsSmsReqPayload `json:"request" xml:"request"`
}

type AcsSmsReqPayload struct {
	Datetime          string `json:"datetime" xml:"smsType"`
	Rrn               string `json:"rrn" xml:"rrn"`
	PartnerId         string `json:"partnerId" xml:"partnerId"`
	PartnerName       string `json:"partnerName" xml:"partnerName"`
	Password          string `json:"password" xml:"password"`
	DestinationNumber string `json:"destinationNumber" xml:"destinationNumber"`
	Message           string `json:"message" xml:"message"`
}

type AcsSmsResponse struct {
	XMLName  xml.Name  `xml:"smsbc" json:"-"`
	Response AcsSmsRes `json:"response" xml:"response"`
}

type AcsSmsRes struct {
	XMLName           xml.Name `xml:"response" json:"-"`
	Datetime          string   `xml:"datetime" json:"datetime"`
	Rrn               string   `xml:"rrn" json:"rrn"`
	PartnerId         string   `xml:"partnerId" json:"partnerId"`
	PartnerName       string   `xml:"partnerName" json:"partnerName"`
	DestinationNumber string   `xml:"destinationNumber" json:"destinationNumber"`
	Message           string   `xml:"message" json:"message"`
	Rc                string   `xml:"rc" json:"rc"`
	Rm                string   `xml:"rm" json:"rm"`
}

type AcsSmsAllResponse struct {
	Response    AcsSmsResponse `json:"response"`
	ResponseStr string         `json:"responseStr"`
}

const (
	// date format
	DateFormat    string = "2006-01-02T15:04:05.99999"
	SmsDateFormat string = "0102150405"
	DateFormatStr string = "2006-01-02 15:04:05"

	// sms content
	SmsOtpSignIn            string = "OTP Login pengguna = %s. Jangan share OTP ke siapapun"
	SmsValidationLimit      string = "SMS terakhir untuk no hp %s, jam %s, minimal jarak %s detik"
	SmsValidationLimitGroup string = "Maaf, nomor %s dapat mengirim sms lagi pada %s"
)
