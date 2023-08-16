package repository

import (
	"github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
)

type Authentication interface {
	CreateCustomer(user store.User) (int, error)
	GetCustomer(username, password string) (store.User, error)
	CreateAdmin(user store.User) (int, error)
	GetAdmin(username, password string) (store.User, error)
}

type Authorization interface{
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

type Cart interface{
	Insert(input store.Cart, customerId int) (int, error)
	Get(customerId int) ([]store.Cart, error)
	Delete(productId, customerId int) error
}

type Repository struct {
	Authentication 
	Authorization
	Catalog
	Manufacturer
	Product
	Cart
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authentication: NewAuthPostgresSQL(db),
		Authorization: NewAuthorizationPostgresSQL(db),
		Catalog: NewStoreCatalogPostgres(db),
		Manufacturer: NewStoreManufacturerPostgres(db),
		Product: NewProductPostgres(db),
		Cart: NewStoreCartPostgres(db),
	}
}