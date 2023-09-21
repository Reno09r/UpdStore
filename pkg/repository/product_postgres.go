package repository

import (
	"errors"
	"fmt"
	"strings"

	store "github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product store.Product) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var productId int
	createProductQuery := fmt.Sprintf(`INSERT INTO %s (product_name, manufacturer_id, catalog_id)
	VALUES ($1, (SELECT manufacturer_id FROM %s WHERE manufacturer_name = $2), (SELECT catalog_id FROM %s WHERE catalog_name = $3))
	RETURNING product_id`, ProductsTable, ManufacturerTable, CatalogTable)

	row := tx.QueryRow(createProductQuery, product.ProductName, product.ManufacturerName, product.CatalogName)
	err = row.Scan(&productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (product_id, date_price_change, new_price) values ($1, NOW(), $2)", PriceTable)
	_, err = tx.Exec(createListItemsQuery, productId, product.Price)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return productId, tx.Commit()
}

func (r *ProductPostgres) GetAll() ([]store.Product, error) {
	var products []store.Product
	query := fmt.Sprintf(`
	SELECT p.product_id, p.product_name, m.manufacturer_name, c.catalog_name, pc.new_price
	FROM
		%s p
	JOIN
		manufacturers m ON p.manufacturer_id = m.manufacturer_id
	JOIN
		catalogs c ON p.catalog_id = c.catalog_id
	JOIN
		price_change pc ON p.product_id = pc.product_id
	WHERE
		pc.date_price_change = (
			SELECT MAX(date_price_change)
			FROM price_change
			WHERE product_id = p.product_id
		)
	ORDER BY
		pc.date_price_change DESC;
	`, ProductsTable)
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductPostgres) GetById(productId int) (store.Product, error) {
	var product store.Product
	query := fmt.Sprintf(`
	SELECT p.product_id, p.product_name, m.manufacturer_name, c.catalog_name, pc.new_price
	FROM
		%s p
	JOIN
		manufacturers m ON p.manufacturer_id = m.manufacturer_id
	JOIN
		catalogs c ON p.catalog_id = c.catalog_id
	JOIN
		price_change pc ON p.product_id = pc.product_id
	WHERE
		pc.date_price_change = (
			SELECT MAX(date_price_change)
			FROM price_change
			WHERE product_id = p.product_id
		)
		AND p.product_id = $1
	ORDER BY
		pc.date_price_change DESC
	LIMIT 1;
	`, ProductsTable)

	err := r.db.Get(&product, query, productId)
	return product, err
}

func (r *ProductPostgres) Delete(productId int) error {
	var product store.Product
	queryCheck := fmt.Sprintf("SELECT product_name FROM %s WHERE product_id = $1", ProductsTable)
	err := r.db.Get(&product, queryCheck, productId)
	if err != nil {
		return errors.New("deletion by non-existent product id")
	}
	var id int
	queryCheck = "SELECT purchase_id FROM purchases WHERE product_id = $1"
	err = r.db.Get(&id, queryCheck, productId)
	if err != nil {
		query := "DELETE FROM purchases WHERE purchase_id  = $1"
		_, err = r.db.Exec(query, productId)
		if err != nil {
			return err
		}
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE product_id  = $1", PriceTable)
	_, err = r.db.Exec(query, productId)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE product_id = $1", ProductsTable)
	_, err = r.db.Exec(query, productId)
	return err
}

func (r *ProductPostgres) Update(productId int, input store.UpdateProductInput) error {
	var product store.Product
	queryCheck := fmt.Sprintf("SELECT product_name FROM %s WHERE product_id = $1", ProductsTable)
	err := r.db.Get(&product, queryCheck, productId)
	if err != nil {
		return errors.New("Update by non-existent productId")
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.ProductName != nil {
		setValues = append(setValues, fmt.Sprintf("product_name=$%d", argId))
		args = append(args, *input.ProductName)
		argId++
	}
	var catalog store.Catalog
	queryCheck = fmt.Sprintf("SELECT catalog_id FROM %s WHERE catalog_name = $1", CatalogTable)
	err = r.db.Get(&catalog, queryCheck, input.CatalogName)
	if err != nil {
		return errors.New("Update by non-existent catalogId")
	}
	if input.CatalogName != nil {
		setValues = append(setValues, fmt.Sprintf("catalog_id=$%d", argId))
		args = append(args, catalog.Id)
		argId++
	}

	var manufacturer store.Manufacturer
	queryCheck = fmt.Sprintf("SELECT manufacturer_id FROM %s WHERE manufacturer_name = $1", ManufacturerTable)
	err = r.db.Get(&manufacturer, queryCheck, input.ManufacturerName)
	if err != nil {
		return errors.New("Update by non-existent manufacturerId")
	}
	if input.ManufacturerName != nil {
		setValues = append(setValues, fmt.Sprintf("manufacturer_id=$%d", argId))
		args = append(args, manufacturer.Id)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE product_id=$%d", ProductsTable, setQuery, argId)
	args = append(args, productId)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec(query, args...)
	if err != nil {
		return err
	}

	if input.Price != nil {
		insertPriceQuery := fmt.Sprintf("INSERT INTO %s (product_id, new_price, date_price_change) VALUES ($1, $2, NOW())", PriceTable)
		_, err = tx.Exec(insertPriceQuery, productId, *input.Price)
		if err != nil {
			return err
		}
		query := "UPDATE purchase_items SET product_price = $1 WHERE product_id =$2"
		_, err = tx.Exec(query, *input.Price, productId)
		if err != nil {
			return err
		}

	}

	return err
}

func (r *ProductPostgres) GetAllByManufacturer(manufacturerID int) ([]store.Product, error) {
	var products []store.Product
	query := `
        SELECT p.product_id, p.product_name, m.manufacturer_name, c.catalog_name, pc.new_price
        FROM products p
        JOIN manufacturers m ON p.manufacturer_id = m.manufacturer_id
        JOIN catalogs c ON p.catalog_id = c.catalog_id
        JOIN price_change pc ON p.product_id = pc.product_id
        WHERE p.manufacturer_id = $1
          AND pc.date_price_change = (
              SELECT MAX(date_price_change)
              FROM price_change
              WHERE product_id = p.product_id
          )
        ORDER BY pc.date_price_change DESC;
    `

	err := r.db.Select(&products, query, manufacturerID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductPostgres) GetAllByCatalog(catalogID int) ([]store.Product, error) {
	var products []store.Product

	query := `
        SELECT p.product_id, p.product_name, m.manufacturer_name, c.catalog_name, pc.new_price
        FROM products p
        JOIN manufacturers m ON p.manufacturer_id = m.manufacturer_id
        JOIN catalogs c ON p.catalog_id = c.catalog_id
        JOIN price_change pc ON p.product_id = pc.product_id
        WHERE p.catalog_id = $1
          AND pc.date_price_change = (
              SELECT MAX(date_price_change)
              FROM price_change
              WHERE product_id = p.product_id
          )
        ORDER BY pc.date_price_change DESC;
    `

	err := r.db.Select(&products, query, catalogID)
	if err != nil {
		return nil, err
	}

	return products, nil
}
