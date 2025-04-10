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

func (s *UserService) RegisterUser(user *RegisterUserDto) error {
	return s.repo.Create(&common.User{
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
	})
}

func (s *UserService) UpdateUser(user *common.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUserById(ID string) (*common.User, error) {
	return s.repo.FindByID(ID)
}

func (s *UserService) GetUserByEmail(email string) (*common.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) LoginUser(loginDto LoginUserDto) (*common.User, *string) {

	user, err := s.GetUserByEmail(loginDto.Email)
	if err != nil {
		return nil, nil
	}

	if user.Password != loginDto.Password {
		return nil, nil
	}

	token, err := common.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return nil, nil
	}

	return user, &token
}

func (s *UserService) GetUserByPhone(phone string) (*common.User, error) {
	return s.repo.FindByPhone(phone)
}
