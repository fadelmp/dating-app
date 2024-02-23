package comparator

import (
	"dating-app/service-user/dto"
	"dating-app/service-user/message"
	"dating-app/service-user/repository"
	"errors"
)

// Interface
type SignUpComparatorContract interface {
	CheckEmail(dto.SignUp) error
	CheckTempEmail(dto.SignUp) error
}

// Class
type SignUpComparator struct {
	userRepo repository.UserRepositoryContract
	tempRepo repository.TempUserRepositoryContract
}

// Constructor
func NewSignUpComparator(
	userRepo repository.UserRepositoryContract,
	tempRepo repository.TempUserRepositoryContract,
) *SignUpComparator {
	return &SignUpComparator{
		userRepo: userRepo,
		tempRepo: tempRepo,
	}
}

// Implementation

func (c *SignUpComparator) CheckEmail(signUp dto.SignUp) error {

	// Get data by name
	SignUp := c.userRepo.GetByEmail(signUp.Email)

	// Return error if data exists
	if SignUp.Id != 0 {
		return errors.New(message.EmailExists)
	}

	return nil
}

func (c *SignUpComparator) CheckTempEmail(signUp dto.SignUp) error {

	// Get data by Id
	SignUp := c.tempRepo.GetByEmail(signUp.Email)

	// Return error if data not found
	if SignUp.Id != 0 {
		return errors.New(message.EmailBeingVerified)
	}

	return nil
}
