package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreBuy struct {
	repo repository.Buy
}

func NewStoreBuy(repo repository.Buy) *StoreBuy {
	return &StoreBuy{repo: repo}
}

func (s *StoreBuy) Confirm(input store.UserCardInput, userId int) error{
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Confirm(input, userId)
}

func (s *StoreBuy) BuyedProducts(userId int) ([]store.BuyedProducts, error){
	return s.repo.BuyedProducts(userId)
}

