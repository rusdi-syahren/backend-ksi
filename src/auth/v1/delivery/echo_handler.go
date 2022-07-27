package delivery

import (
	"net/http"

	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/auth/v1/usecase"
	"github.com/Klinisia/backend-ksi/src/shared"

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
	group.POST("/sign-up/by-mobile-phone", h.AddAuth)
	group.POST("/login/patient/mobile-phone-password", h.LoginByPhonePassword)
	group.POST("/login/patient/mobile-phone-otp", h.LoginByPhoneOtp)

}

// SignUp by Phone
func (h *EchoHandler) AddAuth(c echo.Context) error {
	//initialize json schema template pointer
	response := new(shared.JSONSchemaTemplate)

	params := new(dto.SignUpByPhoneRequest)

	err := c.Bind(params)

	if err != nil {
		response.Success = false
		response.Message = "SignUp by Phone is failed"
		response.Code = http.StatusBadRequest
		response.SetData(shared.Empty{})

		return response.ShowHTTPResponse(c)
	}

	SignUp := h.AuthUsecase.SignUpByPhone(params)

	if SignUp.Error != nil {
		response.Success = false
		response.Message = SignUp.Error.Error()
		response.Code = http.StatusBadRequest

		return response.ShowHTTPResponse(c)
	}

	Auth := SignUp.Result.(domain.SignUpByPhone)

	response.Success = true
	response.Message = "SignUp By Phone"
	response.Code = http.StatusOK
	response.SetData(Auth)

	return response.ShowHTTPResponse(c)
}

func (h *EchoHandler) LoginByPhonePassword(c echo.Context) error {
	//initialize json schema template pointer
	response := new(shared.JSONResponse)

	params := new(dto.LoginByPhoneRequest)

	err := c.Bind(params)

	if err != nil {
		response.Error = err.Error()
		response.Message = "Login by Phone is failed"
		response.Status = http.StatusBadRequest
		response.SetPayload(shared.Empty{})

		return response.ShowHTTPResponse(c)
	}

	signIn := h.AuthUsecase.LoginByPhonePassword(params)

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

func (h *EchoHandler) LoginByPhoneOtp(c echo.Context) error {
	//initialize json schema template pointer
	response := new(shared.JSONSchemaTemplate)

	params := new(dto.LoginByPhoneOtpRequest)

	err := c.Bind(params)

	if err != nil {
		response.Success = false
		response.Message = "Login by otp is failed"
		response.Code = http.StatusBadRequest
		response.SetData(shared.Empty{})

		return response.ShowHTTPResponse(c)
	}

	SignUp := h.AuthUsecase.LoginByPhoneOtp(params)

	if SignUp.Error != nil {
		response.Success = false
		response.Message = SignUp.Error.Error()
		response.Code = http.StatusBadRequest

		return response.ShowHTTPResponse(c)
	}

	Auth := SignUp.Result.(domain.SignUpByPhone)

	response.Success = true
	response.Message = "Login By Phone"
	response.Code = http.StatusOK
	response.SetData(Auth)

	return response.ShowHTTPResponse(c)
}
