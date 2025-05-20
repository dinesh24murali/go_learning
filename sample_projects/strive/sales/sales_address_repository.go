package sales

import (
	"strive/common"

	"gorm.io/gorm"
)

type SalesAddressRepository struct {
	*common.GormRepository[common.SalesAddress]
}

func NewSalesAddressRepository(db *gorm.DB) *SalesAddressRepository {
	return &SalesAddressRepository{
		GormRepository: common.NewGormRepository[common.SalesAddress](db),
	}
}

func (r *SalesAddressRepository) FindByUserID(userID string) ([]common.SalesAddress, error) {
	var addresses []common.SalesAddress
	err := r.Db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}
