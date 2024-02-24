package routes

import (
	"dating-app/service-user/handler"

	"github.com/labstack/echo/v4"
)

func UserRoute(routes *echo.Echo, handler *handler.UserHandlerImpl) {

	user := routes.Group("/user")
	{
		user.POST("/sign-up", handler.SignUp)
	}
}
