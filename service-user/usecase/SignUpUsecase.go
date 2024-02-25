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
type SignUpUsecase interface {
	SignUp(dto.SignUp) (string, error)
}

// Class
type SignUpUsecaseImpl struct {
	mapper     mapper.UserMapper
	comparator comparator.UserComparator
	repo       repository.TempUserRepository
}

// Constructor
func NewSignUpUsecase(
	repo repository.TempUserRepository,
	mapper mapper.UserMapper,
	comparator comparator.UserComparator) *SignUpUsecaseImpl {
	return &SignUpUsecaseImpl{
		repo:       repo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *SignUpUsecaseImpl) SignUp(signUpDto dto.SignUp) (string, error) {

	// Check Email Whether Exists
	if err := u.comparator.CheckEmail(signUpDto.Email); err != nil {
		return "", err
	}

	// Check Email Whether Being Verified
	if err := u.comparator.CheckTempEmail(signUpDto.Email); err != nil {
		return "", err
	}

	// Decode Password from the Encoded Password
	decodedPass, err := utils.DecodeString(signUpDto.Password)
	if err != nil {
		return "", errors.New(message.SignUpFailed)
	}

	// Hash Password
	hashPass, err := utils.HashPassword(decodedPass)
	if err != nil {
		return "", errors.New(message.SignUpFailed)
	}

	// Generate OTP Code
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return "", errors.New(message.SignUpFailed)
	}

	// Generate UUID
	id, err := sharedUtils.GenerateUUID()
	if err != nil {
		return "", errors.New(message.SignUpFailed)
	}

	// Create Base data
	base := sharedDomain.Create(signUpDto.Email)

	// Map SignUp dto to Temp User domain
	tempUser := u.mapper.ToTempUser(id, signUpDto.Email, hashPass, otpCode, base)

	// Create Temp User and return
	tempUserRow, err := u.repo.Create(tempUser)
	if err != nil {
		return "", errors.New(message.SignUpFailed)
	}

	// Send this to Mail Server to Send the Otp Code

	return tempUserRow.Id, nil
}
