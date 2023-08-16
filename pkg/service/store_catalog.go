package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/server/repository"
)

type StoreCatalog struct {
	repo repository.Catalog
}

func NewStoreCatalog(repo repository.Catalog) *StoreCatalog {
	return &StoreCatalog{repo: repo}
}

func (s *StoreCatalog) Create(catalog store.Catalog) (int, error) {
	return s.repo.Create(catalog)
}

func (s *StoreCatalog) GetAll() ([]store.Catalog, error) {
	return s.repo.GetAll()
}

func (s *StoreCatalog) GetById(CatalogId int) (store.Catalog, error) {
	return s.repo.GetById(CatalogId)
}

func (s *StoreCatalog) Delete(CatalogId int) error {
	return s.repo.Delete(CatalogId)
}

func (s *StoreCatalog) Update(CatalogId int, input store.UpdateInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(CatalogId, input)
}
