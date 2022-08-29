package delivery

import (
	"net/http"

	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/dto"
	"github.com/rusdi-syahren/backend-ksi/src/auth/v1/usecase"
	"github.com/rusdi-syahren/backend-ksi/src/shared"

	"github.com/labstack/echo/v4"
)

// EchoHandler structure
type EchoHandler struct {
	AuthUsecase usecase.AuthUsecase
}

// NewEchoHandler function
// Returns *EchoHandler
func NewEchoHandler(AuthUsecase usecase.AuthUsecase) *EchoHandler {
	return &EchoHandler{AuthUsecase: AuthUsecase}
}

// Mount function
// Params : *echo.Group
func (h *EchoHandler) Mount(group *echo.Group) {
	group.POST("/login/patient/mobile-phone-password", h.SignInByPhonePassword)
	group.POST("/login/patient/mobile-phone-otp", h.SignInByPhoneOtp)

}

func (h *EchoHandler) SignInByPhonePassword(c echo.Context) error {
	//initialize json schema template pointer
	response := new(shared.JSONResponse)

	params := new(dto.LoginByPhoneRequest)

	err := c.Bind(params)

	if err != nil {
		response.Error = err.Error()
		response.Message = "SignIn by Phone is failed"
		response.Status = http.StatusBadRequest
		response.SetPayload(shared.Empty{})

		return response.ShowHTTPResponse(c)
	}

	signIn := h.AuthUsecase.SignInByPhonePassword(params)

	if signIn.Error != nil {
		errorResponse := signIn.Errors.(shared.ErrorResponse)
		var errorList []shared.ErrorResponse
		errorList = append(errorList, errorResponse)
		response.Error = signIn.Error.Error()
		response.Message = ""
		response.Status = signIn.Code
		response.Errors = errorList
		response.SetPayload(signIn.Result)

		return response.ShowHTTPResponse(c)
	}

	Auth := signIn.Result.(dto.PatientLoginPasswordResp)

	response.Status = http.StatusOK

	response.SetPayload(Auth)

	return response.ShowHTTPResponse(c)
}

func (h *EchoHandler) SignInByPhoneOtp(c echo.Context) error {
	//initialize json schema template pointer
	response := new(shared.JSONResponse)

	params := new(dto.LoginByPhoneOtpRequest)

	err := c.Bind(params)

	if err != nil {
		response.Error = err.Error()
		response.Message = "SignIn by OTP is failed"
		response.Status = http.StatusBadRequest
		response.SetPayload(shared.Empty{})

		return response.ShowHTTPResponse(c)
	}

	signIn := h.AuthUsecase.SignInByPhoneOtp(params)

	if signIn.Error != nil {
		errorResponse := signIn.Errors.(shared.ErrorResponse)
		var errorList []shared.ErrorResponse
		errorList = append(errorList, errorResponse)
		response.Error = signIn.Error.Error()
		response.Message = ""
		response.Status = signIn.Code
		response.Errors = errorList
		response.SetPayload(signIn.Result)

		return response.ShowHTTPResponse(c)
	}

	Auth := signIn.Result.(dto.LoginByPhoneOtpResponse)

	response.Status = http.StatusOK

	response.SetPayload(Auth)

	return response.ShowHTTPResponse(c)
}
