package repository

import (
	"errors"
	"fmt"
	"time"

	store "github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
)

type StoreBuyPostgres struct {
	db *sqlx.DB
}

func NewStoreBuyPostgres(db *sqlx.DB) *StoreBuyPostgres {
	return &StoreBuyPostgres{db: db}
}

func (r *StoreBuyPostgres) Confirm(input store.UserCardInput, userId int) error {
	var ids []int
	queryCheck := "SELECT purchase_id FROM purchases WHERE user_id = $1"
	rows, err := r.db.Query(queryCheck, userId)
	if err != nil {
		return errors.New("no products in cart")
	}
	defer rows.Close()
	for rows.Next() {
		var p int
		err = rows.Scan(&p)
		if err != nil {
			return err
		}
		ids = append(ids, p)
	}
	for _, purchaseID := range ids {
		query := fmt.Sprintf(`INSERT INTO %s (user_id, product_id, product_count, full_price, buy_date)
		VALUES ($1, (SELECT product_id FROM purchase_items WHERE purchase_id = $2), 
				(SELECT product_count FROM purchase_items WHERE purchase_id = $2),
				(SELECT product_count * product_price FROM purchase_items WHERE purchase_id = $2),
				NOW());`, BuysTable)
		_, err = r.db.Exec(query, userId, purchaseID)
		if err != nil {
			return err
		}
		query = "DELETE FROM purchase_items WHERE purchase_id = $1"
		_, err = r.db.Exec(query, purchaseID)
		if err != nil {
			return err
		}

		query = "DELETE FROM purchases WHERE purchase_id = $1 AND NOT EXISTS (SELECT 1 FROM purchase_items WHERE purchase_id = $1)"
		_, err = r.db.Exec(query, purchaseID)
		if err != nil {
			return err
		}
	}

	return err
}

func (r *StoreBuyPostgres) BoughtProducts(userId int) ([]store.BoughtProducts, error) {
	var buyedProducts []store.BoughtProducts
	query := fmt.Sprintf(`SELECT
    p.product_name,
    m.manufacturer_name,
    c.catalog_name,
    p.product_id,
    pc.new_price,
    b.full_price,
    b.product_count,
    b.buy_id,
    b.buy_date
FROM
    %s p
JOIN
%s m ON p.manufacturer_id = m.manufacturer_id
JOIN
%s c ON p.catalog_id = c.catalog_id
JOIN
%s b ON p.product_id = b.product_id
JOIN
%s u ON b.user_id = u.user_id
LEFT JOIN LATERAL (
    SELECT
        pc.product_id,
        pc.new_price,
        pc.date_price_change
    FROM
	%s pc
    WHERE
        pc.product_id = p.product_id
        AND pc.date_price_change <= b.buy_date
    ORDER BY
        pc.date_price_change DESC
    LIMIT 1
) pc ON true
WHERE
    u.user_id = $1;
	`, ProductsTable, ManufacturerTable, CatalogTable, BuysTable, UsersTable, PriceTable)
	err := r.db.Select(&buyedProducts, query, userId)
	for i := range buyedProducts {
		buyedProducts[i].TimeBuy, err = time.Parse("2006-01-02 15:04:05", buyedProducts[i].TimeBuy.Local().Format("2006-01-02 15:04:05"))
		if err != nil {
			return nil, err
		}
	}

	return buyedProducts, err
}
