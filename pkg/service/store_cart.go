package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreCart struct {
	repo repository.Cart
}

func NewStoreCart(repo repository.Cart) *StoreCart {
	return &StoreCart{repo: repo}
}

func (s *StoreCart) Insert(input store.Cart, userId int) (int, error) {
	return s.repo.Insert(input, userId)
}

func (s *StoreCart) Get(userId int) ([]store.Cart, error) {
	return s.repo.Get(userId)
}

func (s *StoreCart) Delete(productId, userId int) error {
	return s.repo.Delete(productId, userId)
}
