package repositories

import (
	"context"
	"fmt"

	"github.com/go-backend/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) RegisterUser(context context.Context, registerData *models.AuthCredentials) (*models.User, error) {

		user := &models.User{
			Email: registerData.Email,
			Password: registerData.Password,
		}

		res := r.db.Model(&models.User{}).Create(&user)

		if res.Error  != nil {
			return nil, res.Error
		}

		fmt.Println("User created successfully")

		return user, nil

}

func (r *AuthRepository) GetUserByEmail(context context.Context, email string) (*models.User, error) {

	user := &models.User{}

	res := r.db.Model(&models.User{}).Where("email = ?", email).First(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil

}

