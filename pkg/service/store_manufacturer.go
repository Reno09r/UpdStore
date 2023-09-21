package service

import (
	"fmt"

	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreManufacturer struct {
	repo repository.Manufacturer
	logs repository.Logs
}

func NewStoreManufacturer(repo repository.Manufacturer, logs repository.Logs) *StoreManufacturer {
	return &StoreManufacturer{repo: repo, logs: logs}
}

func (s *StoreManufacturer) Create(manufacturer store.Manufacturer) (int, error) {
	id, err := s.repo.Create(manufacturer)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ManufacturerCreateFailAttempt+err.Error())
		return 0, err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(ManufacturerCreateSuccessful, id, manufacturer.Title))
	return id, nil
}

func (s *StoreManufacturer) GetAll() ([]store.Manufacturer, error) {
	manufacturer, err := s.repo.GetAll()
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ManufacturerGetFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, ManufacturerGetSuccessful)
	return manufacturer, nil
}

func (s *StoreManufacturer) GetById(ManufacturerId int) (store.Manufacturer, error) {
	manufacturer, err := s.repo.GetById(ManufacturerId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ManufacturerGetFailAttempt+err.Error())
		return store.Manufacturer{}, err
	}
	s.logs.PublishLog(ToLogs, Info, ManufacturerGetSuccessful)
	return manufacturer, nil
}

func (s *StoreManufacturer) Delete(ManufacturerId int) error {
 	err := s.repo.Delete(ManufacturerId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ManufacturerDeleteFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(ManufacturerDeleteSuccessful, ManufacturerId))
	return nil
}

func (s *StoreManufacturer) Update(ManufacturerId int, input store.UpdateInput) error {
	if err := input.Validate(); err != nil {
		s.logs.PublishLog(ToLogs, Warning, ManufacturerUpdateFailAttempt+err.Error())
		return err
	}
	err := s.repo.Update(ManufacturerId, input)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, ManufacturerUpdateFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info,fmt.Sprintf(ManufacturerUpdateSuccessful, ManufacturerId))
	return nil
}
