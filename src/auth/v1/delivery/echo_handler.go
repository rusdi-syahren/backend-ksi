package delivery

import (
	"net/http"

	"github.com/Klinisia/backend-ksi/src/auth/v1/domain"
	"github.com/Klinisia/backend-ksi/src/auth/v1/dto"
	"github.com/Klinisia/backend-ksi/src/auth/v1/usecase"
	"github.com/Klinisia/backend-ksi/src/shared"

	"github.com/labstack/echo"
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
