package dto

import (
	dto "dating-app/shared/dto"
)

// Payload Design
type SignUp struct {
	dto.Base
	Id       string `json:"id"`
	Email    string `json:"email" validate:"required,email" validateErrorMsg:"Invalid Email Format"`
	Password string `json:"password" validate:"required" validateErrorMsg:"Password Required"`
}
