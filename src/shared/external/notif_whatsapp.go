package external

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Klinisia/backend-ksi/src/shared"
	"github.com/google/go-querystring/query"
)

// NotifWhatsapp struct
type NotifWhatsapp struct {
	whatsappBaseURL *url.URL
	basicAuth       string
	httpRequest     shared.HTTPRequest
}

// NewPaymentMidtrans function, PaymentMidtrans's constructor
func NewNotifWhatsapp() (*NotifWhatsapp, error) {

	//product detail env variable
	whatsappBaseURLStr, ok := os.LookupEnv("WHATSAPP_BASE_URL")
	if !ok {
		err := errors.New("you need to specify WHATSAPP_BASE_URL in the environment variable")
		return nil, err
	}

	whatsappBaseURL, err := url.Parse(whatsappBaseURLStr)
	if err != nil {
		err = errors.New("error URL parse")
		return nil, err
	}

	basicAuth, ok := os.LookupEnv("WHATSAPP_BASIC_AUTH")
	if !ok {
		err := errors.New("you need to specify WHATSAPP_BASIC_AUTH in the environment variable")
		return nil, err
	}

	httpRequest := shared.NewRequest(3, 500*time.Millisecond)

	return &NotifWhatsapp{
		whatsappBaseURL: whatsappBaseURL,
		basicAuth:       basicAuth,
		httpRequest:     httpRequest,
	}, nil
}

// SendMessage function
func (s *NotifWhatsapp) SendMessage(params WhatsappPayload) shared.Output {

	url := fmt.Sprintf("%s%s", s.whatsappBaseURL.String(), `/send-message`)
	v, _ := query.Values(params)

	payload := strings.NewReader(v.Encode())

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", s.basicAuth)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return shared.Output{Error: err}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return shared.Output{Error: err}
	}

	var waResponse ResponseWhatsApp

	err = json.Unmarshal(body, &waResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	fmt.Println(waResponse)
	fmt.Println(v.Encode())

	return shared.Output{Result: waResponse}
}

// SendMessage function
func (s *NotifWhatsapp) CheckNumber(params WhatsappPayloadCheckNum) shared.Output {

	url := fmt.Sprintf("%s%s", s.whatsappBaseURL.String(), `/check-phonenumber`)
	v, _ := query.Values(params)

	payload := strings.NewReader(v.Encode())

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", s.basicAuth)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return shared.Output{Error: err}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return shared.Output{Error: err}
	}

	var waResponse ResponseWhatsApp

	err = json.Unmarshal(body, &waResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	fmt.Println(waResponse)
	fmt.Println(v.Encode())

	return shared.Output{Result: waResponse}
}
