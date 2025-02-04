package sales

import (
	"strive/common"

	"gorm.io/gorm"
)

type SaleDetailsRepository struct {
	*common.GormRepository[common.SaleDetails]
}

func NewSaleDetailsRepository(db *gorm.DB) *SaleDetailsRepository {
	return &SaleDetailsRepository{
		GormRepository: common.NewGormRepository[common.SaleDetails](db),
	}
}

func (r *SaleDetailsRepository) FindBySaleID(saleID string) ([]common.SaleDetails, error) {
	var details []common.SaleDetails
	err := r.Db.Where("sale_id = ?", saleID).Find(&details).Error
	return details, err
}
