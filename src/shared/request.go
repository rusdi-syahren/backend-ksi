package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	log "github.com/sirupsen/logrus"
	"gopkg.in/eapache/go-resiliency.v1/retrier"
)

/*
	Request

	A tiny wrapper Go's Http Client with Circuit Breaker and Retry

	Configure your command before HTTP call

	hystrix.ConfigureCommand("POST-REVIEWS", hystrix.CommandConfig{
		Timeout:               2000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
	})

*/

/*
	Todo :
	- Add Hystrix Command Configuration
	- Check every http error code on retry function
*/
type (
	// Request struct
	Request struct {
		// retries count, but we can use this retries as circuit breaker's failure threshold
		retries int

		// sleepBetweenRetry
		// example time.Millisecond * 500
		sleepBetweenRetry time.Duration
		client            *http.Client
	}

	// HTTPRequest interface
	HTTPRequest interface {
		Do(breakerName, method, url string, body interface{}, headers map[string]string) ([]byte, error)
	}
)

// NewRequest function
// Request's Constructor
// Returns : *Request
func NewRequest(retries int, sleepBetweenRetry time.Duration) *Request {
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}

	client := &http.Client{}
	client.Transport = transport
	return &Request{
		retries:           retries,
		sleepBetweenRetry: sleepBetweenRetry,
		client:            client,
	}
}

// Do function, for http client call
func (request *Request) Do(breakerName, method, url string, body interface{}, headers map[string]string) ([]byte, error) {
	ctx := "Request-Do"
	output := make(chan []byte, 1)    // Declare the channel where the hystrix goroutine will put success responses.
	errors := hystrix.Go(breakerName, // Pass the name of the circuit breaker as first parameter.
		// 2nd parameter, the inlined func to run inside the breaker.
		func() error {
			// For hystrix, forward the err from the retrier. It's nil if successful.
			return request.retry(output, method, url, body, headers)
		},
		// 3rd parameter, the fallback func. In this case, we just do a bit of logging and return the error.
		func(err error) error {
			circuit, _, _ := hystrix.GetCircuit(breakerName)
			Log(log.ErrorLevel, err.Error(), ctx, "fallback_hystrix_error")

			Log(log.ErrorLevel, fmt.Sprintf("Circuit state is Open = %v", circuit.IsOpen()), ctx, "fallback_hystrix_circuit_is_open")
			return err
		})
	// Response and error handling. If the call was successful, the output channel gets the response. Otherwise,
	// the errors channel gives us the error.
	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return nil, err
	}
}

func (request *Request) retry(output chan []byte, method, url string, body interface{}, headers map[string]string) error {
	ctx := "Request-retry"

	// Create a retrier with constant backoff, retries number of attempts  with a sleepBetweenRetry sleep between retries.
	r := retrier.New(retrier.ConstantBackoff(request.retries, request.sleepBetweenRetry), nil)
	// This counter is just for getting some logging for showcasing, remove in production code.
	attempt := 0
	// Retrier works similar to hystrix, we pass the actual work (doing the HTTP request) in a func.
	err := r.Run(func() error {
		attempt++
		// Do HTTP request and handle response. If successful, pass resp.Body over output channel,
		// otherwise, do a bit of error logging and return the err.
		// Create the request. Omitted err handling for brevity
		payload, _ := json.Marshal(body)
		buf := bytes.NewBuffer(payload)
		req, _ := http.NewRequest(method, url, buf)

		// iterate optional data of headers
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		resp, err := request.client.Do(req)
		if err == nil && resp.StatusCode < 499 {
			responseBody, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				output <- responseBody
				return nil
			}
			return err
		} else if err == nil {
			err = fmt.Errorf("Status was %v", resp.StatusCode)
		}
		Log(log.ErrorLevel, fmt.Sprintf("Attempt : %d", attempt), ctx, "retrier")
		return err
	})
	return err
}
