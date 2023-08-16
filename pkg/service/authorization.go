package service

import (
	"github.com/Reno09r/Store/pkg/repository"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization ) *AuthorizationService{
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService)CurrentUserIsAdmin(userId int) (bool, error){
	return s.repo.CurrentUserIsAdmin(userId)
}