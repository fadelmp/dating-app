package mapper

import (
	"dating-app/service-user/domain"
	"dating-app/service-user/dto"
	sharedBase "dating-app/shared/domain"
)

// Interface
type UserMapper interface {
	ToTempUser(dto.SignUp, sharedBase.Base) domain.TempUser
}

// Class
type UserMapperImpl struct {
}

// Constructor
func NewUserMapper() *UserMapperImpl {
	return &UserMapperImpl{}
}

// Implementation

func (m *UserMapperImpl) ToTempUser(signUp dto.SignUp, base sharedBase.Base) domain.TempUser {

	return domain.TempUser{
		Email:      signUp.Email,
		Password:   signUp.Password,
		TryAttempt: 3,
		Base:       base,
	}
}
