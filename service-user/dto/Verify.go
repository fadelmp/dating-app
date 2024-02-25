package dto

import (
	dto "dating-app/shared/dto"
)

// Payload Design
type Verify struct {
	dto.Base
	Id      string `json:"id" validate:"required"`
	OtpCode string `json:"otp_code" validate:"required,len=6" validateErrorMsg:"Invalid OTP Code"`
}
