package store

import (
	"errors"
	"time"
)

type Product struct {
	ID               int     `json:"id" db:"product_id"`
	ProductName      string  `json:"title" db:"product_name"`
	ManufacturerName string  `json:"manufacturer" db:"manufacturer_name"`
	CatalogName      string  `json:"catalog" db:"catalog_name"`
	Price            float64 `json:"price" db:"new_price"`
}

type Cart struct {
	PurchaseID int `json:"purchase_id" db:"purchase_id"`
	Product 
	Quantity int `json:"count" db:"product_count"`
}

type BoughtProducts struct {
	BuyID int `json:"id" db:"buy_id"`
	TimeBuy time.Time `json:"time_buy" db:"buy_date"`
	Product
	Quantity  int     `json:"count" db:"product_count"`
	FullPrice float64 `json:"fullprice" db:"full_price"`
}

type Catalog struct {
	Id    int    `json:"id" db:"catalog_id"`
	Title string `json:"title" db:"catalog_name" binding:"required"`
}

type Manufacturer struct {
	Id    int    `json:"id" db:"manufacturer_id"`
	Title string `json:"title" db:"manufacturer_name" binding:"required"`
}

type UpdateProductInput struct {
	ProductName      *string  `json:"title"`
	ManufacturerName *string  `json:"manufacturer"`
	CatalogName      *string  `json:"catalog"`
	Price            *float64 `json:"price"`
}

type UpdateInput struct {
	Title *string `json:"title"`
}

func (i UpdateInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

func (i UpdateProductInput) Validate() error {
	if i.ProductName == nil && i.ManufacturerName == nil && i.CatalogName == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
