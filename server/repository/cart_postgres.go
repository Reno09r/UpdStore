package repository

import (
	"errors"
	"fmt"

	store "github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
)

type StoreCartPostgres struct {
	db *sqlx.DB
}

func NewStoreCartPostgres(db *sqlx.DB) *StoreCartPostgres {
	return &StoreCartPostgres{db: db}
}

func (r *StoreCartPostgres) Insert(input store.Cart, customerId int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	if purchaseId, err := r.CheckProductInCart(customerId, input.ID); purchaseId == 0 && err != nil {

		insertToCartQuery := fmt.Sprintf(` 
		INSERT INTO purchases (purchase_date, customer_id, store_id)
		VALUES (NOW(), $1, $2)
		RETURNING purchase_id;
		`)

		row := tx.QueryRow(insertToCartQuery, customerId, 1)
		if err := row.Scan(&id); err != nil {
			tx.Rollback()
			return 0, err
		}
		if input.Quantity == 0{
			input.Quantity++
		}
		insertToCartQuery = fmt.Sprintf(`
		INSERT INTO purchase_items (purchase_id, product_id, product_count, product_price) 
		VALUES (
			$1, 
			$2, 
			$3,
			(
				SELECT new_price 
				FROM price_change 
				WHERE date_price_change = (
					SELECT MAX(date_price_change) 
					FROM price_change 
					WHERE product_id = $2
				) 
				ORDER BY date_price_change DESC
			)
		);
		`)
		_, err = tx.Exec(insertToCartQuery, id, input.ID, input.Quantity)

		if err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {
		if input.Quantity == 0{
			input.Quantity++
		}
		updateQuantity := fmt.Sprintf(`UPDATE purchase_items 
		SET product_count = product_count + $1 
		WHERE purchase_id = $2
		AND product_id = $3`)
		r.db.Exec(updateQuantity, input.Quantity, purchaseId, input.ID)
		id = purchaseId
	}
	return id, tx.Commit()
}

func (r *StoreCartPostgres) CheckProductInCart(customerId, productId int) (int, error) {
	var id int
	query := fmt.Sprintf(`SELECT pi.purchase_id
	FROM purchase_items pi
	JOIN purchases ps ON pi.purchase_id = ps.purchase_id
	WHERE ps.customer_id = $1
	  AND pi.product_id = $2
	ORDER BY ps.purchase_date DESC 
	LIMIT 1;`)
	err := r.db.Get(&id, query, customerId, productId)
	return id, err
}

func (r *StoreCartPostgres) Get(customerId int) ([]store.Cart, error) {
	var cart []store.Cart
	query := fmt.Sprintf(`SELECT p.product_name, m.manufacturer_name, c.catalog_name, p.product_id, pc.new_price, pi.product_count, pi.purchase_id
	FROM products p
	JOIN manufacturers m ON p.manufacturer_id = m.manufacturer_id
	JOIN catalogs c ON p.catalog_id = c.catalog_id
	JOIN price_change pc ON p.product_id = pc.product_id
	JOIN purchase_items pi ON p.product_id = pi.product_id
	JOIN purchases ps ON pi.purchase_id = ps.purchase_id
	JOIN customers cs ON ps.customer_id = cs.customer_id  
	WHERE pc.date_price_change = (
		SELECT MAX(date_price_change)
		FROM price_change
		WHERE product_id = p.product_id
	)
	AND cs.customer_id = $1
	ORDER BY pc.date_price_change DESC;
	`)
	err := r.db.Select(&cart, query, customerId)
	return cart, err
}

func (r *StoreCartPostgres) Delete(productId, customerId int) error {
	queryCheck := fmt.Sprintf("SELECT purchase_id FROM purchase_items WHERE product_id = $1 AND purchase_id IN (SELECT purchase_id FROM purchases WHERE customer_id = $2)")
	var purchaseID int
	err := r.db.Get(&purchaseID, queryCheck, productId, customerId)
	if err != nil {
		return errors.New("Deletion by non-existent purchase ID")
	}

	query := fmt.Sprintf("DELETE FROM purchase_items WHERE product_id = $1 AND purchase_id = $2")
	_, err = r.db.Exec(query, productId, purchaseID)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("DELETE FROM purchases WHERE purchase_id = $1 AND NOT EXISTS (SELECT 1 FROM purchase_items WHERE purchase_id = $1)")
	_, err = r.db.Exec(query, purchaseID)

	return err
}
