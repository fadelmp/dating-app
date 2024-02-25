package usecase

// import (
// 	"dating-app/service-user/comparator"
// 	"dating-app/service-user/dto"
// 	"dating-app/service-user/mapper"
// 	"dating-app/service-user/message"
// 	"dating-app/service-user/repository"
// 	"dating-app/service-user/utils"
// 	sharedDomain "dating-app/shared/domain"
// 	sharedUtils "dating-app/shared/utils"
// 	"errors"
// )

// // Interface
// type VerifyEmailUsecase interface {
// 	VerifyEmail(dto.Verify) error
// }

// // Class
// type VerifyEmailUsecaseImpl struct {
// 	mapper     mapper.UserMapper
// 	comparator comparator.UserComparator
// 	tempRepo   repository.TempUserRepository
// 	userRepo   repository.UserRepository
// }

// // Constructor
// func NewVerifyEmailUsecase(
// 	tempRepo repository.TempUserRepository,
// 	userRepo repository.UserRepository,
// 	mapper mapper.UserMapper,
// 	comparator comparator.UserComparator) *VerifyEmailUsecaseImpl {
// 	return &VerifyEmailUsecaseImpl{
// 		tempRepo:   tempRepo,
// 		userRepo:   userRepo,
// 		mapper:     mapper,
// 		comparator: comparator,
// 	}
// }

// // Implementation

// func (u *VerifyEmailUsecaseImpl) VerifyEmail(verifyDto dto.Verify) error {

// 	// Check Email Whether Exists
// 	if err := u.comparator.CheckEmail(verifyDto.ema); err != nil {
// 		return err
// 	}

// 	// Check Email Whether Being Verified
// 	if err := u.comparator.CheckTempEmail(VerifyEmailDto); err != nil {
// 		return err
// 	}

// 	// Decode Password from the Encoded Password
// 	decodedPass, err := utils.DecodeString(VerifyEmailDto.Password)
// 	if err != nil {
// 		return errors.New(message.VerifyEmailFailed)
// 	}

// 	// Hash Password
// 	hashPass, err := utils.HashPassword(decodedPass)
// 	if err != nil {
// 		return errors.New(message.VerifyEmailFailed)
// 	}

// 	// Generate OTP Code
// 	otpCode, err := utils.GenerateOTP()
// 	if err != nil {
// 		return errors.New(message.VerifyEmailFailed)
// 	}

// 	// Generate UUID
// 	id, err := sharedUtils.GenerateUUID()
// 	if err != nil {
// 		return errors.New(message.VerifyEmailFailed)
// 	}

// 	// Create Base data
// 	base := sharedDomain.Create(VerifyEmailDto.Email)

// 	// Map VerifyEmail dto to Temp User domain
// 	tempUser := u.mapper.ToTempUser(id, VerifyEmailDto.Email, hashPass, otpCode, base)

// 	// Create Temp User and return
// 	if _, err := u.repo.Create(tempUser); err != nil {
// 		return errors.New(message.VerifyEmailFailed)
// 	}

// 	// Send this to Mail Server to Send the Otp Code

// 	return nil
// }
