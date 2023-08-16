package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type StoreCatalogPostgres struct {
	db *sqlx.DB
}

func NewStoreCatalogPostgres(db *sqlx.DB) *StoreCatalogPostgres {
	return &StoreCatalogPostgres{db: db}
}

func (r *StoreCatalogPostgres) Create(catalog store.Catalog) (int, error){
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createcatalogQuery := fmt.Sprintf("INSERT INTO %s (catalog_name) VALUES ($1) RETURNING catalog_id", CatalogTable)
	row := tx.QueryRow(createcatalogQuery, catalog.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *StoreCatalogPostgres)GetAll() ([]store.Catalog, error){
	var catalogs []store.Catalog
	query := fmt.Sprintf("SELECT * FROM %s", CatalogTable)
	err := r.db.Select(&catalogs, query)
	return catalogs, err
}

func (r *StoreCatalogPostgres)GetById(CatalogId int) (store.Catalog, error){
	var catalog store.Catalog

	query := fmt.Sprintf("SELECT * FROM %s WHERE catalog_id = $1", CatalogTable)
	err := r.db.Get(&catalog, query, CatalogId)

	return catalog, err
}

func (r *StoreCatalogPostgres)Delete(CatalogId int) error{
	var catalog store.Catalog
	queryCheck := fmt.Sprintf("SELECT * FROM %s WHERE catalog_id = $1", CatalogTable)
	err := r.db.Get(&catalog, queryCheck, CatalogId)
	if err != nil {
		return errors.New("Delete by non-existent CatalogId")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE catalog_id = $1", CatalogTable)
	_, err = r.db.Exec(query, CatalogId)
	return err
}

func (r *StoreCatalogPostgres)Update(CatalogId int, input store.UpdateInput) error{
	var manufacturer store.Catalog
	queryCheck := fmt.Sprintf("SELECT * FROM %s WHERE catalog_id = $1", CatalogTable)
	err := r.db.Get(&manufacturer, queryCheck, CatalogId)
	if err != nil {
		return errors.New("Update by non-existent CatalogId")
	}
	setValues := make([]string, 0)
	var arg interface{}
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("catalog_name=$%d", argId))
		arg = *input.Title
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ct SET %s WHERE ct.catalog_id =$%d",
		CatalogTable, setQuery, argId)
	
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", arg)
	_, err = r.db.Exec(query, arg, CatalogId)
	return err
}