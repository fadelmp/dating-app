package dto

import (
	dto "dating-app/shared/dto"
)

// Payload Design
type SignUp struct {
	dto.Base
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
