package service

import (
	"github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/repository"
)

type Authentication interface {
	CreateCustomer(user store.User) (int, error)
	CreateAdmin(user store.User) (int, error)
	GenerateToken(username, password string, isAdmin bool) (string, error)
	ParseToken(token string, isAdmin bool) (int, error)
}

type Authorization interface {
	CurrentUserIsAdmin(userId int) (bool, error)
}

type Catalog interface {
	Create(catalog store.Catalog) (int, error)
	GetAll() ([]store.Catalog, error)
	GetById(CatalogId int) (store.Catalog, error)
	Delete(CatalogId int) error
	Update(CatalogId int, input store.UpdateInput) error
}

type Manufacturer interface {
	Create(manufacturer store.Manufacturer) (int, error)
	GetAll() ([]store.Manufacturer, error)
	GetById(manufacturerId int) (store.Manufacturer, error)
	Delete(manufacturerId int) error
	Update(manufacturerId int, input store.UpdateInput) error
}

type Product interface {
	Create(product store.Product) (int, error)
	GetAll() ([]store.Product, error)
	GetById(productId int) (store.Product, error)
	Delete(productId int) error
	Update(productId int, input store.UpdateProductInput) error
	GetAllByManufacturer(manufacturerID int) ([]store.Product, error)
	GetAllByCatalog(catalogID int) ([]store.Product, error)
}

type Cart interface {
	Insert(input store.Cart, customerId int) (int, error)
	Get(customerId int) ([]store.Cart, error)
	Delete(productId, customerId int) error
}

type Service struct {
	Authentication
	Authorization
	Catalog
	Manufacturer
	Product
	Cart
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos.Authentication),
		Authorization:  NewAuthorizationService(repos.Authorization),
		Catalog:        NewStoreCatalog(repos.Catalog),
		Manufacturer:   NewStoreManufacturer(repos.Manufacturer),
		Product:        NewStoreProduct(repos.Product),
		Cart:           NewStoreCart(repos.Cart),
	}
}
