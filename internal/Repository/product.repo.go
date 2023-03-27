package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type ProductRepo interface {
	Add(product *models.Product) error
	GetAll(page int) ([]*models.ProductRating, error)
	GetId(productId string) (*models.ProductRating, error)
	Edit(product *models.Product) error
	Delete(productId string) error
	AddRating(rating *models.Rating) error
	// VerifyRating(userId, productId string) error
}

type productSql struct {
	conn *sql.DB
}

func (p *productSql) Add(product *models.Product) error {
	query := `INSERT INTO products(id, user_id, name, description, price, date_created, date_updated, image) VALUES ('%[1]v', '%[2]v', '%[3]v', '%[4]v', %[5]v, '%[6]v', '%[7]v', '%[8]v')`

	stmt := fmt.Sprintf(query, product.Id, product.UserId, product.Name, product.Description, product.Price, product.DateCreated, product.DateUpdated, product.Image)

	_, err := p.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) GetAll(page int) ([]*models.ProductRating, error) {
	tx, err := p.conn.Begin()
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

	var products []*models.ProductRating

	limit := 20
	offset := limit * (page - 1)

	query1 := `SELECT id, user_id, name, description, price, date_updated, date_created, image FROM products ORDER BY id LIMIT %[1]v OFFSET %[2]v`
	stmt := fmt.Sprintf(query1, limit, offset)

	rows, err := tx.Query(stmt)
	if err != nil {
		if err == sql.ErrNoRows {
			return []*models.ProductRating{}, err
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.ProductRating
		product.Product = &models.Product{}

		err := rows.Scan(&product.Product.Id, &product.Product.UserId, &product.Product.Name, &product.Product.Description, &product.Product.Price, &product.Product.DateCreated, &product.Product.DateUpdated, &product.Product.Image)

		if err != nil {
			return nil, err
		}

		product.ProductId = product.Product.Id
		products = append(products, &product)
	}

	for _, product := range products {
		query2 := `SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%[1]v'`
		stmt2 := fmt.Sprintf(query2, product.ProductId)

		if err := tx.QueryRow(stmt2).Scan(&product.Rating); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func (p *productSql) GetId(productId string) (*models.ProductRating, error) {
	tx, err := p.conn.Begin()
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

	var product models.ProductRating
	product.Product = &models.Product{}

	query1 := `SELECT id, user_id, name, description, price, date_updated, date_created, image FROM products WHERE id = '%[1]s'`
	query2 := `SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE id = '%[1]v'`

	stmt1 := fmt.Sprintf(query1, productId)
	row := tx.QueryRow(stmt1)

	err = row.Scan(&product.Product.Id, &product.Product.UserId, &product.Product.Name,
		&product.Product.Description, &product.Product.Price, &product.Product.DateCreated,
		&product.Product.DateUpdated, &product.Product.Image)

	if err != nil {
		if err == sql.ErrNoRows {
			return &models.ProductRating{}, err
		}

		return nil, err
	}

	product.ProductId = product.Product.Id

	stmt2 := fmt.Sprintf(query2, productId)
	if err := tx.QueryRow(stmt2).Scan(&product.Rating); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productSql) Edit(product *models.Product) error {
	query := `UPDATE products SET name = '%[1]v', description = '%[2]v', price = %[3]v, date_updated = '%[4]v', image = '%[5]v' WHERE user_id = '%[6]v' AND id = '%[7]v'`

	stmt := fmt.Sprintf(query, product.Name, product.Description, product.Price, product.DateUpdated, product.Image, product.UserId, product.Id)

	_, err := p.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) Delete(productId string) error {
	stmt1 := fmt.Sprintf(`DELETE FROM products WHERE id = '%v'`, productId)
	_, err := p.conn.Exec(stmt1)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) AddRating(rating *models.Rating) error {
	query := `INSERT INTO ratings(
		rating, product_id, user_id, date_created, date_updated) VALUES(
			%[1]v, '%[2]v', '%[3]v', '%[4]v', '%[5]v'
		) ON DUPLICATE KEY UPDATE rating = %[1]v, date_updated = '%[5]v'`

	stmt := fmt.Sprintf(query, rating.Value, rating.ProductId, rating.UserId, rating.DateCreated, rating.DateUpdated)

	_, err := p.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

// func (p *productSql) VerifyRating(userId, productId string) error {
// 	tx, err := p.conn.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		} else {
// 			tx.Commit()
// 		}
// 	}()

// 	var id string

// 	stmt := fmt.Sprintf(`SELECT rating_id FROM ratings WHERE user_id = '%s' AND product_id = '%s'`, userId, productId)
// 	row := tx.QueryRow(stmt)

// 	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
// 		return err
// 	}

// 	return nil
// }

func NewProductRepo(conn *sql.DB) ProductRepo {
	return &productSql{conn: conn}
}
