package dto

import (
	dto "dating-app/shared/dto"
)

// Payload Design
type Verification struct {
	dto.Base
	Id      uint   `json:"id"`
	Email   string `json:"email"`
	OtpCode string `json:"otp_code"`
}
