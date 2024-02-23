package routes

import (
	"dating-app/service-user/handler"

	"github.com/labstack/echo/v4"
)

func SignUpRoute(routes *echo.Echo, handler *handler.SignUpHandler) {

	position := routes.Group("/sign-up")
	{
		position.POST("", handler.SignUp)
	}
}
