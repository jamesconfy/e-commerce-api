package cartRepo

import (
	"database/sql"
	"e-commerce/internal/models/cartModels"
	"e-commerce/internal/models/errorModels"
	"e-commerce/internal/models/productModels"
	"e-commerce/internal/models/userModels"
	"errors"
	"fmt"
)

type CartRepo interface {
	GetProduct(productId string) (*productModels.GetProductRes, *errorModels.ServiceError)
	GetUser(userId string) (*userModels.GetByIdRes, *errorModels.ServiceError)
	AddToCart(req *cartModels.AddToCartReq) *errorModels.ServiceError
	CheckIfProductInCart(productId, cartId string) *errorModels.ServiceError
	GetItem(itemId string) (*cartModels.GetItemByIdRes, *errorModels.ServiceError)
	DeleteItem(itemId string) *errorModels.ServiceError
}

type mySql struct {
	conn *sql.DB
}

func (m *mySql) GetProduct(productId string) (*productModels.GetProductRes, *errorModels.ServiceError) {
	tx, err := m.conn.Begin()
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var product productModels.GetProductRes

	stmt := fmt.Sprintf(`SELECT product_id, user_id, name, description, price, date_updated, date_created, image FROM products WHERE product_id = '%s'`, productId)
	row := tx.QueryRow(stmt)

	err = row.Scan(&product.ProductId, &product.UserId, &product.Name,
		&product.Description, &product.Price, &product.DateCreated,
		&product.DateUpdated, &product.Image)

	if err != nil {
		if err == sql.ErrNoRows {
			return &productModels.GetProductRes{}, errorModels.NewCustomServiceError("Product not found", err)
		}

		return nil, errorModels.NewCustomServiceError("Error when getting product", err)
	}

	stmt = fmt.Sprintf(`SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%v'`, productId)
	row1 := tx.QueryRow(stmt)
	row1.Scan(&product.Rating)

	return &product, nil
}

func (m *mySql) GetUser(userId string) (*userModels.GetByIdRes, *errorModels.ServiceError) {
	tx, err := m.conn.Begin()
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var user userModels.GetByIdRes

	stmt := fmt.Sprintf(`
		SELECT user_id, email, password, first_name, last_name, phone_number, date_created, cart_id
		FROM users
		WHERE user_id = '%s'
	`, userId)

	err = tx.QueryRow(stmt).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.DateCreated,
		&user.CartId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &userModels.GetByIdRes{}, errorModels.NewCustomServiceError("User not found", err)
		}

		return nil, errorModels.NewCustomServiceError("Internal server error", err)
	}

	return &user, nil
}

func (m *mySql) AddToCart(req *cartModels.AddToCartReq) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	stmt := fmt.Sprintf(`INSERT INTO cart_item(
		cart_item_id, cart_id, product_id, quantity, price, date_created, date_updated
	) VALUES ('%v','%v','%v','%v','%v','%v','%v')`, req.CartItemId, req.CartId, req.ProductId, req.Quantity, req.Price, req.DateCreated, req.DateUpdated)

	_, err = tx.Exec(stmt)
	if err != nil {
		return errorModels.NewCustomServiceError("Error inserting data into cart", err)
	}

	return nil
}

func (m *mySql) CheckIfProductInCart(productId, cartId string) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	stmt := fmt.Sprintf(`SELECT * FROM cart_item WHERE product_id = '%v' AND cart_id = '%v'`, productId, cartId)
	rows, err := tx.Query(stmt)
	if err != nil {
		return errorModels.NewCustomServiceError("Internal server error", err)
	}

	if rows.Next() {
		return errorModels.NewCustomServiceError("Product already in cart", errors.New("rows returned"))
	}

	return nil
}

func (m *mySql) GetItem(itemId string) (*cartModels.GetItemByIdRes, *errorModels.ServiceError) {
	tx, err := m.conn.Begin()
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	var item cartModels.GetItemByIdRes
	stmt := fmt.Sprintf(`SELECT ct.cart_item_id, c.cart_id, ct.product_id, ct.quantity, ct.price, ct.date_created, ct.date_updated, c.user_id 
							FROM e_commerce_api.cart_item ct
							JOIN e_commerce_api.carts c
							ON ct.cart_id = c.cart_id
							WHERE ct.cart_item_id = '%v'`, itemId)

	err = tx.QueryRow(stmt).Scan(
		&item.CartItemId,
		&item.CartId,
		&item.ProductId,
		&item.Quantity,
		&item.Price,
		&item.DateCreated,
		&item.DateUpdated,
		&item.UserId,
	)

	if err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return &cartModels.GetItemByIdRes{}, errorModels.NewCustomServiceError("Item not found", err)
			}

			return nil, errorModels.NewCustomServiceError("Internal server error", err)
		}
	}
	return &item, nil
}

func (m *mySql) DeleteItem(itemId string) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	stmt := fmt.Sprintf(`DELETE FROM cart_item WHERE cart_item_id = '%v'`, itemId)
	_, err = tx.Exec(stmt)
	if err != nil {
		return errorModels.NewCustomServiceError("Internal server error", err)
	}

	return nil
}

func NewMySqlCartRepo(conn *sql.DB) CartRepo {
	return &mySql{conn: conn}
}
