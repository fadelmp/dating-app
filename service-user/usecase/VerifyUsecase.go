package usecase

import (
	"dating-app/service-user/comparator"
	"dating-app/service-user/domain"
	"dating-app/service-user/dto"
	"dating-app/service-user/mapper"
	"dating-app/service-user/message"
	"dating-app/service-user/repository"
	sharedDomain "dating-app/shared/domain"
	sharedUtils "dating-app/shared/utils"
	"errors"
)

// Interface
type VerifyUsecase interface {
	VerifyEmail(dto.Verify) error
}

// Class
type VerifyUsecaseImpl struct {
	mapper     mapper.UserMapper
	comparator comparator.UserComparator
	tempRepo   repository.TempUserRepository
	userRepo   repository.UserRepository
}

// Constructor
func NewVerifyUsecase(
	tempRepo repository.TempUserRepository,
	userRepo repository.UserRepository,
	mapper mapper.UserMapper,
	comparator comparator.UserComparator) *VerifyUsecaseImpl {
	return &VerifyUsecaseImpl{
		tempRepo:   tempRepo,
		userRepo:   userRepo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *VerifyUsecaseImpl) VerifyEmail(verifyDto dto.Verify) error {

	// Get Temp User By Id
	tempUser := u.tempRepo.GetById(verifyDto.Id)

	// Check if Temp User Not Found
	if err := u.comparator.CheckTempId(tempUser.Id); err != nil {
		return err
	}

	// Check Email Whether Exists
	if err := u.comparator.CheckEmail(tempUser.Email); err != nil {
		return err
	}

	// Check OTP Code
	if err := u.comparator.CheckOtpCode(verifyDto, tempUser); err != nil {

		// Minus Try Attempt if OTP Code Wrong
		u.updateTemp(tempUser)

		return err
	}

	// Generate UUID
	id, err := sharedUtils.GenerateUUID()
	if err != nil {
		return errors.New(message.VerifyFailed)
	}

	// Create Base data
	base := sharedDomain.Create(tempUser.Email)

	// Map Temp User to User
	user := u.mapper.ToUser(id, tempUser, base)

	// Create Temp User and return
	if _, err := u.userRepo.Create(user); err != nil {
		return errors.New(message.VerifyFailed)
	}

	u.tempRepo.Delete(tempUser)

	return nil
}

func (u *VerifyUsecaseImpl) updateTemp(tempUser domain.TempUser) {

	tryAttempt := tempUser.TryAttempt

	if tryAttempt == 0 {

		// Delete Temp User
		u.tempRepo.Delete(tempUser)

		return
	}

	tempUser.TryAttempt = tryAttempt - 1

	u.tempRepo.Update(tempUser)
}
