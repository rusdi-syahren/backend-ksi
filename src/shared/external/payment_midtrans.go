package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/labstack/echo/v4"
	"github.com/rusdi-syahren/backend-ksi/src/shared"
)

// PaymentMidtrans struct
type PaymentMidtrans struct {
	midtransBaseURL *url.URL
	apiKey          string
	clientKey       string
	httpRequest     shared.HTTPRequest
}

// NewPaymentMidtrans function, PaymentMidtrans's constructor
func NewPaymentMidtrans() (*PaymentMidtrans, error) {

	//product detail env variable
	midtransBaseURLStr, ok := os.LookupEnv("MIDTRANS_BASE_URL")
	if !ok {
		err := errors.New("you need to specify MIDTRANS_BASE_URL in the environment variable")
		return nil, err
	}

	midtransBaseURL, err := url.Parse(midtransBaseURLStr)
	if err != nil {
		err = errors.New("error URL parse")
		return nil, err
	}

	apiKey, ok := os.LookupEnv("MIDTRANS_APIKEY")
	if !ok {
		err := errors.New("you need to specify MIDTRANS_APIKEY in the environment variable")
		return nil, err
	}

	clientKey, ok := os.LookupEnv("MIDTRANS_CLIENT_KEY")
	if !ok {
		err := errors.New("you need to specify MIDTRANS_CLIENT_KEY in the environment variable")
		return nil, err
	}

	httpRequest := shared.NewRequest(3, 500*time.Millisecond)

	return &PaymentMidtrans{
		midtransBaseURL: midtransBaseURL,
		apiKey:          apiKey,
		clientKey:       clientKey,
		httpRequest:     httpRequest,
	}, nil
}

// Charge function
func (s *PaymentMidtrans) Charge(c echo.Context, body interface{}) shared.Output {
	// Configure Hystrix
	hystrix.ConfigureCommand("Get-Payment-Midtrans", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	var paymentResponse PaymentResponse

	urlMidtrans := fmt.Sprintf("%s%s", s.midtransBaseURL.String(), "/v2/charge")

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": s.apiKey,
	}

	resp, err := s.httpRequest.Do("Get-Payment-Midtrans", "POST", urlMidtrans, body, headers)

	if err != nil {
		return shared.Output{Error: err}
	}

	err = json.Unmarshal(resp, &paymentResponse)
	fmt.Printf("%v", paymentResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	if paymentResponse.StatusCode != "201" && paymentResponse.StatusCode != "200" {
		return shared.Output{Error: errors.New(paymentResponse.StatusMessage)}
	}

	return shared.Output{Result: &paymentResponse}
}

// CheckStatusPayment function
// func (s *PaymentMidtrans) CheckStatusPayment(c echo.Context, orderID string) shared.Output {
// 	// Configure Hystrix
// 	hystrix.ConfigureCommand("Get-Payment-Midtrans", hystrix.CommandConfig{
// 		Timeout:               5000,
// 		MaxConcurrentRequests: 100,
// 		ErrorPercentThreshold: 25,
// 	})

// 	var paymentResponse notifDomain.Notification

// 	urlMidtrans := fmt.Sprintf("%s%s", s.midtransBaseURL.String(), "/v2/"+orderID+"/status")

// 	headers := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Accept":        "application/json",
// 		"Authorization": s.apiKey,
// 	}

// 	resp, err := s.httpRequest.Do("Get-Payment-Midtrans", "GET", urlMidtrans, nil, headers)

// 	if err != nil {
// 		return shared.Output{Error: err}
// 	}

// 	err = json.Unmarshal(resp, &paymentResponse)

// 	if err != nil {
// 		return shared.Output{Error: err}
// 	}

// 	return shared.Output{Result: &paymentResponse}
// }

// GetToken function
func (s *PaymentMidtrans) GetToken(c echo.Context, body TokenParams) shared.Output {
	// Configure Hystrix
	hystrix.ConfigureCommand("Get-Payment-Midtrans", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

	var paymentResponse TokenResponse

	urlMidtrans := fmt.Sprintf("%s%s", s.midtransBaseURL.String(), `/v2/token`+
		`?gross_amount=`+body.GrossAmount+
		`&card_number=`+body.CardNumber+
		`&card_exp_month=`+body.CardExpMonth+
		`&card_exp_year=`+body.CardExpYear+
		"&card_cvv="+body.CardCVV+
		"&client_key="+s.clientKey)

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	resp, err := s.httpRequest.Do("Get-Token-cc-Midtrans", "GET", urlMidtrans, nil, headers)

	if err != nil {
		return shared.Output{Error: err}
	}
	err = json.Unmarshal(resp, &paymentResponse)
	// fmt.Printf("%v", paymentResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	return shared.Output{Result: paymentResponse}
}
