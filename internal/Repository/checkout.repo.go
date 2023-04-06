package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type CheckoutRepo interface {
	Add(checkout *models.Checkout) (*models.Checkout, error)
}

type checkoutSql struct {
	conn *sql.DB
}

func (co *checkoutSql) Add(checkout *models.Checkout) (*models.Checkout, error) {
	query := `INSERT INTO checkout(id, quantity, cart_id, product_id, payment_method, date_created, date_updated, status) VALUES('%v', %v, '%v', '%v', '%v', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), 'ACTIVE')`
	stmt := fmt.Sprintf(query, checkout.Id, checkout.Quantity, checkout.CartId, checkout.ProductId, checkout.PaymentMethod)

	_, err := co.conn.Exec(stmt)
	if err != nil {
		return nil, err
	}

	return co.Get(checkout.Id)
}

func (co *checkoutSql) Get(checkoutId string) (*models.Checkout, error) {
	var checkout models.Checkout

	query := `SELECT id, quantity, cart_id, product_id, status, payment_method, date_created, date_updated FROM checkout WHERE id = '%v'`
	stmt := fmt.Sprintf(query, checkoutId)

	row := co.conn.QueryRow(stmt)

	err := row.Scan(&checkout.Id, &checkout.Quantity, &checkout.CartId, &checkout.ProductId, &checkout.Status, &checkout.PaymentMethod, &checkout.DateCreated, &checkout.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &checkout, nil
}

func NewCheckoutRepo(conn *sql.DB) CheckoutRepo {
	return &checkoutSql{conn: conn}
}
