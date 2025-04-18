package sales

import (
	"strive/common"

	"gorm.io/gorm"
)

type SalesRepository struct {
	*common.GormRepository[common.Sale]
}

func NewSalesRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{
		GormRepository: common.NewGormRepository[common.Sale](db),
	}
}

func (r *SalesRepository) FindByUserID(userID string) ([]common.Sale, error) {
	var sales []common.Sale
	err := r.Db.Where("user_id = ?", userID).Find(&sales).Error
	return sales, err
}

func (r *SalesRepository) FindByIDWithDetails(salesID string) (*common.Sale, error) {
	var sale common.Sale
	err := r.Db.Preload("SaleDetails").First(&sale, "id = ?", salesID).Error
	return &sale, err
}
