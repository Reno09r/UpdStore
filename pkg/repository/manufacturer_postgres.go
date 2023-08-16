package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type StoreManufacturerPostgres struct {
	db *sqlx.DB
}

func NewStoreManufacturerPostgres(db *sqlx.DB) *StoreManufacturerPostgres {
	return &StoreManufacturerPostgres{db: db}
}

func (r *StoreManufacturerPostgres) Create(manufacturer store.Manufacturer) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createManufacturerQuery := fmt.Sprintf("INSERT INTO %s (manufacturer_name) VALUES ($1) RETURNING manufacturer_id", ManufacturerTable)
	row := tx.QueryRow(createManufacturerQuery, manufacturer.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *StoreManufacturerPostgres) GetAll() ([]store.Manufacturer, error) {
	var manufacturers []store.Manufacturer
	query := fmt.Sprintf("SELECT * FROM %s", ManufacturerTable)
	err := r.db.Select(&manufacturers, query)
	return manufacturers, err
}

func (r *StoreManufacturerPostgres) GetById(manufacturerId int) (store.Manufacturer, error) {
	var manufacturer store.Manufacturer

	query := fmt.Sprintf("SELECT * FROM %s WHERE manufacturer_id = $1", ManufacturerTable)
	err := r.db.Get(&manufacturer, query, manufacturerId)

	return manufacturer, err
}

func (r *StoreManufacturerPostgres) Delete(manufacturerId int) error {
	var manufacturer store.Manufacturer
	queryCheck := fmt.Sprintf("SELECT * FROM %s WHERE manufacturer_id = $1", ManufacturerTable)
	err := r.db.Get(&manufacturer, queryCheck, manufacturerId)
	if err != nil {
		return errors.New("Delete by non-existent ManufacturerId")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE manufacturer_id = $1", ManufacturerTable)
	_, err = r.db.Exec(query, manufacturerId)
	return err
}

func (r *StoreManufacturerPostgres) Update(manufacturerId int, input store.UpdateInput) error {
	var manufacturer store.Manufacturer
	queryCheck := fmt.Sprintf("SELECT * FROM %s WHERE manufacturer_id = $1", ManufacturerTable)
	err := r.db.Get(&manufacturer, queryCheck, manufacturerId)
	if err != nil {
		return errors.New("Update by non-existent ManufacturerId")
	}
	setValues := make([]string, 0)
	var arg interface{}
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("manufacturer_name=$%d", argId))
		arg = *input.Title
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ct SET %s WHERE ct.manufacturer_id =$%d",
		ManufacturerTable, setQuery, argId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", arg)
	_, err = r.db.Exec(query, arg, manufacturerId)
	return err
}
