package sales

import "github.com/google/uuid"

type PartialSale struct {
	UserID       uuid.UUID
	AddressID    uuid.UUID
	SalesDetails []ProductList
}
