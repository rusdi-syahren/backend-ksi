package external

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
	PartnerId         string `json:"partnerId" url:"partnerId"`
	PartnerName       string `json:"partnerName" url:"partnerName"`
	Password          string `json:"password" url:"password"`
	DestinationNumber string `json:"destinationNumber" url:"destinationNumber"`
	Message           string `json:"message" url:"message"`
}

type AcsSmsResponse struct {
	SmsBc AcsSmsRes `json:"smsbc" xml:"smsbc"`
}

type AcsSmsRes struct {
	Response AcsSmsReqPayload `json:"response" xml:"response"`
}

type AcsSmsResPayload struct {
	Datetime          string `json:"datetime" xml:"smsType"`
	Rrn               string `json:"rrn" xml:"rrn"`
	PartnerId         string `json:"partnerId" url:"partnerId"`
	PartnerName       string `json:"partnerName" url:"partnerName"`
	DestinationNumber string `json:"destinationNumber" url:"destinationNumber"`
	Message           string `json:"message" url:"message"`
	Rc                string `json:"rc" url:"rc"`
	Rm                string `json:"rm" url:"rm"`
}
