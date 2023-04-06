package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
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
	query := `INSERT INTO products(id, user_id, name, description, price, image) VALUES ('%[1]v', '%[2]v', '%[3]v', '%[4]v', %[5]v, '%[6]v')`

	stmt := fmt.Sprintf(query, product.Id, product.UserId, product.Name, product.Description, product.Price, product.Image)

	_, err := p.conn.Exec(stmt)
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

	query1 := `SELECT p.id, p.user_id, p.name, p.description, p.price, p.date_updated, p.date_created, p.image, IFNULL(AVG(r.rating), 0.00) AS rating FROM products p LEFT JOIN ratings r ON r.product_id = p.id GROUP BY p.id ORDER BY p.id LIMIT %[1]v OFFSET %[2]v;`
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

	query := `SELECT p.id, p.user_id, p.name, p.description, p.price, p.date_updated, p.date_created, p.image, IFNULL(AVG(r.rating), 0.00) AS rating FROM products p LEFT JOIN ratings r ON r.product_id = p.id WHERE p.id = '%[1]v' GROUP BY p.id ORDER BY p.id`

	stmt := fmt.Sprintf(query, productId)
	row := tx.QueryRow(stmt)

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
	query := `UPDATE products SET name = '%[1]v', description = '%[2]v', price = %[3]v, date_updated = CURRENT_TIMESTAMP(), image = '%[4]v' WHERE user_id = '%[5]v' AND id = '%[6]v'`

	stmt := fmt.Sprintf(query, product.Name, product.Price, product.Description, product.Image, product.UserId, product.Id)

	_, err := p.conn.Exec(stmt)
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
	stmt1 := fmt.Sprintf(`DELETE FROM products WHERE id = '%v'`, productId)
	_, err := p.conn.Exec(stmt1)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSql) AddRating(rating *models.Rating) (*models.Rating, error) {
	query := `INSERT INTO ratings(
		rating, product_id, user_id, date_created, date_updated) VALUES(
			%[1]v, '%[2]v', '%[3]v', '%[4]v', '%[5]v'
		) ON DUPLICATE KEY UPDATE rating = %[1]v, date_updated = '%[5]v'`

	stmt := fmt.Sprintf(query, rating.Value, rating.ProductId, rating.UserId, rating.DateCreated, rating.DateUpdated)

	_, err := p.conn.Exec(stmt)
	if err != nil {
		return nil, err
	}

	return p.GetRating(rating.ProductId, rating.UserId)
}

func (p *productSql) GetRating(productId, userId string) (*models.Rating, error) {
	query := `SELECT rating, product_id, user_id, date_created, date_updated FROM ratings WHERE product_id = '%[1]v' AND user_id = '%[2]v'`
	stmt := fmt.Sprintf(query, productId, userId)
	row := p.conn.QueryRow(stmt)

	var rating models.Rating

	err := row.Scan(&rating.Value, &rating.ProductId, &rating.UserId, &rating.DateCreated, &rating.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &rating, nil
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
