package service

import (
	"github.com/Reno09r/Store"
	"github.com/Reno09r/Store/server/repository"
)

type Authentication interface{
	CreateUser(user store.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Product interface{
	Create(product store.Product) (int, error)
	GetAll() ([]store.Product, error)
	GetById(productId int) (store.Product, error)
	Delete(productId int) error
	Update(productId int, input store.UpdateProductInput) error
	GetAllByManufacturer(manufacturerID int) ([]store.Product, error)
	GetAllByCatalog(catalogID int) ([]store.Product, error)
}

type Cart interface{
	Insert(input store.Cart, customerId int) (int, error)
	Get(customerId int) ([]store.Cart, error)
	Delete(productId, customerId int) error
}

type Service struct{
	Authentication 
	Catalog
	Manufacturer
	Product
	Cart
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Authentication: NewAuthService(repos.Authentication),
		Catalog: NewStoreCatalog(repos.Catalog),
		Manufacturer: NewStoreManufacturer(repos.Manufacturer),
		Product: NewStoreProduct(repos.Product),
		Cart: NewStoreCart(repos.Cart),
	}
}