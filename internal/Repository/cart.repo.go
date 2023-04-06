package repo

import (
	"database/sql"
	"e-commerce/internal/models"
)

type CartRepo interface {
	// Cart
	CreateCart(cart *models.Cart) (*models.Cart, error)
	GetCart(userId string) (*models.Cart, error)
	ClearCart(userId string) error

	// Cart Item
	AddItem(req *models.CartItem, userId string) (*models.CartItem, error)
	GetItem(productId, userId string) (*models.CartItem, error)
	DeleteItem(productId, userId string) error
}

type cartSql struct {
	conn *sql.DB
}

func (c *cartSql) CreateCart(cart *models.Cart) (*models.Cart, error) {
	query := `INSERT INTO carts(id, user_id, date_created) VALUES (?, ?, CURRENT_TIMESTAMP())`

	_, err := c.conn.Exec(query, cart.Id, cart.UserId)
	if err != nil {
		return nil, err
	}

	return c.GetCart(cart.UserId)
}

func (c *cartSql) GetCart(userId string) (*models.Cart, error) {
	var cart models.Cart
	var items []*models.CartItem

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

	query = `SELECT ct.cart_id, ct.product_id, ct.quantity, ct.date_created, ct.date_updated, p.id "product_id", p.user_id, p.name, p.description, p.price, p.date_created "product_date_created", p.date_updated "product_date_updated", p.image FROM cart_item ct LEFT JOIN products p ON p.id=ct.product_id WHERE ct.cart_id = ?;`

	rows, err := tx.Query(query, cart.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item models.CartItem
		item.Product = &models.Product{}

		err := rows.Scan(&item.CartId, &item.ProductId, &item.Quantity, &item.DateCreated, &item.DateUpdated, &item.Product.Id, &item.Product.UserId, &item.Product.Name, &item.Product.Description, &item.Product.Price, &item.Product.DateCreated, &item.Product.DateUpdated, &item.Product.Image)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	cart.Items = items
	cart.TotalPrice = cart.Total()
	return &cart, nil
}

func (c *cartSql) ClearCart(userId string) error {
	cart, err := c.GetCart(userId)
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

func (c *cartSql) AddItem(item *models.CartItem, userId string) (*models.CartItem, error) {
	// SQL – Add item to cart on conflict update item quantity
	query := `INSERT INTO cart_item (cart_id, product_id, quantity)
				VALUES (?, ?, ?)
				ON DUPLICATE KEY UPDATE quantity = ?, date_updated = CURRENT_TIMESTAMP()`

	_, err := c.conn.Exec(query, item.CartId, item.ProductId, item.Quantity, item.Quantity)
	if err != nil {
		return nil, err
	}

	return c.GetItem(item.ProductId, userId)
}

func (c *cartSql) GetItem(productId, userId string) (*models.CartItem, error) {
	var item models.CartItem

	cart, err := c.GetCart(userId)
	if err != nil {
		return nil, err
	}

	query := `SELECT cart_id, product_id, quantity, date_created, date_updated FROM cart_item WHERE product_id = ? AND cart_id = ?`

	err = c.conn.QueryRow(query, productId, cart.Id).Scan(
		&item.CartId,
		&item.ProductId,
		&item.Quantity,
		&item.DateCreated,
		&item.DateUpdated,
	)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *cartSql) DeleteItem(productId, userId string) error {
	cart, err := c.GetCart(userId)
	if err != nil {
		return err
	}

	query := `DELETE FROM cart_item WHERE product_id = ? AND cart_id = ?`

	_, err = c.conn.Exec(query, productId, cart.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewCartRepo(conn *sql.DB) CartRepo {
	return &cartSql{conn: conn}
}
