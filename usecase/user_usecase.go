package usecase

import (
	"echo-go-api/models"
	"echo-go-api/repository"
)

type UserUsecase interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error) // New method
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUserByID(id int) (*models.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userUsecase) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

// New method to get all users
func (u *userUsecase) GetAllUsers() ([]*models.User, error) {
	return u.repo.GetAllUsers()
}
