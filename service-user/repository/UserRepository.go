package repository

import (
	"dating-app/service-user/domain"
	config "dating-app/shared/config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Interface
type UserRepositoryContract interface {
	GetByEmail(string) domain.User
	Create(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
}

// Class
type UserRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Constructor
func NewUserRepository(DB *gorm.DB, Redis *redis.Client) *UserRepository {
	return &UserRepository{
		DB:    DB,
		Redis: Redis,
	}
}

// Implementation

func (r *UserRepository) GetByEmail(email string) domain.User {

	var user domain.User

	keys := "user_" + email
	query := r.DB.Model(&user).
		Where("is_deleted=?", false).
		Where("email=?", email).
		Find(&user)

	config.Query(r.Redis, query, keys)

	return user
}

func (r *UserRepository) Create(User domain.User) (domain.User, error) {

	// Create User
	err := r.DB.Create(&User).Error

	return User, err
}

func (r *UserRepository) Update(User domain.User) (domain.User, error) {

	// Update User by id
	err := r.DB.Model(&User).Update(&User).Error

	return User, err
}
