package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreProduct struct {
	repo repository.Product
}

func NewStoreProduct(repo repository.Product) *StoreProduct {
	return &StoreProduct{repo: repo}
}

func (s *StoreProduct) Create(product store.Product) (int, error) {
	return s.repo.Create(product)
}

func (s *StoreProduct) GetAll() ([]store.Product, error) {
	return s.repo.GetAll()
}

func (s *StoreProduct) GetById(productId int) (store.Product, error) {
	return s.repo.GetById(productId)
}

func (s *StoreProduct) Delete(productId int) error {
	return s.repo.Delete(productId)
}

func (s *StoreProduct) Update(productId int, input store.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(productId, input)
}

func (s *StoreProduct) GetAllByManufacturer(manufacturerID int) ([]store.Product, error) {
	return s.repo.GetAllByManufacturer(manufacturerID)
}

func (s *StoreProduct) GetAllByCatalog(catalogID int) ([]store.Product, error) {
	return s.repo.GetAllByCatalog(catalogID)
}
