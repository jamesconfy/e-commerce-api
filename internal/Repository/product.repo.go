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
	VerifyRating(userId, productId string) error
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

	query1 := `SELECT id, user_id, name, description, price, date_updated, date_created, image FROM products ORDER BY product_id LIMIT %[1]v OFFSET %[2]v`
	query2 := `SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%[1]v'`

	stmt := fmt.Sprintf(query1, limit, offset)
	rows, err := tx.Query(stmt)
	if err != nil {
		if err != sql.ErrNoRows {
			return []*models.ProductRating{}, err
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.ProductRating

		if err = rows.Scan(&product.Product.Id, &product.Product.UserId, &product.Product.Name,
			&product.Product.Description, &product.Product.Price, &product.Product.DateCreated,
			&product.Product.DateUpdated, &product.Product.Image); err != nil {
			return nil, err
		}

		product.ProductId = product.Product.Id
		products = append(products, &product)
	}

	for _, product := range products {
		stmt2 := fmt.Sprintf(query2, product.Product.Id)
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

	query1 := `SELECT id, user_id, name, description, price, date_updated, date_created, image FROM products WHERE product_id = '%[1]s'`
	query2 := `SELECT IFNULL(AVG(rating), 0.00) AS rating FROM ratings WHERE product_id = '%[1]v'`

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
	tx, err := p.conn.Begin()
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

	stmt := fmt.Sprintf(`UPDATE products 
						SET name = '%v', description = '%v', price = %v, date_updated = '%v', image = '%v'
						WHERE user_id = '%v' AND product_id = '%v'`, product.Name, product.Description, product.Price, product.DateUpdated, product.Image, product.UserId, product.Id)

	_, err = tx.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) Delete(productId string) error {
	tx, err := p.conn.Begin()
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

	stmt1 := fmt.Sprintf(`DELETE FROM products WHERE product_id = '%v'`, productId)
	_, err = tx.Exec(stmt1)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) AddRating(rating *models.Rating) error {
	tx, err := p.conn.Begin()
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
		)`, rating.Id, rating.Value, rating.ProductId, rating.UserId, rating.DateCreated, rating.DateUpdated)

	_, err = tx.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) VerifyRating(userId, productId string) error {
	tx, err := p.conn.Begin()
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

	var id string

	stmt := fmt.Sprintf(`SELECT rating_id FROM ratings WHERE user_id = '%s' AND product_id = '%s'`, userId, productId)
	row := tx.QueryRow(stmt)

	if err := row.Scan(&id); err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func NewProductRepo(conn *sql.DB) ProductRepo {
	return &productSql{conn: conn}
}
