package external

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/rusdi-syahren/backend-ksi/src/shared"
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
func (s *SmsAcs) SendSms(params shared.AcsSmsRequest, isSimulation bool) shared.Output {

	var body string
	if isSimulation {
		body = shared.CreateSmsXmlResponse(&params)

	} else {
		smsPayload := shared.CreateSmsXmlRequest(&params)
		url := fmt.Sprintf("%s", s.acsBaseURL.String())
		payload := strings.NewReader(smsPayload)

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/xml")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			return shared.Output{Error: err}
		}

		defer res.Body.Close()
		rs, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return shared.Output{Error: err}
		}
		body = string(rs)
	}

	var output AcsSmsResponse

	xml.Unmarshal([]byte(body), &output)

	var allResponse AcsSmsAllResponse
	allResponse.Response = output
	allResponse.ResponseStr = body

	// spew.Dump(allResponse)

	return shared.Output{Result: &allResponse}
}
