package repo

import (
	"database/sql"
	"e-commerce/internal/models"
)

type CartRepo interface {
	Add(cart *models.Cart) (*models.Cart, error)
	Get(userId string) (*models.Cart, error)
	Clear(userId string) error
}

type cartSql struct {
	conn *sql.DB
}

func (c *cartSql) Add(cart *models.Cart) (*models.Cart, error) {
	query := `INSERT INTO carts(id, user_id, date_created) VALUES (?, ?, CURRENT_TIMESTAMP())`

	_, err := c.conn.Exec(query, cart.Id, cart.UserId)
	if err != nil {
		return nil, err
	}

	return c.Get(cart.UserId)
}

func (c *cartSql) Get(userId string) (*models.Cart, error) {
	var cart models.Cart

	tx, err := c.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	query := `SELECT id, user_id, date_created, date_updated FROM carts WHERE user_id = ?`

	row := tx.QueryRow(query, userId)
	err = row.Scan(&cart.Id, &cart.UserId, &cart.DateCreated, &cart.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (c *cartSql) Clear(userId string) error {
	cart, err := c.Get(userId)
	if err != nil {
		return err
	}

	query := `DELETE FROM cart_item WHERE cart_id = ?`

	_, err = c.conn.Exec(query, cart.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewCartRepo(conn *sql.DB) CartRepo {
	return &cartSql{conn: conn}
}
