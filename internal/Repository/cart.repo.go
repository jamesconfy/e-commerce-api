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
	// CheckProduct(productId, cartId string) error
	GetId(productId, cartId string) (*models.CartItem, error)
	// Edit(req *cartModels.EditItemReq) error
	Delete(productId, cartId string) error
}

type cartSql struct {
	conn *sql.DB
}

// func (m *cartSql) GetProduct(productId string) (*productModels.GetProductRes, *errorModels.ServiceError) {
// 	tx, err := m.conn.Begin()
// 	if err != nil {
// 		return nil, errorModels.NewCustomServiceError("Error when starting transaction", err)
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		} else {
// 			tx.Commit()
// 		}
// 	}()

// 	var product productModels.GetProductRes

// 	stmt := fmt.Sprintf(`SELECT product_id, user_id, name, description, price, date_updated, date_created, image FROM products WHERE product_id = '%s'`, productId)
// 	row := tx.QueryRow(stmt)

// 	err = row.Scan(&product.ProductId, &product.UserId, &product.Name,
// 		&product.Description, &product.Price, &product.DateCreated,
// 		&product.DateUpdated, &product.Image)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return &productModels.GetProductRes{}, errorModels.NewCustomServiceError("Product not found", err)
// 		}

// 		return nil, errorModels.NewCustomServiceError("Error when getting product", err)
// 	}

// 	stmt = fmt.Sprintf(`SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%v'`, productId)
// 	row1 := tx.QueryRow(stmt)
// 	row1.Scan(&product.Rating)

// 	return &product, nil
// }

// func (m *cartSql) GetUser(userId string) (*userModels.GetByIdRes, *errorModels.ServiceError) {
// 	tx, err := m.conn.Begin()
// 	if err != nil {
// 		return nil, errorModels.NewCustomServiceError("Error when starting transaction", err)
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		} else {
// 			tx.Commit()
// 		}
// 	}()

// 	var user userModels.GetByIdRes

// 	stmt := fmt.Sprintf(`
// 		SELECT user_id, email, password, first_name, last_name, phone_number, date_created, cart_id
// 		FROM users
// 		WHERE user_id = '%s'
// 	`, userId)

// 	err = tx.QueryRow(stmt).Scan(
// 		&user.UserId,
// 		&user.Email,
// 		&user.Password,
// 		&user.FirstName,
// 		&user.LastName,
// 		&user.PhoneNumber,
// 		&user.DateCreated,
// 		&user.CartId,
// 	)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return &userModels.GetByIdRes{}, errorModels.NewCustomServiceError("User not found", err)
// 		}

// 		return nil, errorModels.NewCustomServiceError("Internal server error", err)
// 	}

// 	return &user, nil
// }

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

func (c *cartSql) GetId(productId, cartId string) (*models.CartItem, error) {
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

// func (m *cartSql) Edit(cartItem *cartModels.EditItemReq) error {
// 	tx, err := m.conn.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}

// 		tx.Commit()
// 	}()

// 	stmt := fmt.Sprintf(`UPDATE cart_item SET quantity = %v, price = %v, date_updated = '%v' WHERE cart_item_id = '%v'`, cartItem.Quantity, cartItem.Price, cartItem.DateUpdated, cartItem.CartItemId)
// 	_, err = tx.Exec(stmt)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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
