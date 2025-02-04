package auth

import (
	"strive/common"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *common.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUserByEmail(email string) (*common.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) GetUserByPhone(phone string) (*common.User, error) {
	return s.repo.FindByPhone(phone)
}
