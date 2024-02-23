package domain

import (
	domain "dating-app/shared/domain"
	"time"
)

// Database Design
type User struct {
	domain.Base
	Id        uint      `gorm:"primaryKey;autoIncrement:true;Index"`
	Email     string    `gorm:"type:VARCHAR(255);NOT NULL"`
	Password  string    `gorm:"type:VARCHAR(255);NOT NULL"`
	LastLogin time.Time `gorm:"type:Datetime"`
}
