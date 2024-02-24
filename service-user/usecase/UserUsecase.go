package usecase

import (
	"dating-app/service-user/comparator"
	"dating-app/service-user/dto"
	"dating-app/service-user/mapper"
	"dating-app/service-user/repository"
	"dating-app/service-user/utils"
	sharedDomain "dating-app/shared/domain"
)

// Interface
type UserUsecase interface {
	SignUp(dto.SignUp) error
}

// Class
type UserUsecaseImpl struct {
	mapper     mapper.UserMapper
	comparator comparator.UserComparator
	repo       repository.TempUserRepository
}

// Constructor
func NewUserUsecase(
	repo repository.TempUserRepository,
	mapper mapper.UserMapper,
	comparator comparator.UserComparator) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo:       repo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *UserUsecaseImpl) SignUp(signUpDto dto.SignUp) error {

	// Check Email Whether Exists
	if err := u.comparator.CheckEmail(signUpDto); err != nil {
		return err
	}

	// Check Email Whether Being Verified
	if err := u.comparator.CheckTempEmail(signUpDto); err != nil {
		return err
	}

	// Create Base data
	base := sharedDomain.Create(signUpDto.Email)

	// Decode Password
	decodedPass, err := utils.DecodeString(signUpDto.Password)
	if err != nil {
		return err
	}

	// Hash Password
	hashPass, err := utils.HashPassword(decodedPass)
	if err != nil {
		return err
	}

	signUpDto.Password = hashPass

	// Map SignUp dto to Temp User domain
	tempUser := u.mapper.ToTempUser(signUpDto, base)

	// Put OTP Code
	tempUser.OtpCode, err = utils.GenerateOTP()
	if err != nil {
		return err
	}

	// Create Temp User and return
	_, err = u.repo.Create(tempUser)

	return err
}
