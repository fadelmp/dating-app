package injection

import (
	"dating-app/service-user/comparator"
	"dating-app/service-user/handler"
	"dating-app/service-user/mapper"
	"dating-app/service-user/repository"
	"dating-app/service-user/usecase"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

func UserInjection(db *gorm.DB, redis *redis.Client) *handler.UserHandlerImpl {

	mapper := mapper.NewUserMapper()
	userRepo := repository.NewUserRepository(db, redis)
	tempRepo := repository.NewTempUserRepository(db, redis)

	comparator := comparator.NewUserComparator(userRepo, tempRepo)
	usecase := usecase.NewSignUpUsecase(tempRepo, mapper, comparator)
	handler := handler.NewUserHandler(usecase)

	return handler
}
