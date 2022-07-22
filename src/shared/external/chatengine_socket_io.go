package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/labstack/echo"
)

// ChatEngineSocketIO struct
type ChatEngineSocketIO struct {
	clientID     string
	clientSecret string
	chatURL      string
	httpRequest  shared.HTTPRequest
}

// NewChatEngineSocketIO function, NewChatEngineSocketIO's constructor
func NewChatEngineSocketIO() (*ChatEngineSocketIO, error) {

	clientID, ok := os.LookupEnv("CLIENT_ID_CHAT")
	if !ok {
		err := errors.New("you need to specify CLIENT_ID_CHAT in the environment variable")
		return nil, err
	}

	clientSecret, ok := os.LookupEnv("CLIENT_SECRET_CHAT")
	if !ok {
		err := errors.New("you need to specify CLIENT_ID_CHAT in the environment variable")
		return nil, err
	}

	chatURL, ok := os.LookupEnv("URL_CHAT")
	if !ok {
		err := errors.New("you need to specify CLIENT_ID_CHAT in the environment variable")
		return nil, err
	}
	httpRequest := shared.NewRequest(3, 5000*time.Millisecond)

	return &ChatEngineSocketIO{
		clientID:     clientID,
		clientSecret: clientSecret,
		chatURL:      chatURL,
		httpRequest:  httpRequest,
	}, nil
}

// GetToken function
func (ct *ChatEngineSocketIO) GetToken(c echo.Context, body *TokenRequest) shared.Output {

	// Configure Hystrix
	hystrix.ConfigureCommand("Get-Payment-Midtrans", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	var tokenResponse TokenResp

	urlChat := fmt.Sprintf("%s%s", ct.chatURL, "/token")

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	resp, err := ct.httpRequest.Do("Get-Token-Chat", "POST", urlChat, body, headers)

	if err != nil {
		return shared.Output{Error: err}
	}

	err = json.Unmarshal(resp, &tokenResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	return shared.Output{Result: &tokenResponse}
}

// CreateRoom function
func (ct *ChatEngineSocketIO) CreateRoom(c echo.Context, body *RoomRequest) shared.Output {

	// Configure Hystrix
	hystrix.ConfigureCommand("Create-Room", hystrix.CommandConfig{
		Timeout:               100000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	var RoomResponse RoomResponse

	urlChat := fmt.Sprintf("%s%s", ct.chatURL, "/room")

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	resp, err := ct.httpRequest.Do("Create-Room", "POST", urlChat, body, headers)

	if err != nil {
		fmt.Printf("%+v\n", body)
		return shared.Output{Error: err}
	}

	err = json.Unmarshal(resp, &RoomResponse)
	fmt.Printf("%v", RoomResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	if RoomResponse.Code != 201 && RoomResponse.Code != 409 {
		return shared.Output{Error: errors.New(RoomResponse.Message)}
	}

	return shared.Output{Result: &RoomResponse}
}

// SendMessage function
func (ct *ChatEngineSocketIO) SendMessage(body *MessageRequest) shared.Output {

	// Configure Hystrix
	hystrix.ConfigureCommand("Send-Message", hystrix.CommandConfig{
		Timeout:               100000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	var RoomResponse RoomResponse

	urlChat := fmt.Sprintf("%s%s", ct.chatURL, "/message")

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	resp, err := ct.httpRequest.Do("Send-Message", "POST", urlChat, body, headers)

	if err != nil {
		fmt.Printf("%+v\n", body)
		return shared.Output{Error: err}
	}

	err = json.Unmarshal(resp, &RoomResponse)
	fmt.Printf("%v", RoomResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	if RoomResponse.Code != 200 && RoomResponse.Code != 201 {
		return shared.Output{Error: errors.New(RoomResponse.Message)}
	}

	return shared.Output{Result: &RoomResponse}
}

// GetListRoom function
func (ct *ChatEngineSocketIO) GetListRoom(c echo.Context, body *RoomRequest) shared.Output {

	var roomResponse RoomResponse

	return shared.Output{Result: &roomResponse}
}
