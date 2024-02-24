package routes

import (
	"dating-app/service-user/injection"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init(routes *echo.Echo, db *gorm.DB, redis *redis.Client) *echo.Echo {

	// SignUp Route & Injection
	user := injection.UserInjection(db, redis)
	UserRoute(routes, user)

	return routes
}
