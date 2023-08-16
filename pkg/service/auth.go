package service

import (
	"errors"
	"log"
	"time"

	"github.com/Reno09r/Store"
	"github.com/Reno09r/Store/server/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	signingKey = "grkjk#4#%35FSFJ1ja#4353KSFjH"
	tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	CustomerId int `json:"customer_id"`
}

type adminTokenClaims struct {
	jwt.StandardClaims
	AdminId int `json:"admin_id"`
}

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateCustomer(user store.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateCustomer(user)
}
func (s *AuthService) CreateAdmin(user store.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateAdmin(user)
}
func (s *AuthService) GenerateToken(username, password string, isAdmin bool) (string, error) {
	var user store.User
	var err error

	if isAdmin {
		user, err = s.repo.GetAdmin(username, password)
	} else {
		user, err = s.repo.GetCustomer(username, password)
	}

	if err != nil {
		return "", err
	}

	var claims jwt.Claims
	if isAdmin {
		claims = &adminTokenClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			AdminId: user.AdminId,
		}
	} else {
		claims = &tokenClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			CustomerId: user.Id,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string, isAdmin bool) (int, error) {
	var claims jwt.Claims
	if isAdmin {
		claims = &adminTokenClaims{}
	} else {
		claims = &tokenClaims{}
	}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	var userID int
	if isAdmin {
		adminClaims, ok := token.Claims.(*adminTokenClaims)
		if !ok {
			return 0, errors.New("token claims are not of type *adminTokenClaims")
		}
		userID = adminClaims.AdminId
	} else {
		userClaims, ok := token.Claims.(*tokenClaims)
		if !ok {
			return 0, errors.New("token claims are not of type *tokenClaims")
		}
		userID = userClaims.CustomerId
	}

	return userID, nil
}

func generatePasswordHash(password string) string {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panicln("Failed to generate password hash:", err)
	}
	return string(hashed_password)

}
