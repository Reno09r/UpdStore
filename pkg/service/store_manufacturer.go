package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/server/repository"
)

type StoreManufacturer struct {
	repo repository.Manufacturer
}

func NewStoreManufacturer(repo repository.Manufacturer) *StoreManufacturer {
	return &StoreManufacturer{repo: repo}
}

func (s *StoreManufacturer) Create(manufacturer store.Manufacturer) (int, error) {
	return s.repo.Create(manufacturer)
}

func (s *StoreManufacturer) GetAll() ([]store.Manufacturer, error) {
	return s.repo.GetAll()
}

func (s *StoreManufacturer) GetById(manufacturerId int) (store.Manufacturer, error) {
	return s.repo.GetById(manufacturerId)
}

func (s *StoreManufacturer) Delete(manufacturerId int) error {
	return s.repo.Delete(manufacturerId)
}

func (s *StoreManufacturer) Update(manufacturerId int, input store.UpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(manufacturerId, input)
}
