package dto

import (
	dto "dating-app/shared/dto"
)

// Payload Design
type SignUp struct {
	dto.Base
	Id       uint   `json:"id"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
