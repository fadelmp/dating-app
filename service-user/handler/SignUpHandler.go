package handler

import (
	"dating-app/service-user/dto"
	"dating-app/service-user/message"
	"dating-app/service-user/usecase"
	handler "dating-app/shared/handler"

	"github.com/labstack/echo/v4"
)

// Interface
type SignUpHandlerContract interface {
	SignUp(e echo.Context) error
}

// Class
type SignUpHandler struct {
	usecase usecase.SignUpUsecaseContract
}

// Constructor
func NewSignUpHandler(usecase usecase.SignUpUsecaseContract) *SignUpHandler {
	return &SignUpHandler{
		usecase: usecase,
	}
}

func (h *SignUpHandler) SignUp(e echo.Context) error {

	var signUp dto.SignUp

	if e.Bind(&signUp) != nil {
		return handler.BadRequest(e)
	}

	if err := h.usecase.SignUp(signUp); err != nil {
		return handler.Error(e, message.TitleCreateFailed)
	}

	return handler.Success(e, "")
}
