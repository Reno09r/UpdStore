package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/server/repository"
)

type StoreCart struct {
	repo repository.Cart
}

func NewStoreCart(repo repository.Cart) *StoreCart {
	return &StoreCart{repo: repo}
}

func (s *StoreCart) Insert(input store.Cart, customerId int) (int, error){
	return s.repo.Insert(input, customerId)
}

func (s *StoreCart) Get(customerId int) ([]store.Cart, error) {
	return s.repo.Get(customerId)
}

func (s *StoreCart) Delete(productId, customerId int) error {
	return s.repo.Delete(productId, customerId)
}