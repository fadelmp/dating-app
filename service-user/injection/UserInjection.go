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

	userRepo := repository.NewUserRepository(db, redis)
	tempRepo := repository.NewTempUserRepository(db, redis)

	mapper := mapper.NewUserMapper()
	comparator := comparator.NewUserComparator(userRepo, tempRepo)

	signUpUsecase := usecase.NewSignUpUsecase(tempRepo, mapper, comparator)
	signInUsecase := usecase.NewSignInUsecase(userRepo, mapper, comparator)
	verifyUsecase := usecase.NewVerifyUsecase(tempRepo, userRepo, mapper, comparator)

	handler := handler.NewUserHandler(signUpUsecase, signInUsecase, verifyUsecase)

	return handler
}
