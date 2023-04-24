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

func (c *cartSql) Add(cart *models.Cart) (cat *models.Cart, err error) {
	cat = new(models.Cart)

	query := `INSERT INTO carts(user_id) VALUES ($1) RETURNING id, user_id, date_created, date_updated`

	err = c.conn.QueryRow(query, cart.UserId).Scan(&cat.Id, &cat.UserId, &cat.DateCreated, &cat.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (c *cartSql) Get(userId string) (cat *models.Cart, err error) {
	cat = new(models.Cart)

	query := `SELECT id, user_id, date_created, date_updated FROM carts WHERE user_id = $1`

	row := c.conn.QueryRow(query, userId)
	err = row.Scan(&cat.Id, &cat.UserId, &cat.DateCreated, &cat.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (c *cartSql) Clear(userId string) (err error) {
	cart, err := c.Get(userId)
	if err != nil {
		return err
	}

	query := `DELETE FROM cart_item WHERE cart_id = $1 `

	_, err = c.conn.Exec(query, cart.Id)
	if err != nil {
		return
	}

	return
}

func NewCartRepo(conn *sql.DB) CartRepo {
	return &cartSql{conn: conn}
}
