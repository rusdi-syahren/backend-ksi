package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Klinisia/backend-ksi/src/shared"

	"github.com/labstack/echo/v4"
)

type response struct {
	Error bool        `json:"error"`
	Body  interface{} `json:"body"`
}

type FirebaseImpl struct {
	host        string
	key         string
	httpRequest *shared.Request
}

// NewFirebaseREST create new firebase rest
func NewFirebaseREST() (*FirebaseImpl, error) {

	urlBaseURL, ok := os.LookupEnv("FIREBASE_BASE_URL")
	if !ok {
		err := errors.New("you need to specify FIREBASE_BASE_URL in the environment variable")
		return nil, err
	}

	urlKey, ok := os.LookupEnv("FIREBASE_API_KEY")
	if !ok {
		err := errors.New("you need to specify FIREBASE_API_KEY in the environment variable")
		return nil, err
	}
	httpRequest := shared.NewRequest(3, 5000*time.Millisecond)

	return &FirebaseImpl{
		host:        urlBaseURL,
		key:         urlKey,
		httpRequest: httpRequest,
	}, nil
}

// SendNotification func
func (f *FirebaseImpl) SendNotification(ctx context.Context, payload *Payload) <-chan []byte {
	operationName := "Firebase-SendNotification"
	output := make(chan []byte)

	go func() {

		var responses []response

		header := map[string]string{
			echo.HeaderAuthorization: fmt.Sprintf("key=%s", f.key),
			echo.HeaderContentType:   echo.MIMEApplicationJSON,
		}

		body, err := f.httpRequest.Do(operationName, http.MethodPost, f.host, *payload, header)
		if err != nil {
			responses = append(responses, response{
				Error: true, Body: err.Error(),
			})
			fmt.Println("[firebase-error] " + err.Error())
		}

		var resp response
		json.Unmarshal(body, &resp.Body)
		responses = append(responses, resp)

		result, _ := json.Marshal(responses)
		output <- result
	}()

	return output
}
