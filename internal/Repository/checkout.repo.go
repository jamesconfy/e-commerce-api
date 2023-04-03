package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type CheckoutRepo interface {
}

type checkoutSql struct {
	conn *sql.DB
}

func (co *checkoutSql) Add(checkout *models.Checkout) (*models.Checkout, error) {
	// tx, err := co.conn.Begin()
	// if err != nil {
	// 	return nil, err
	// }

	// defer func() {
	// 	if err != nil {
	// 		tx.Rollback()
	// 	}
	// 	tx.Commit()
	// }()

	query := `INSERT INTO checkout(id, amount, quantity, user_id, product_id, price, payment_method, date_created, date_updated) VALUES('%v', %v, %v, '%v', '%v', %v, '%v', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP())`
	stmt := fmt.Sprintf(query, checkout.Id, checkout.Amount, checkout.Quantity, checkout.UserId, checkout.ProductId, checkout.Price(), checkout.PaymentMethod)

	_, err := co.conn.Exec(stmt)
	if err != nil {
		return nil, err
	}

	return co.Get(checkout.Id)
}

func (co *checkoutSql) Get(checkoutId string) (*models.Checkout, error) {
	var checkout models.Checkout

	query := `SELECT id, amount, quantity, user_id, product_id, status, payment_method, date_created, date_updated FROM checkout WHERE id = '%v'`
	stmt := fmt.Sprintf(query, checkoutId)

	row := co.conn.QueryRow(stmt)

	err := row.Scan(&checkout.Id, &checkout.Amount, &checkout.Quantity, &checkout.UserId, &checkout.ProductId, &checkout.Status, &checkout.PaymentMethod, &checkout.DateCreated, &checkout.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &checkout, nil
}

func NewCheckoutRepo(conn *sql.DB) CheckoutRepo {
	return &checkoutSql{conn: conn}
}
