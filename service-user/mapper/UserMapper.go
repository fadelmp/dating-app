package mapper

import (
	"dating-app/service-user/domain"
	sharedBase "dating-app/shared/domain"
)

// Interface
type UserMapper interface {
	ToTempUser(string, string, string, string, sharedBase.Base) domain.TempUser
}

// Class
type UserMapperImpl struct {
}

// Constructor
func NewUserMapper() *UserMapperImpl {
	return &UserMapperImpl{}
}

// Implementation

func (m *UserMapperImpl) ToTempUser(
	id string,
	email string,
	password string,
	otpCode string,
	base sharedBase.Base) domain.TempUser {

	return domain.TempUser{
		Id:         id,
		Email:      email,
		Password:   password,
		OtpCode:    otpCode,
		TryAttempt: 3,
		Base:       base,
	}
}
