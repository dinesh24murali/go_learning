package inventory

import (
	"strive/common"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	*common.GormRepository[common.Category]
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		GormRepository: common.NewGormRepository[common.Category](db),
	}
}
