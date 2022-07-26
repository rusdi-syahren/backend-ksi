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
)

// SmsAcs struct
type SmsAcs struct {
	acsBaseURL  *url.URL
	httpRequest shared.HTTPRequest
}

func NewSmsAcs() (*SmsAcs, error) {

	acsBaseURLStr, ok := os.LookupEnv("ACS_URL")
	if !ok {
		err := errors.New("you need to specify ACS_URL in the environment variable")
		return nil, err
	}

	acsBaseURL, err := url.Parse(acsBaseURLStr)
	if err != nil {
		err = errors.New("error URL parse")
		return nil, err
	}

	httpRequest := shared.NewRequest(3, 500*time.Millisecond)

	return &SmsAcs{
		acsBaseURL:  acsBaseURL,
		httpRequest: httpRequest,
	}, nil
}

// SendSms function
func (s *SmsAcs) SendSms(params string) shared.Output {

	url := fmt.Sprintf("%s", s.acsBaseURL.String())
	// v, _ := query.Values(params)

	payload := strings.NewReader(params)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/xml")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return shared.Output{Error: err}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return shared.Output{Error: err}
	}

	var smsResponse AcsSmsResponse

	err = json.Unmarshal(body, &smsResponse)

	if err != nil {
		return shared.Output{Error: err}
	}

	fmt.Println(smsResponse)

	return shared.Output{Result: smsResponse}
}
