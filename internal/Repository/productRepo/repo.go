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
	GetProductById(productId string) (*productModels.GetProductById, error)
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

func (m *mySql) GetProductById(productId string) (*productModels.GetProductById, error) {
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

	var product productModels.GetProductById

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
