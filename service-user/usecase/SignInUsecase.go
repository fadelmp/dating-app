package usecase

import (
	"dating-app/service-user/comparator"
	"dating-app/service-user/dto"
	"dating-app/service-user/mapper"
	"dating-app/service-user/message"
	"dating-app/service-user/repository"
	"dating-app/service-user/utils"
	sharedDomain "dating-app/shared/domain"
	"errors"
)

// Interface
type SignInUsecase interface {
	SignIn(dto.Sign) (dto.Sign, error)
}

// Class
type SignInUsecaseImpl struct {
	mapper     mapper.UserMapper
	comparator comparator.UserComparator
	repo       repository.UserRepository
}

// Constructor
func NewSignInUsecase(
	repo repository.UserRepository,
	mapper mapper.UserMapper,
	comparator comparator.UserComparator) *SignInUsecaseImpl {
	return &SignInUsecaseImpl{
		repo:       repo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *SignInUsecaseImpl) SignIn(signDto dto.Sign) (dto.Sign, error) {

	// Decode Password from the Encoded Password
	decodedPass, err := utils.DecodeString(signDto.Password)
	if err != nil {
		return signDto, errors.New(message.SignInFailed)
	}

	// Hash Password
	hashPass, err := utils.HashPassword(decodedPass)
	if err != nil {
		return signDto, errors.New(message.SignInFailed)
	}

	// Get User By Email
	user := u.repo.GetByEmail(signDto.Email)

	// Check Login Data
	if err := u.comparator.CheckLogin(user, hashPass); err != nil {
		return signDto, err
	}

	// Create Base data
	base := sharedDomain.Update(signDto.Email)

	// Map SignIn dto to Temp User domain
	user = u.mapper.Login(user, base)

	// Update User Data
	if _, err := u.repo.Update(user); err != nil {
		return signDto, errors.New(message.SignInFailed)
	}

	// Generate Access and Token Refresh
	signDto.Password = ""
	signDto.AccessToken, _ = utils.AccessToken(user.Id, user.Email)
	signDto.RefreshToken, _ = utils.RefreshToken(user.Id, user.Email)

	return signDto, nil
}
