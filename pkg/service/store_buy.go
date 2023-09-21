package service

import (
	"fmt"

	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreBuy struct {
	repo repository.Buy
	logs repository.Logs
}

func NewStoreBuy(repo repository.Buy, logs repository.Logs) *StoreBuy {
	return &StoreBuy{repo: repo, logs: logs}
}

func (s *StoreBuy) Confirm(input store.UserCardInput, userId int) error{
	if err := input.Validate(); err != nil {
		s.logs.PublishLog(ToLogs, Warning, BuyProductsFailAttempt+err.Error())
		return err
	}
	err := s.repo.Confirm(input, userId)
	if err != nil{
		s.logs.PublishLog(ToLogs, Warning, BuyProductsFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(BuyProductsSuccessful, userId))
	return err
}

func (s *StoreBuy) BoughtProducts(userId int) ([]store.BoughtProducts, error){
	boughtProducts, err := s.repo.BoughtProducts(userId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, GetBoughtProductsFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(GetBoughtProductsSuccessful, userId))
	return boughtProducts, nil
}

