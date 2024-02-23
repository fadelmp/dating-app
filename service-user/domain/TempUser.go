package domain

import (
	domain "dating-app/shared/domain"
)

// Database Design
type TempUser struct {
	domain.Base
	Id         uint   `gorm:"primaryKey;autoIncrement:true;Index"`
	Email      string `gorm:"type:VARCHAR(255);NOT NULL"`
	Password   string `gorm:"type:VARCHAR(255);NOT NULL"`
	OtpCode    string `gorm:"type:VARCHAR(6);NOT NULL"`
	TryAttempt int    `gorm:"type:integer"`
}
