package sales

import (
	"strive/common"
)

type SaleDetailsService struct {
	repo *SaleDetailsRepository
}

func NewSaleDetailsService(repo *SaleDetailsRepository) *SaleDetailsService {
	return &SaleDetailsService{repo: repo}
}

func (s *SaleDetailsService) CreateSaleDetail(detail *common.SaleDetails) error {
	return s.repo.Create(detail)
}

func (s *SaleDetailsService) GetSaleDetailsBySale(saleID string) ([]common.SaleDetails, error) {
	return s.repo.FindBySaleID(saleID)
}
