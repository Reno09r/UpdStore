package service

import (
	"fmt"

	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreCart struct {
	repo repository.Cart
	logs repository.Logs
}

func NewStoreCart(repo repository.Cart, logs repository.Logs) *StoreCart {
	return &StoreCart{repo: repo, logs: logs}
}

func (s *StoreCart) Insert(input store.Cart, userId int) (int, error) {
	id, err := s.repo.Insert(input, userId)
	if err != nil{
		s.logs.PublishLog(ToLogs, Warning, ProductToCartFailAttempt+err.Error())
		return 0, err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(ProductToCartSuccessful, userId))
	return id, nil
}

func (s *StoreCart) Get(userId int) ([]store.Cart, error) {
	cart, err := s.repo.Get(userId)
	if err != nil{
		s.logs.PublishLog(ToLogs, Warning, GetCartFailAttempt+err.Error())
		return nil, err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(GetCartSuccessful, userId))
	return cart, nil
}

func (s *StoreCart) Delete(productId, userId int) error {
	err := s.repo.Delete(productId, userId)
	if err != nil{
		s.logs.PublishLog(ToLogs, Warning, DeleteProductInCartFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, fmt.Sprintf(DeleteProductInCartSuccessful, userId))
	return nil
}
