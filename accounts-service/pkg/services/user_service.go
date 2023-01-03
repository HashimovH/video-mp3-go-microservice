package services

import "github.com/HashimovH/accounts-service/pkg/models"

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewService(ur UserRepository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	result, err := us.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
