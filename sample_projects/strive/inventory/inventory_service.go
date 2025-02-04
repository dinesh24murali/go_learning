package inventory

import (
	"strive/common"
)

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *common.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetAllProducts() ([]common.Product, error) {
	return s.repo.FindAll()
}
