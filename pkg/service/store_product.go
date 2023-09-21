package service

import (
	"fmt"

	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreProduct struct {
	repo repository.Product
	logs repository.Logs
}

func NewStoreProduct(repo repository.Product, logs repository.Logs) *StoreProduct {
	return &StoreProduct{repo: repo, logs: logs}
}

func (s *StoreProduct) Create(product store.Product) (int, error) {
	id, err := s.repo.Create(product)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductCreateFailAttempt+err.Error())
		return 0, err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(ProductCreateSuccessful, id, product.ProductName))
	return id, nil
}

func (s *StoreProduct) GetAll() ([]store.Product, error) {
	product, err := s.repo.GetAll()
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductGetFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, ProductGetSuccessful)
	return product, nil
}

func (s *StoreProduct) GetById(ProductId int) (store.Product, error) {
	product, err := s.repo.GetById(ProductId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductGetFailAttempt+err.Error())
		return store.Product{}, err
	}
	s.logs.PublishLog(ToLogs, Info, ProductGetSuccessful)
	return product, nil
}

func (s *StoreProduct) Delete(ProductId int) error {
 	err := s.repo.Delete(ProductId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductDeleteFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(ProductDeleteSuccessful, ProductId))
	return nil
}

func (s *StoreProduct) Update(ProductId int, input store.UpdateProductInput) error {
	if err := input.Validate(); err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductUpdateFailAttempt+err.Error())
		return err
	}
	err := s.repo.Update(ProductId, input)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductUpdateFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(ProductUpdateSuccessful, ProductId))
	return nil
}


func (s *StoreProduct) GetAllByManufacturer(manufacturerID int) ([]store.Product, error) {
	product, err := s.repo.GetAllByManufacturer(manufacturerID)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductGetFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, ProductGetSuccessful)
	return product, nil
}

func (s *StoreProduct) GetAllByCatalog(catalogID int) ([]store.Product, error) {
	product, err := s.repo.GetAllByManufacturer(catalogID)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ProductGetFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, ProductGetSuccessful)
	return product, nil
}

