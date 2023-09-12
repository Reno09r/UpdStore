package repository

import (
	"github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
)

type Authentication interface {
	CreateUser(user store.User) (int, error)
	GetUser(username, password string) (store.User, error)
}

type Authorization interface {
	CurrentUserIsAdmin(userId int) (bool, error)
	CurrentUserIsCustomer(userId int) (bool, error)
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
	Insert(input store.Cart, userId int) (int, error)
	Get(userId int) ([]store.Cart, error)
	Delete(productId, userId int) error
}

type Buy interface {
	Confirm(input store.UserCardInput, userId int) error
	BuyedProducts(userId int) ([]store.BuyedProducts, error)
}

type User interface {
	Get(userId int) (store.User, error)
	Update(userId int, input store.UpdateUserInput) error
	Delete(userId int) error
}

type Repository struct {
	Authentication
	Authorization
	Catalog
	Manufacturer
	Product
	Cart
	User
	Buy
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authentication: NewAuthPostgresSQL(db),
		Authorization:  NewAuthorizationPostgresSQL(db),
		Catalog:        NewStoreCatalogPostgres(db),
		Manufacturer:   NewStoreManufacturerPostgres(db),
		Product:        NewProductPostgres(db),
		Cart:           NewStoreCartPostgres(db),
		User:           NewUserPostgres(db),
		Buy: 			NewStoreBuyPostgres(db),
	}
}
