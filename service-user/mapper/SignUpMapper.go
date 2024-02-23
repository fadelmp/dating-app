package mapper

import (
	"dating-app/service-user/domain"
	"dating-app/service-user/dto"
	sharedBase "dating-app/shared/domain"
)

// Interface
type SignUpMapperContract interface {
	ToTempUser(dto.SignUp, sharedBase.Base) domain.TempUser
}

// Class
type SignUpMapper struct {
}

// Constructor
func NewSignUpMapper() *SignUpMapper {
	return &SignUpMapper{}
}

// Implementation

func (m *SignUpMapper) ToTempUser(signUp dto.SignUp, base sharedBase.Base) domain.TempUser {

	return domain.TempUser{
		Email:      signUp.Email,
		Password:   signUp.Password,
		OtpCode:    "123456",
		TryAttempt: 3,
		Base:       base,
	}
}
