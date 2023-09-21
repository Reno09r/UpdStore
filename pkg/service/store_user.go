package service

import (
	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type StoreUser struct {
	repo repository.User
	logs repository.Logs
}

func NewStoreUser(repo repository.User, logs repository.Logs) *StoreUser {
	return &StoreUser{repo: repo, logs: logs}
}

func (s *StoreUser) Get(userId int) (store.User, error) {
	user, err := s.repo.Get(userId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, UserGetProfileFailAttempt+err.Error())
		return store.User{}, err
	}
	s.logs.PublishLog(ToLogs, Info, UserGetProfileSuccessful)
	return user, nil
}

func (s *StoreUser) Update(userId int, input store.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		s.logs.PublishLog(ToLogs, Warning, UserUpdateProfileFailAttempt+err.Error())
		return err
	}
	if input.NewPassword != nil {
		*input.NewPassword = GeneratePasswordHash(*input.NewPassword)
	}
	err := s.repo.Update(userId, input)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, UserUpdateProfileFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, UserUpdateProfileSuccessful)
	return nil
}

func (s *StoreUser) Delete(userId int) error {
	err := s.repo.Delete(userId)
	if err != nil {
		s.logs.PublishLog(ToLogs, Warning, UserDeleteProfileFailAttempt+err.Error())
		return err
	}
	s.logs.PublishLog(ToLogs, Info, UserDeleteProfileSuccessful)
	return nil
}
