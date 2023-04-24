package repo

import (
	"database/sql"
	"e-commerce/internal/models"
)

type ProductRepo interface {
	Add(product *models.Product) (*models.Product, error)
	GetAll(page int) ([]*models.ProductRating, error)
	Get(productId string) (*models.ProductRating, error)
	Edit(product *models.Product) (*models.Product, error)
	Delete(productId string) error
	AddRating(rating *models.Rating) (*models.Rating, error)
	GetRating(productId, userId string) (*models.Rating, error)
}

type productSql struct {
	conn *sql.DB
}

func (p *productSql) Add(product *models.Product) (prod *models.Product, err error) {
	prod = new(models.Product)

	query := `INSERT INTO products(user_id, name, description, price, image) VALUES($1, $2, $3, $4, $5) RETURNING id, user_id, name, description, price, image, date_created, date_updated`

	err = p.conn.QueryRow(query, product.UserId, product.Name, product.Description, product.Price, product.Image).Scan(&prod.Id, &prod.UserId, &prod.Name, &prod.Description, &prod.Price, &prod.Image, &prod.DateCreated, &prod.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (p *productSql) GetAll(page int) ([]*models.ProductRating, error) {
	var products []*models.ProductRating

	limit := 20
	offset := limit * (page - 1)

	query := `SELECT p.id, p.user_id, p.name, p.description, p.price, p.date_updated, p.date_created, p.image, COALESCE(AVG(r.rating), 0.00) AS rating FROM products p LEFT JOIN ratings r ON r.product_id = p.id GROUP BY p.id ORDER BY p.id LIMIT $1 OFFSET $2;`

	rows, err := p.conn.Query(query, limit, offset)
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

func (p *productSql) Get(productId string) (prod *models.ProductRating, err error) {
	prod = new(models.ProductRating)
	prod.Product = &models.Product{}

	query := `SELECT p.id, p.user_id, p.name, p.description, p.price, p.date_updated, p.date_created, p.image, COALESCE(AVG(r.rating), 0.00) AS rating FROM products p LEFT JOIN ratings r ON r.product_id = p.id WHERE p.id = $1 GROUP BY p.id ORDER BY p.id`

	err = p.conn.QueryRow(query, productId).Scan(&prod.Product.Id, &prod.Product.UserId, &prod.Product.Name, &prod.Product.Description, &prod.Product.Price, &prod.Product.DateCreated, &prod.Product.DateUpdated, &prod.Product.Image, &prod.Rating)

	if err != nil {
		return
	}

	prod.ProductId = prod.Product.Id

	return
}

func (p *productSql) Edit(product *models.Product) (prod *models.Product, err error) {
	prod = new(models.Product)

	query := `UPDATE products SET name = $1, description = $2, price = $3, image = $4, date_updated = CURRENT_TIMESTAMP WHERE user_id = $5 AND id = $6 RETURNING id, user_id, name, description, price, image, date_created, date_updated`

	err = p.conn.QueryRow(query, product.Name, product.Description, product.Price, product.Image, product.UserId, product.Id).Scan(&prod.Id, &prod.UserId, &prod.Name, &prod.Description, &prod.Price, &prod.Image, &prod.DateCreated, &prod.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (p *productSql) Delete(productId string) (err error) {
	query := `DELETE FROM products WHERE id = $1`
	_, err = p.conn.Exec(query, productId)
	if err != nil {
		return
	}

	return
}

func (p *productSql) AddRating(rating *models.Rating) (rate *models.Rating, err error) {
	rate = new(models.Rating)

	query := `INSERT INTO ratings(rating, product_id, user_id) VALUES($1, $2, $3) ON CONFLICT (product_id, user_id)
	DO UPDATE SET rating = $1, date_updated = CURRENT_TIMESTAMP RETURNING rating, product_id, user_id, date_created, date_updated`

	err = p.conn.QueryRow(query, rating.Value, rating.ProductId, rating.UserId).Scan(&rate.Value, &rate.ProductId, &rate.UserId, &rate.DateCreated, &rate.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (p *productSql) GetRating(productId, userId string) (rate *models.Rating, err error) {
	rate = new(models.Rating)
	query := `SELECT rating, product_id, user_id, date_created, date_updated FROM ratings WHERE product_id = $1 AND user_id = $2`

	err = p.conn.QueryRow(query, productId, userId).Scan(&rate.Value, &rate.ProductId, &rate.UserId, &rate.DateCreated, &rate.DateUpdated)
	if err != nil {
		return
	}

	return
}

func NewProductRepo(conn *sql.DB) ProductRepo {
	return &productSql{conn: conn}
}
