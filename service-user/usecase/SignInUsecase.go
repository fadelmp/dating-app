package usecase

import (
	"dating-app/service-user/comparator"
	"dating-app/service-user/dto"
	"dating-app/service-user/mapper"
	"dating-app/service-user/message"
	"dating-app/service-user/repository"
	"dating-app/service-user/utils"
	sharedDomain "dating-app/shared/domain"
	sharedUtils "dating-app/shared/utils"
	"errors"
)

// Interface
type SignInUsecase interface {
	SignIn(dto.Sign) (string, error)
}

// Class
type SignInUsecaseImpl struct {
	mapper     mapper.UserMapper
	comparator comparator.UserComparator
	repo       repository.TempUserRepository
}

// Constructor
func NewSignInUsecase(
	repo repository.TempUserRepository,
	mapper mapper.UserMapper,
	comparator comparator.UserComparator) *SignInUsecaseImpl {
	return &SignInUsecaseImpl{
		repo:       repo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *SignInUsecaseImpl) SignIn(SignInDto dto.Sign) (string, error) {

	// Check Email Whether Exists
	if err := u.comparator.CheckEmail(SignInDto.Email); err != nil {
		return "", err
	}

	// Check Email Whether Being Verified
	if err := u.comparator.CheckTempEmail(SignInDto.Email); err != nil {
		return "", err
	}

	// Decode Password from the Encoded Password
	decodedPass, err := utils.DecodeString(SignInDto.Password)
	if err != nil {
		return "", errors.New(message.SignInFailed)
	}

	// Hash Password
	hashPass, err := utils.HashPassword(decodedPass)
	if err != nil {
		return "", errors.New(message.SignInFailed)
	}

	// Generate OTP Code
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return "", errors.New(message.SignInFailed)
	}

	// Generate UUID
	id, err := sharedUtils.GenerateUUID()
	if err != nil {
		return "", errors.New(message.SignInFailed)
	}

	// Create Base data
	base := sharedDomain.Create(SignInDto.Email)

	// Map SignIn dto to Temp User domain
	tempUser := u.mapper.ToTempUser(id, SignInDto.Email, hashPass, otpCode, base)

	// Create Temp User and return
	tempUserRow, err := u.repo.Create(tempUser)
	if err != nil {
		return "", errors.New(message.SignInFailed)
	}

	// Send this to Mail Server to Send the Otp Code

	return tempUserRow.Id, nil
}
