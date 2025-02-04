package sales

import (
	"strive/common"
)

type SalesService struct {
	repo *SalesRepository
}

func NewSalesService(repo *SalesRepository) *SalesService {
	return &SalesService{repo: repo}
}

func (s *SalesService) CreateSale(sale *common.Sale) error {
	return s.repo.Create(sale)
}

func (s *SalesService) GetAllSales() ([]common.Sale, error) {
	return s.repo.FindAll()
}

func (s *SalesService) GetSalesByUser(userID string) ([]common.Sale, error) {
	return s.repo.FindByUserID(userID)
}
