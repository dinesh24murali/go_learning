package auth

import (
	"strive/common"

	"gorm.io/gorm"
)

type AddressRepository struct {
	*common.GormRepository[common.Address]
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		GormRepository: common.NewGormRepository[common.Address](db),
	}
}

func (r *AddressRepository) FindByUserID(userID string) ([]common.Address, error) {
	var addresses []common.Address
	err := r.Db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}
