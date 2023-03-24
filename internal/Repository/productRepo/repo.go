package productRepo

import (
	"database/sql"
	"e-commerce/internal/models/errorModels"
	"e-commerce/internal/models/productModels"
	"fmt"
)

type ProductRepo interface {
	AddProduct(req *productModels.AddProductReq) *errorModels.ServiceError
	GetProducts(page int) ([]*productModels.GetProductRes, *errorModels.ServiceError)
	GetProduct(productId string) (*productModels.GetProductRes, *errorModels.ServiceError)
	EditProduct(req *productModels.EditProductReq) *errorModels.ServiceError
	DeleteProduct(productId string) *errorModels.ServiceError
	AddRating(req *productModels.AddRatingsReq) *errorModels.ServiceError
	VerifyUserRatings(userId, productId string) *errorModels.ServiceError
}

type mySql struct {
	conn *sql.DB
}

func (m *mySql) AddProduct(req *productModels.AddProductReq) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
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

	_, err = tx.Exec(stmt)
	if err != nil {
		return errorModels.NewCustomServiceError("Internal server error", err)
	}

	return nil
}

func (m *mySql) GetProducts(page int) ([]*productModels.GetProductRes, *errorModels.ServiceError) {
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

	var products []*productModels.GetProductRes

	limit := 20
	offset := limit * (page - 1)
	stmt := fmt.Sprintf(`SELECT product_id, user_id, name, description, price, date_updated, date_created, image 
			FROM products
			ORDER BY product_id
			LIMIT %v
			OFFSET %v`, limit, offset)
	rows, err := tx.Query(stmt)
	if err != nil {
		return nil, errorModels.NewCustomServiceError("Error when fetching products from database", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product productModels.GetProductRes

		if err = rows.Scan(&product.ProductId, &product.UserId, &product.Name,
			&product.Description, &product.Price, &product.DateCreated,
			&product.DateUpdated, &product.Image); err != nil {
			return nil, errorModels.NewCustomServiceError("Error when saving product to structure", err)
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

func (m *mySql) EditProduct(req *productModels.EditProductReq) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt := fmt.Sprintf(`UPDATE products 
						SET name = '%v', description = '%v', price = %v, date_updated = '%v', image = '%v'
						WHERE user_id = '%v' AND product_id = '%v'`, req.Name, req.Description, req.Price, req.DateUpdated, req.Image, req.UserId, req.ProductId)

	_, err = tx.Exec(stmt)
	if err != nil {
		return errorModels.NewCustomServiceError("Error when updating product", err)
	}

	return nil
}

func (m *mySql) DeleteProduct(productId string) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt1 := fmt.Sprintf(`DELETE FROM products WHERE product_id = '%v'`, productId)
	_, err = tx.Exec(stmt1)
	if err != nil {
		return errorModels.NewCustomServiceError("Error when deleting product", err)
	}

	return nil
}

func (m *mySql) AddRating(req *productModels.AddRatingsReq) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
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
		return errorModels.NewCustomServiceError("error when saving rating", err)
	}

	return nil
}

func (m *mySql) VerifyUserRatings(userId, productId string) *errorModels.ServiceError {
	tx, err := m.conn.Begin()
	if err != nil {
		return errorModels.NewCustomServiceError("Error when starting transaction", err)
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
		return errorModels.NewCustomServiceError("Product already rated, you can edit it though", err) // errors.New("product already rated, you can edit it though")
	}
	defer rows.Close()

	return nil
}

func NewMySqlProductRepo(conn *sql.DB) ProductRepo {
	return &mySql{conn: conn}
}
