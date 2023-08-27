package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreUser struct {
	repo repository.User
}

func NewStoreUser(repo repository.User) *StoreUser {
	return &StoreUser{repo: repo}
}

func (s *StoreUser) Get(userId int) (store.User, error) {
	return s.repo.Get(userId)
}

func (s *StoreUser) Update(userId int, input store.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	if input.NewPassword != nil{
		*input.NewPassword = GeneratePasswordHash(*input.NewPassword)
	}
	return s.repo.Update(userId, input)
}

func (s *StoreUser) Delete(userId int) error {
	return s.repo.Delete(userId)
}
