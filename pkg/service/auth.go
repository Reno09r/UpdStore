package service

import (
	"errors"
	"log"
	"time"

	"github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	signingKey = "grkjk#4#%35FSFJ1ja#4353KSFjH"
	tokenTTL = 24 * time.Hour
)

type tokenClaims struct{
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authentication 
}

func NewAuthService(repo repository.Authentication ) *AuthService{
	return &AuthService{repo: repo}
}

func(s *AuthService) CreateUser(user store.User) (int, error){
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error){
	user, err := s.repo.GetUser(username, password)
	if err != nil{
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt: time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok{
		return 0, errors.New("token claims ate not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string{
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panicln("Failed to generate password hash:", err)
	}
	return string(hashed_password)

}

