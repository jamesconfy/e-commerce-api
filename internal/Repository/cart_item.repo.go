package repo

import (
	"database/sql"
	"e-commerce/internal/models"
)

type CartItemRepo interface {
	Add(cart *models.Cart, item *models.Item) (*models.Item, error)
	Get(cart *models.Cart, productId string) (*models.Item, error)
	GetItems(cart *models.Cart) (*models.CartItem, error)
	Delete(cart *models.Cart, productId string) error
}

type itemSql struct {
	conn *sql.DB
}

func (c *itemSql) Add(cart *models.Cart, item *models.Item) (itm *models.Item, err error) {
	// SQL – Add item to cart on conflict update item quantity
	itm = new(models.Item)

	query := `INSERT INTO cart_item(cart_id, product_id, quantity) VALUES ($1, $2, $3) ON CONFLICT (cart_id, product_id)
	DO UPDATE SET quantity = $3, date_updated = CURRENT_TIMESTAMP RETURNING product_id, quantity, date_created, date_updated`

	err = c.conn.QueryRow(query, cart.Id, item.ProductId, item.Quantity).Scan(&itm.ProductId, &itm.Quantity, &itm.DateCreated, &itm.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (c *itemSql) Get(cart *models.Cart, productId string) (*models.Item, error) {
	var item models.Item

	query := `SELECT product_id, quantity, date_created, date_updated FROM cart_item WHERE product_id = $1 AND cart_id = $2`

	err := c.conn.QueryRow(query, productId, cart.Id).Scan(&item.ProductId, &item.Quantity, &item.DateCreated, &item.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *itemSql) GetItems(cart *models.Cart) (*models.CartItem, error) {
	var items []*models.Item
	var cartItem models.CartItem

	query := `SELECT ct.product_id, ct.quantity, ct.date_created, ct.date_updated, p.id "product_id", p.user_id, p.name, p.description, p.price, p.date_created "product_date_created", p.date_updated "product_date_updated", p.image FROM cart_item ct LEFT JOIN products p ON p.id=ct.product_id WHERE ct.cart_id = $1;`

	rows, err := i.conn.Query(query, cart.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item models.Item
		item.Product = &models.Product{}

		err := rows.Scan(&item.ProductId, &item.Quantity, &item.DateCreated, &item.DateUpdated, &item.Product.Id, &item.Product.UserId, &item.Product.Name, &item.Product.Description, &item.Product.Price, &item.Product.DateCreated, &item.Product.DateUpdated, &item.Product.Image)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	cartItem.Cart = cart
	cartItem.Items = items
	cartItem.TotalPrice = cartItem.Total()

	return &cartItem, nil
}

func (c *itemSql) Delete(cart *models.Cart, productId string) error {
	query := `DELETE FROM cart_item WHERE product_id = $1 AND cart_id = $2`

	_, err := c.conn.Exec(query, productId, cart.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewCartItemRepo(conn *sql.DB) CartItemRepo {
	return &itemSql{conn: conn}
}
