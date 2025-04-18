package sales

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strive/auth"
	"strive/common"
	"strive/inventory"

	"github.com/google/uuid"
)

type SalesService struct {
	repo         *SalesRepository
	userRepo     *auth.UserRepository
	addressRepo  *auth.AddressRepository
	productRepo  *inventory.ProductRepository
	emailService *common.EmailService
}

type Item struct {
	Name     string
	Price    int
	Quantity int
}

type SalesEmail struct {
	Name          string
	Url           string
	Address       string
	CustomerEmail string
	PhoneNo       string
	OrderNo       string
	SubTotal      int
	Total         int
	// discount uint
	PackagingCharges int
	Products         []Item
}

func NewSalesService(
	repo *SalesRepository,
	userRepo *auth.UserRepository,
	addressRepo *auth.AddressRepository,
	productRepo *inventory.ProductRepository,
	emailService *common.EmailService,
) *SalesService {
	return &SalesService{
		repo:         repo,
		userRepo:     userRepo,
		emailService: emailService,
		productRepo:  productRepo,
		addressRepo:  addressRepo,
	}
}

func (s *SalesService) CreateSale(sale *PartialSale) (string, error) {

	id := uuid.New()

	salesDetail := []common.SaleDetails{}
	for _, v := range sale.SalesDetails {
		salesDetail = append(salesDetail, common.SaleDetails{
			SaleID:    id,
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	payload := common.Sale{
		ID:          id,
		UserID:      sale.UserID,
		AddressID:   sale.AddressID,
		SaleDetails: salesDetail,
	}
	return id.String(), s.repo.Create(&payload)
}

func (s *SalesService) GetAllSales() ([]common.Sale, error) {
	return s.repo.FindAll()
}

func (s *SalesService) GetSalesByUser(userID string) ([]common.Sale, error) {
	return s.repo.FindByUserID(userID)
}

func (s *SalesService) GetSalesByID(salesID string) (*common.Sale, error) {
	return s.repo.FindByID(salesID)
}

func (s *SalesService) prepareSalesEmailContent(
	user *common.User,
	address *common.Address,
	sale *common.Sale,
) *SalesEmail {
	addressString := fmt.Sprintf("%s, %s, %s, %s %s",
		address.AddressLine1,
		address.AddressLine2,
		address.City,
		address.State,
		address.Pincode)
	items := []Item{}
	ids := []string{}
	var qtys = map[string]int{}

	fmt.Println("sale")

	// Prepare product IDs
	for _, sd := range sale.SaleDetails {
		ids = append(ids, sd.ProductID.String())
		qtys[sd.ProductID.String()] = sd.Quantity
	}
	// Fetch products
	products, _ := s.productRepo.FindByIDs(ids)
	var subTotal int = 0
	// Prepare item list
	for _, p := range products {
		qty := qtys[p.ID.String()]
		items = append(items, Item{
			Name:     p.Name,
			Price:    (int(p.Price) * qty),
			Quantity: qty,
		})
		subTotal = subTotal + (int(p.Price) * qty)
	}
	fmt.Println(items)
	var packagingCharges int = 80 // @todo need to discuss this
	salesEmailContent := SalesEmail{
		Address:          addressString,
		Name:             user.FirstName + " " + user.LastName,
		Url:              sale.ID.String(), // @todo need to change this to URL after frontend route is defined
		CustomerEmail:    user.Email,
		PhoneNo:          user.Phone,
		OrderNo:          sale.ID.String(), // @todo may need to generate orderID
		SubTotal:         subTotal,
		Total:            packagingCharges + subTotal,
		PackagingCharges: packagingCharges,
		Products:         items,
	}
	return &salesEmailContent
}

func (s *SalesService) SendSalesEmail(salesID string) error {

	sale, salesErr := s.repo.FindByIDWithDetails(salesID)
	if salesErr != nil {
		return salesErr
	}

	user, userErr := s.userRepo.FindByID(sale.UserID.String())
	if userErr != nil {
		return userErr
	}

	address, addressErr := s.addressRepo.FindByID(sale.AddressID.String())
	if addressErr != nil {
		return addressErr
	}

	salesEmailContent := s.prepareSalesEmailContent(
		user,
		address,
		sale,
	)

	// Generate HTML content from template
	var htmlContent bytes.Buffer
	tmpl, err := template.ParseFiles(filepath.Join("email_templates", "order_confirmation.html"))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}
	err = tmpl.Execute(&htmlContent, salesEmailContent)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	// Create email payload
	email := common.EmailType{
		Body:       htmlContent.String(),
		Subject:    fmt.Sprintf("Order Confirmation - Order #%s", salesID),
		Recepients: []string{user.Email},
	}

	// Send email using email service
	return s.emailService.SendEmail(email)
}
