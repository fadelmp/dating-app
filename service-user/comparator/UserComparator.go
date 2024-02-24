package comparator

import (
	"dating-app/service-user/dto"
	"dating-app/service-user/message"
	"dating-app/service-user/repository"
	"errors"
)

// Interface
type UserComparator interface {
	CheckEmail(dto.SignUp) error
	CheckTempEmail(dto.SignUp) error
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

func (c *UserComparatorImpl) CheckEmail(signUp dto.SignUp) error {

	// Get data by name
	SignUp := c.userRepo.GetByEmail(signUp.Email)

	// Return error if data exists
	if SignUp.Id != 0 {
		return errors.New(message.EmailExists)
	}

	return nil
}

func (c *UserComparatorImpl) CheckTempEmail(signUp dto.SignUp) error {

	// Get data by Id
	SignUp := c.tempRepo.GetByEmail(signUp.Email)

	// Return error if data not found
	if SignUp.Id != 0 {
		return errors.New(message.EmailBeingVerified)
	}

	return nil
}
