package service

import (
	"fmt"

	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreCatalog struct {
	repo repository.Catalog
	logs repository.Logs
}

func NewStoreCatalog(repo repository.Catalog, logs repository.Logs) *StoreCatalog {
	return &StoreCatalog{repo: repo, logs: logs}
}

func (s *StoreCatalog) Create(catalog store.Catalog) (int, error) {
	id, err := s.repo.Create(catalog)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, CatalogCreateFailAttempt+err.Error())
		return 0, err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(CatalogCreateSuccessful, id, catalog.Title))
	return id, nil
}

func (s *StoreCatalog) GetAll() ([]store.Catalog, error) {
	catalog, err := s.repo.GetAll()
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, CatalogGetFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, CatalogGetSuccessful)
	return catalog, nil
}

func (s *StoreCatalog) GetById(CatalogId int) (store.Catalog, error) {
	catalog, err := s.repo.GetById(CatalogId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, CatalogGetFailAttempt+err.Error())
		return store.Catalog{}, err
	}
	s.logs.PublishLog(ToLogs, Info, CatalogGetSuccessful)
	return catalog, nil
}

func (s *StoreCatalog) Delete(CatalogId int) error {
 	err := s.repo.Delete(CatalogId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, CatalogDeleteFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(CatalogDeleteSuccessful, CatalogId))
	return nil
}

func (s *StoreCatalog) Update(CatalogId int, input store.UpdateInput) error {
	if err := input.Validate(); err != nil {
		s.logs.PublishLog(ToLogs, Warning, CatalogUpdateFailAttempt+err.Error())
		return err
	}
	err := s.repo.Update(CatalogId, input)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, CatalogUpdateFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(CatalogUpdateSuccessful, CatalogId))
	return nil
}
