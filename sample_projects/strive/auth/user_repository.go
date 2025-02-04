package auth

import (
	"strive/common"

	"gorm.io/gorm"
)

type UserRepository struct {
	*common.GormRepository[common.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		GormRepository: common.NewGormRepository[common.User](db),
	}
}

func (r *UserRepository) FindByEmail(email string) (*common.User, error) {
	var user common.User
	err := r.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhone(phone string) (*common.User, error) {
	var user common.User
	err := r.Db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
