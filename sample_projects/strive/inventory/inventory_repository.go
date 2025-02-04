package inventory

import (
	"strive/common"

	"gorm.io/gorm"
)

type ProductRepository struct {
	*common.GormRepository[common.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		GormRepository: common.NewGormRepository[common.Product](db),
	}
}

func (r *ProductRepository) FindByCategory(categoryID string) ([]common.Product, error) {
	var products []common.Product
	err := r.Db.Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}
