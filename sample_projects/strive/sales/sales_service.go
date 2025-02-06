package sales

import (
	"strive/common"

	"github.com/google/uuid"
)

type SalesService struct {
	repo *SalesRepository
}

func NewSalesService(repo *SalesRepository) *SalesService {
	return &SalesService{repo: repo}
}

func (s *SalesService) CreateSale(sale *PartialSale) error {

	id := uuid.New()

	salesDetail := []common.SaleDetails{}
	for _, v := range *sale.SalesDetails {
		salesDetail = append(salesDetail, common.SaleDetails{
			SaleID:    id,
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	payload := common.Sale{
		UserID:      sale.UserID,
		AddressID:   sale.AddressID,
		SaleDetails: salesDetail,
	}
	return s.repo.Create(&payload)
}

func (s *SalesService) GetAllSales() ([]common.Sale, error) {
	return s.repo.FindAll()
}

func (s *SalesService) GetSalesByUser(userID string) ([]common.Sale, error) {
	return s.repo.FindByUserID(userID)
}
