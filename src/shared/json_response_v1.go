package shared

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
)

// JSONResponse data structure
type JSONResponse struct {
	Timestamp string          `json:"timestamp"`
	Path      string          `json:"path"`
	Status    int             `json:"status"`
	Error     string          `json:"error"`
	Message   string          `json:"message"`
	RequestId string          `json:"requestId"`
	TraceId   string          `json:"traceId"`
	ProcMs    int             `json:"procMs"`
	Errors    []ErrorResponse `json:"errors"`
	Payload   interface{}     `json:"payload"`
}

type ErrorResponse struct {
	Field   string `json:"field"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SetData function
func (t *JSONResponse) SetPayload(payload interface{}) {
	t.Payload = payload
	t.Timestamp = time.Now().Local().Format("2006-01-02T15:04:05.99999")
}

// ShowHTTPResponse function
func (t *JSONResponse) ShowHTTPResponse(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(t.Status)
	return json.NewEncoder(c.Response()).Encode(t)
}
