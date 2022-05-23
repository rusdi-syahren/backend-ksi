package external

// TokenRequest struct
type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	UserID       string `json:"user_id"`
	Type         int    `json:"type"`
	RoomID       string `json:"room_id"`
}

// TokenResp struct
type TokenResp struct {
	Token string `json:"token"`
}

// Partisipant struct
type Partisipant struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

// Partisipants list
type Partisipants []Partisipant

// RoomRequest struct
type RoomRequest struct {
	RoomID       string       `json:"room_id"`
	Partisipants Partisipants `json:"participants"`
	Message      string       `json:"message"`
	FromUsername string       `json:"from_username"`
	FromType     int          `json:"from_type"`
	UserID       string       `json:"user_id"`
}

// RoomResponse struct
type RoomResponse struct {
	Code         int          `json:"code"`
	DataResponse DataResponse `json:"data"`
	Message      string       `json:"message"`
}

// DataResponse struct
type DataResponse struct {
	ID           string `json:"id"`
	ClientRoomID string `json:"client_room_id"`
	Status       bool   `json:"status"`
}

// MessageRequest struct
type MessageRequest struct {
	RoomID       string `json:"room_id"`
	Message      string `json:"message"`
	FromUsername string `json:"from_username"`
	FromType     int    `json:"from_type"`
	Media        string `json:"media"`
}
