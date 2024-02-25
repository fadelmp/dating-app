package comparator

import (
	"dating-app/service-user/domain"
	"dating-app/service-user/dto"
	"dating-app/service-user/message"
	"dating-app/service-user/repository"
	"errors"
)

// Interface
type UserComparator interface {
	CheckEmail(string) error
	CheckTempEmail(string) error
	CheckTempId(string) error
	CheckOtpCode(verifyDto dto.Verify, tempUser domain.TempUser) error
}

// Class
type UserComparatorImpl struct {
	userRepo repository.UserRepository
	tempRepo repository.TempUserRepository
}

// Constructor
func NewUserComparator(
	userRepo repository.UserRepository,
	tempRepo repository.TempUserRepository,
) *UserComparatorImpl {
	return &UserComparatorImpl{
		userRepo: userRepo,
		tempRepo: tempRepo,
	}
}

// Implementation

func (c *UserComparatorImpl) CheckEmail(email string) error {

	// Get data by name
	signUp := c.userRepo.GetByEmail(email)

	// Return error if data exists
	if signUp.Id != "" {
		return errors.New(message.EmailExists)
	}

	return nil
}

func (c *UserComparatorImpl) CheckTempEmail(email string) error {

	// Get Data By Email
	signUp := c.tempRepo.GetByEmail(email)

	// Return error if data not found
	if signUp.Id != "" {
		return errors.New(message.EmailBeingVerified)
	}

	return nil
}

func (c *UserComparatorImpl) CheckTempId(id string) error {

	// Return error if data not found
	if id == "" {
		return errors.New(message.EmailNotFound)
	}

	return nil
}

func (c *UserComparatorImpl) CheckOtpCode(verifyDto dto.Verify, tempUser domain.TempUser) error {

	// Check OTP Code
	if verifyDto.OtpCode != tempUser.OtpCode {
		return errors.New(message.WrongOtpCode)
	}

	return nil
}
