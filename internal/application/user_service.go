package application

import (
	"errors"

	"auction-systems/internal/domain"
	"auction-systems/internal/domain/repositories"
)

type UserService struct {
	repo *repositories.Repositories
}

func NewUserService(repo *repositories.Repositories) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) RegisterClient(user *domain.User) (int, error) {
	existingUser, err := us.repo.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		return 0, err
	}

	if existingUser != nil {
		return 0, errors.New("error user already exists")
	}

	userId, err := us.repo.UserRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (us *UserService) UpdateUserBalance(userId int, amount float64) error {
	_, err := us.repo.UserRepository.GetUserById(userId, domain.BuyerRole)
	if err != nil {
		return err
	}

	exists, err := us.repo.UserRepository.UpdateUserBalance(userId, amount, domain.PlusSymbol)
	if err != nil {
		return err
	}
	if !exists {
		err = us.repo.UserRepository.AddUserBalance(userId, amount)
		if err != nil {
			return err
		}
	}
	return nil
}
