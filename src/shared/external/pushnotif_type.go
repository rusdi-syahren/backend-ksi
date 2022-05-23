package external

// Payload model
type Payload struct {
	To               string                 `json:"to"`
	Data             map[string]interface{} `json:"data"`
	Priority         string                 `json:"priority"`
	ContentAvailable bool                   `json:"content_available"`

	// Notification *Notification          `json:"notification"`
	// Android *Android               `json:"android"`
	// Apns    *Apns                  `json:"apns"`
}

// Notification model
type Notification struct {
	Title          string `json:"title"`
	Body           string `json:"body"`
	Image          string `json:"image"`
	Sound          string `json:"sound"`
	MutableContent bool   `json:"mutable-content"`
	ResourceID     string `json:"resourceId"`
	ResourceName   string `json:"resourceName"`
}

// Apns struct
type Apns struct {
	ApnsPayload *ApnsPayload `json:"payload"`
}

// ApnsPayload struct
type ApnsPayload struct {
	Aps *Aps `json:"aps"`
}

// Aps struct
type Aps struct {
	Sound string `json:"sound"`
}

// Android struct
type Android struct {
	AndroidNotification *AndroidNotification `json:"notification"`
}

// AndroidNotification struct
type AndroidNotification struct {
	Sound string `json:"sound"`
}
