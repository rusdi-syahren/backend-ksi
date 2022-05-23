package external

// Payload model
type WhatsappPayload struct {
	Number  string `json:"number" url:"number"`
	Message string `json:"message" url:"message"`
}

type WhatsappPayloadCheckNum struct {
	Number string `json:"number" url:"number"`
}

type ResponseWhatsApp struct {
	Status   bool        `json:"status"`
	Response interface{} `json:"response"`
	Message  string      `json:"message"`
}
