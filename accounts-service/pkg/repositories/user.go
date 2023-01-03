package repositories

import (
	"fmt"

	"github.com/HashimovH/accounts-service/pkg/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db}
}

func (r repository) Create(user *models.User) (*models.User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return user, nil
}
