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
	VerifyEmail(e echo.Context) error
}

// Class
type UserHandlerImpl struct {
	signUpUsecase usecase.SignUpUsecase
	verifyUsecase usecase.VerifyUsecase
}

// Constructor
func NewUserHandler(
	signUpUsecase usecase.SignUpUsecase,
	verifyUsecase usecase.VerifyUsecase,
) *UserHandlerImpl {
	return &UserHandlerImpl{
		signUpUsecase: signUpUsecase,
		verifyUsecase: verifyUsecase,
	}
}

func (h *UserHandlerImpl) SignUp(e echo.Context) error {

	var signUp dto.SignUp

	if e.Bind(&signUp) != nil {
		return handler.BadRequest(e)
	}

	result, err := h.signUpUsecase.SignUp(signUp)
	if err != nil {
		return handler.Error(e, err.Error())
	}

	return handler.Success(e, message.SignUpSuccess, result)
}

func (h *UserHandlerImpl) VerifyEmail(e echo.Context) error {

	var verify dto.Verify

	if e.Bind(&verify) != nil {
		return handler.BadRequest(e)
	}

	if err := h.verifyUsecase.VerifyEmail(verify); err != nil {
		return handler.Error(e, err.Error())
	}

	return handler.Success(e, message.VerifySuccess, "")
}
