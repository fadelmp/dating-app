package handler

import (
	"dating-app/service-user/dto"
	"dating-app/service-user/message"
	"dating-app/service-user/usecase"
	handler "dating-app/shared/handler"

	"github.com/labstack/echo/v4"
)

// Interface
type UserHandler interface {
	SignUp(e echo.Context) error
}

// Class
type UserHandlerImpl struct {
	usecase usecase.UserUsecase
}

// Constructor
func NewUserHandler(usecase usecase.UserUsecase) *UserHandlerImpl {
	return &UserHandlerImpl{
		usecase: usecase,
	}
}

func (h *UserHandlerImpl) SignUp(e echo.Context) error {

	var signUp dto.SignUp

	if e.Bind(&signUp) != nil {
		return handler.BadRequest(e)
	}

	if err := h.usecase.SignUp(signUp); err != nil {
		return handler.Error(e, message.TitleCreateFailed)
	}

	return handler.Success(e, "")
}
