package productRepo

import (
	"context"
	"database/sql"
	"e-commerce/internal/models/productModels"
	"fmt"
	"time"
)

type ProductRepo interface {
	AddProduct(req *productModels.AddProductReq) error
	GetProducts(page int) ([]*productModels.GetProduct, error)
	GetProduct(productId string) (*productModels.GetProduct, error)
	DeleteProduct(productId string) (*productModels.DeleteProduct, error)
	AddRating(req *productModels.AddRatingsReq) error
	VerifyUserRatings(userId, productId string) error
}

type mySql struct {
	conn *sql.DB
}

func (m *mySql) AddProduct(req *productModels.AddProductReq) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	tx, err := m.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt := fmt.Sprintf(`INSERT INTO products(
		product_id, user_id, name, description, price, date_created, date_updated, image) 
		VALUES (
			'%v', '%v', '%v', '%v', %v, '%v', '%v', '%v'
		)`, req.ProductId, req.UserId, req.Name, req.Description, req.Price, req.DateCreated, req.DateUpdated, req.Image)

	_, err = tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}

func (m *mySql) GetProducts(page int) ([]*productModels.GetProduct, error) {
	tx, err := m.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var products []*productModels.GetProduct

	limit := 20
	offset := limit * (page - 1)
	stmt := fmt.Sprintf(`SELECT product_id, user_id, name, description, price, date_updated, date_created, image 
			FROM products
			ORDER BY product_id
			LIMIT %v
			OFFSET %v`, limit, offset)
	rows, err := tx.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product productModels.GetProduct

		if err = rows.Scan(&product.ProductId, &product.UserId, &product.Name,
			&product.Description, &product.Price, &product.DateCreated,
			&product.DateUpdated, &product.Image); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	for _, product := range products {
		stmt = fmt.Sprintf(`SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%v'`, product.ProductId)
		row1 := tx.QueryRow(stmt)
		row1.Scan(&product.Rating)
	}

	return products, nil
}

func (m *mySql) GetProduct(productId string) (*productModels.GetProduct, error) {
	tx, err := m.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var product productModels.GetProduct

	stmt := fmt.Sprintf(`SELECT product_id, user_id, name, description, price, date_updated, date_created, image FROM products WHERE product_id = '%s'`, productId)
	row := tx.QueryRow(stmt)

	if err = row.Scan(&product.ProductId, &product.UserId, &product.Name,
		&product.Description, &product.Price, &product.DateCreated,
		&product.DateUpdated, &product.Image); err != nil {
		return nil, err
	}

	stmt = fmt.Sprintf(`SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%v'`, productId)
	row1 := tx.QueryRow(stmt)
	row1.Scan(&product.Rating)

	return &product, nil
}

func (m *mySql) DeleteProduct(productId string) (*productModels.DeleteProduct, error) {
	tx, err := m.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var product productModels.DeleteProduct

	stmt := fmt.Sprintf(`SELECT product_id, user_id, name, description, price, date_updated, date_created, image 
			FROM products
			WHERE product_id = '%v'
			`, productId)
	row := tx.QueryRow(stmt)

	if err = row.Scan(&product.ProductId, &product.UserId, &product.Name,
		&product.Description, &product.Price, &product.DateCreated,
		&product.DateUpdated, &product.Image); err != nil {
		return nil, err
	}

	stmt1 := fmt.Sprintf(`DELETE FROM products WHERE product_id = '%v'`, productId)
	_, err = tx.Exec(stmt1)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (m *mySql) AddRating(req *productModels.AddRatingsReq) error {
	tx, err := m.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt := fmt.Sprintf(`INSERT INTO ratings(
		rating_id, rating, product_id, user_id, date_created, date_updated) VALUES(
			'%v', %v, '%v', '%v', '%v', '%v'
		)`, req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated, req.DateUpdated)

	_, err = tx.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (m *mySql) VerifyUserRatings(userId, productId string) error {
	tx, err := m.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt := fmt.Sprintf(`SELECT * FROM ratings WHERE user_id = '%s' AND product_id = '%s'`, userId, productId)
	rows, err := tx.Query(stmt)

	if rows.Next() {
		return fmt.Errorf("product already rated, you can edit it though")
	}
	defer rows.Close()

	return nil
}

func NewMySqlUserRepo(conn *sql.DB) ProductRepo {
	return &mySql{conn: conn}
}
