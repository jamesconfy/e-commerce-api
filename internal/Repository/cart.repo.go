package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type CartRepo interface {
	// GetProduct(productId string) (*productModels.GetProductRes, *errorModels.ServiceError)
	// GetUser(userId string) (*userModels.GetByIdRes, *errorModels.ServiceError)
	Add(req *models.CartItem) error
	GetCart(cartId string) (*models.Cart, error)
	// GetItems(cartId string) ([]*models.CartItem, error)
	// CheckProduct(productId, cartId string) error
	Get(productId, cartId string) (*models.CartItem, error)
	// Edit(req *cartModels.EditItemReq) error
	Delete(productId, cartId string) error
}

type cartSql struct {
	conn *sql.DB
}

func (c *cartSql) Add(req *models.CartItem) error {
	// SQL – Add item to cart on conflict update item quantity
	query := `INSERT INTO cart_item (cart_id, product_id, quantity, date_created, date_updated)
	VALUES ('%[1]v', '%[2]v', '%[3]v', CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP())
	ON DUPLICATE KEY UPDATE quantity = %[3]v, date_updated = CURRENT_TIMESTAMP()`

	stmt := fmt.Sprintf(query, req.CartId, req.ProductId, req.Quantity)

	_, err := c.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (c *cartSql) GetCart(cartId string) (*models.Cart, error) {
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

	query := `SELECT cart_id, product_id, quantity, date_created, date_updated FROM cart_item WHERE cart_id = '%[1]v'`
	stmt := fmt.Sprintf(query, cartId)

	rows, err := tx.Query(stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item models.CartItem

		err := rows.Scan(&item.CartId, &item.ProductId, &item.Quantity, &item.DateCreated, item.DateUpdated)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	for _, item := range items {
		var product models.Product

		query1 := `SELECT id, user_id, name, description, price, date_created, date_updated, image FROM products WHERE id = '%[1]v'`
		stmt1 := fmt.Sprintf(query1, item.ProductId)

		row := tx.QueryRow(stmt1)

		err := row.Scan(&product.Id, &product.UserId, &product.Name, &product.Description, &product.Price, &product.DateCreated, &product.DateUpdated, &product.Image)

		if err != nil {
			return nil, err
		}

		item.Product = &product
	}

	query3 := `SELECT date_created, date_updated FROM carts WHERE id = '%[1]v'`
	stmt3 := fmt.Sprintf(query3, cartId)

	row := tx.QueryRow(stmt3)
	if err := row.Scan(&cart.DateCreated, &cart.DateUpdated); err != nil {
		return nil, err
	}

	query4 := `SELECT ifnull(sum(p.price * ct.quantity), 0.00) AS total_price FROM cart_item ct JOIN products p ON ct product_id = p.id WHERE ct.cart_id = '%[1]v'`
	stmt4 := fmt.Sprintf(query4, cartId)

	row = tx.QueryRow(stmt4)
	if err := row.Scan(&cart.TotalPrice); err != nil {
		return nil, err
	}

	cart.Items = items

	return &cart, nil
}

func (c *cartSql) Get(productId, cartId string) (*models.CartItem, error) {
	var item models.CartItem

	query := `SELECT cart_id, product_id, quantity, date_created, date_updated FROM cart_item WHERE product_id = '%[1]v' AND cart_id = '%[2]v'`

	stmt := fmt.Sprintf(query, productId, cartId)

	err := c.conn.QueryRow(stmt).Scan(
		&item.CartId,
		&item.ProductId,
		&item.Quantity,
		&item.DateCreated,
		&item.DateUpdated,
	)

	if err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return &models.CartItem{}, err
			}

			return nil, err
		}
	}
	return &item, nil
}

func (c *cartSql) Delete(productId, cartId string) error {
	query := `DELETE FROM cart_item WHERE product_id = '%[1]v' AND cart_id = '%[2]v'`
	stmt := fmt.Sprintf(query, productId, cartId)
	_, err := c.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func NewCartRepo(conn *sql.DB) CartRepo {
	return &cartSql{conn: conn}
}
