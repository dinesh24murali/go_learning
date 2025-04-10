package auth

import (
	"strive/common"
)

type AddressService struct {
	repo *AddressRepository
}

func NewAddressService(repo *AddressRepository) *AddressService {
	return &AddressService{repo: repo}
}

func (s *AddressService) AddAddress(address *common.Address) error {
	return s.repo.Create(address)
}

func (s *AddressService) GetAddressesByUser(userID string) ([]common.Address, error) {
	return s.repo.FindByUserID(userID)
}

func (s *AddressService) UpdateAddress(addressID string, address *common.Address) error {
	return s.repo.UpdatePartialByID(addressID, address)
}

func (s *AddressService) DeleteAddress(addressID string) error {
	return s.repo.DeleteByID(addressID)
}
