package usecase

import (
	"echo-go-api/models"
	"echo-go-api/repository"
	"errors"
	"strconv"
)

type IUserUsecase interface {
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) error
	GetUsers() ([]*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
}

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) IUserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUser(id int) (*models.User, error) {
	return u.repo.GetUser(id)
}

func (u *UserUsecase) CreateUser(user *models.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserUsecase) GetUsers() ([]*models.User, error) {
	return u.repo.GetUsers()
}

func (u *UserUsecase) UpdateUser(id string, user *models.User) error {
	return u.repo.UpdateUser(id, user)
}

func (u *UserUsecase) DeleteUser(id string) error {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New(err.Error())
	}

	user, err := u.repo.GetUser(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	return u.repo.DeleteUser(id)
}
