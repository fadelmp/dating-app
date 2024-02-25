package handler

import (
	"dating-app/service-user/dto"
	"dating-app/service-user/message"
	"dating-app/service-user/usecase"
	handler "dating-app/shared/handler"
	"fmt"

	"github.com/labstack/echo/v4"
)

// Interface
type UserHandler interface {
	SignUp(e echo.Context) error
}

// Class
type UserHandlerImpl struct {
	usecase usecase.SignUpUsecase
}

// Constructor
func NewUserHandler(usecase usecase.SignUpUsecase) *UserHandlerImpl {
	return &UserHandlerImpl{
		usecase: usecase,
	}
}

func (h *UserHandlerImpl) SignUp(e echo.Context) error {

	var signUp dto.SignUp

	if e.Bind(&signUp) != nil {
		return handler.BadRequest(e)
	}

	// validate := validator.New()
	// err := validate.Struct(signUp)
	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		fmt.Println("Field:", err.Field())
	// 		fmt.Println("Error:", err.ActualTag())
	// 		fmt.Println("Message:", err.Param())
	// 		return handler.Error(e, err.Field()+" "+err.ActualTag()+" "+err.Param())
	// 	}
	// }

	if err := h.usecase.SignUp(signUp); err != nil {
		fmt.Println(err)
		return handler.Error(e, err.Error())
	}

	return handler.Success(e, message.SignUpSuccess, "")
}
