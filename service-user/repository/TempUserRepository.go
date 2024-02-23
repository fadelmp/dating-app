package repository

import (
	"dating-app/service-user/domain"
	config "dating-app/shared/config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type TempUserRepositoryContract interface {
	GetByEmail(string) domain.TempUser
	Create(domain.TempUser) (domain.TempUser, error)
	Update(domain.TempUser) (domain.TempUser, error)
	Delete(domain.TempUser) error
}

// Class
type TempUserRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewTempUserRepository(DB *gorm.DB, Redis *redis.Client) *TempUserRepository {
	return &TempUserRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *TempUserRepository) GetByEmail(email string) domain.TempUser {

	var tempUser domain.TempUser

	keys := "temp_user_" + email
	query := r.DB.Model(&tempUser).
		Where("is_deleted=?", false).
		Where("email=?", email).
		Find(&tempUser)

	config.Query(r.Redis, query, keys)

	return tempUser
}

func (r *TempUserRepository) Create(tempUser domain.TempUser) (domain.TempUser, error) {

	// Create User
	err := r.DB.Create(&tempUser).Error

	return tempUser, err
}

func (r *TempUserRepository) Update(tempUser domain.TempUser) (domain.TempUser, error) {

	// Update User by id
	err := r.DB.Model(&tempUser).Update(&tempUser).Error

	return tempUser, err
}

func (r *TempUserRepository) Delete(tempUser domain.TempUser) error {

	// Delete Temp User by Id
	return r.DB.Model(&tempUser).Delete(&tempUser).Error
}
