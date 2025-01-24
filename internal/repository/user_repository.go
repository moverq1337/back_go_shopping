package repository

import (
	"product_api/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := r.db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}
