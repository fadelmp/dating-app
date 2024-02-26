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
	SignIn(e echo.Context) error
	VerifyEmail(e echo.Context) error
}

// Class
type UserHandlerImpl struct {
	signUpUsecase usecase.SignUpUsecase
	signInUsecase usecase.SignInUsecase
	verifyUsecase usecase.VerifyUsecase
}

// Constructor
func NewUserHandler(
	signUpUsecase usecase.SignUpUsecase,
	signInUsecase usecase.SignInUsecase,
	verifyUsecase usecase.VerifyUsecase,
) *UserHandlerImpl {
	return &UserHandlerImpl{
		signUpUsecase: signUpUsecase,
		signInUsecase: signInUsecase,
		verifyUsecase: verifyUsecase,
	}
}

func (h *UserHandlerImpl) SignUp(e echo.Context) error {

	var sign dto.Sign

	if e.Bind(&sign) != nil {
		return handler.BadRequest(e)
	}

	result, err := h.signUpUsecase.SignUp(sign)
	if err != nil {
		return handler.Error(e, err.Error())
	}

	return handler.Success(e, message.SignUpSuccess, result)
}

func (h *UserHandlerImpl) SignIn(e echo.Context) error {

	var sign dto.Sign

	if e.Bind(&sign) != nil {
		return handler.BadRequest(e)
	}

	result, err := h.signInUsecase.SignIn(sign)
	if err != nil {
		return handler.Error(e, err.Error())
	}

	return handler.Success(e, message.SignInSuccess, result)
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
