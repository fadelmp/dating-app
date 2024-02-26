package dto

import (
	dto "dating-app/shared/dto"
)

// Payload Design
type Sign struct {
	dto.Base
	Email        string `json:"email" validate:"required,email" validateErrorMsg:"Invalid Email Format"`
	Password     string `json:"password" validate:"required" validateErrorMsg:"Password Required"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
