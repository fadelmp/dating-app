package repository

import (
	"dating-app/service-user/domain"
	config "dating-app/shared/config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type TempUserRepository interface {
	GetByEmail(string) domain.TempUser
	Create(domain.TempUser) (domain.TempUser, error)
	Update(domain.TempUser) (domain.TempUser, error)
	Delete(domain.TempUser) error
}

// Class
type TempUserRepositoryImpl struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewTempUserRepository(DB *gorm.DB, Redis *redis.Client) *TempUserRepositoryImpl {
	return &TempUserRepositoryImpl{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *TempUserRepositoryImpl) GetByEmail(email string) domain.TempUser {

	var tempUser domain.TempUser

	keys := "temp_user_" + email
	query := r.DB.Model(&tempUser).
		Where("is_deleted=?", false).
		Where("email=?", email).
		Find(&tempUser)

	config.Query(r.Redis, query, keys)

	return tempUser
}

func (r *TempUserRepositoryImpl) Create(tempUser domain.TempUser) (domain.TempUser, error) {

	// Create User
	err := r.DB.Create(&tempUser).Error

	return tempUser, err
}

func (r *TempUserRepositoryImpl) Update(tempUser domain.TempUser) (domain.TempUser, error) {

	// Update User by id
	err := r.DB.Model(&tempUser).Update(&tempUser).Error

	return tempUser, err
}

func (r *TempUserRepositoryImpl) Delete(tempUser domain.TempUser) error {

	// Delete Temp User by Id
	return r.DB.Model(&tempUser).Delete(&tempUser).Error
}
