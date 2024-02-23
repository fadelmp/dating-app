package usecase

import (
	"dating-app/service-user/comparator"
	"dating-app/service-user/dto"
	"dating-app/service-user/mapper"
	"dating-app/service-user/repository"
	sharedDomain "dating-app/shared/domain"
)

// Interface
type SignUpUsecaseContract interface {
	SignUp(dto.SignUp) error
}

// Class
type SignUpUsecase struct {
	mapper     mapper.SignUpMapperContract
	comparator comparator.SignUpComparatorContract
	repo       repository.TempUserRepositoryContract
}

// Constructor
func NewSignUpUsecase(
	repo repository.TempUserRepositoryContract,
	mapper mapper.SignUpMapperContract,
	comparator comparator.SignUpComparatorContract) *SignUpUsecase {
	return &SignUpUsecase{
		repo:       repo,
		mapper:     mapper,
		comparator: comparator,
	}
}

// Implementation

func (u *SignUpUsecase) SignUp(signUpDto dto.SignUp) error {

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

	// Map SignUp dto to Temp User domain
	tempUser := u.mapper.ToTempUser(signUpDto, base)

	// Create Temp User and return
	_, err := u.repo.Create(tempUser)

	return err
}
