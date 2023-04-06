package repo

import (
	"database/sql"
	"e-commerce/internal/models"
)

type ProductRepo interface {
	Add(product *models.Product) (*models.Product, error)
	GetAll(page int) ([]*models.ProductRating, error)
	GetId(productId string) (*models.ProductRating, error)
	Edit(product *models.Product) (*models.Product, error)
	Delete(productId string) error
	AddRating(rating *models.Rating) (*models.Rating, error)
	GetRating(productId, userId string) (*models.Rating, error)
}

type productSql struct {
	conn *sql.DB
}

func (p *productSql) Add(product *models.Product) (*models.Product, error) {
	query := `INSERT INTO products(id, user_id, name, description, price, image) 
				VALUES(?, ?, ?, ?, ?, ?)`

	_, err := p.conn.Exec(query, product.Id, product.UserId, product.Name, product.Description, product.Price, product.Image)
	if err != nil {
		return nil, err
	}

	result, err := p.GetId(product.Id)
	if err != nil {
		return nil, err
	}

	return result.Product, nil
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

	query := `SELECT p.id, p.user_id, p.name, p.description, p.price, p.date_updated, p.date_created, p.image, 
			IFNULL(AVG(r.rating), 0.00) AS rating 
			FROM products p 
			LEFT JOIN ratings r 
			ON r.product_id = p.id 
			GROUP BY p.id 
			ORDER BY p.id LIMIT ? OFFSET ?;`

	rows, err := tx.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.ProductRating
		product.Product = &models.Product{}

		err := rows.Scan(&product.Product.Id, &product.Product.UserId, &product.Product.Name, &product.Product.Description, &product.Product.Price, &product.Product.DateCreated, &product.Product.DateUpdated, &product.Product.Image, &product.Rating)

		if err != nil {
			return nil, err
		}

		product.ProductId = product.Product.Id
		products = append(products, &product)
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

	query := `SELECT p.id, p.user_id, p.name, p.description, p.price, p.date_updated, p.date_created, p.image, 
			IFNULL(AVG(r.rating), 0.00) 
			AS rating 
			FROM products p 
			LEFT JOIN ratings r 
			ON r.product_id = p.id 
			WHERE p.id = ? 
			GROUP BY p.id 
			ORDER BY p.id`

	row := tx.QueryRow(query, productId)

	err = row.Scan(&product.Product.Id, &product.Product.UserId, &product.Product.Name,
		&product.Product.Description, &product.Product.Price, &product.Product.DateCreated,
		&product.Product.DateUpdated, &product.Product.Image, &product.Rating)

	if err != nil {
		return nil, err
	}

	product.ProductId = product.Product.Id

	return &product, nil
}

func (p *productSql) Edit(product *models.Product) (*models.Product, error) {
	query := `UPDATE products SET name = ?, description = ?, price = ?, image = ?, date_updated = CURRENT_TIMESTAMP() WHERE user_id = ? AND id = ?`

	_, err := p.conn.Exec(query, product.Name, product.Description, product.Price, product.Image, product.UserId, product.Id)
	if err != nil {
		return nil, err
	}

	result, err := p.GetId(product.Id)
	if err != nil {
		return nil, err
	}

	return result.Product, nil
}

func (p *productSql) Delete(productId string) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := p.conn.Exec(query, productId)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) AddRating(rating *models.Rating) (*models.Rating, error) {
	query := `INSERT INTO ratings(rating, product_id, user_id)
			VALUES(?, ?, ?) 
			ON DUPLICATE KEY UPDATE rating = ?, date_updated = CURRENT_TIMESTAMP()`

	_, err := p.conn.Exec(query, rating.Value, rating.ProductId, rating.UserId, rating.Value)
	if err != nil {
		return nil, err
	}

	return p.GetRating(rating.ProductId, rating.UserId)
}

func (p *productSql) GetRating(productId, userId string) (*models.Rating, error) {
	query := `SELECT rating, product_id, user_id, date_created, date_updated 
			FROM ratings WHERE product_id = ? AND user_id = ?`

	row := p.conn.QueryRow(query, productId, userId)

	var rating models.Rating

	err := row.Scan(&rating.Value, &rating.ProductId, &rating.UserId, &rating.DateCreated, &rating.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func NewProductRepo(conn *sql.DB) ProductRepo {
	return &productSql{conn: conn}
}
