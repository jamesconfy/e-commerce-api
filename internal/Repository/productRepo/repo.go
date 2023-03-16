package productRepo

import (
	"database/sql"
	"e-commerce/internal/models/productModels"
)

type ProductRepo interface {
	AddProduct(req *productModels.AddProductReq) error
}

type mySql struct {
	conn *sql.DB
}

func (m *mySql) AddProduct(req *productModels.AddProductReq) error {
	return nil
}

func NewMySqlUserRepo(conn *sql.DB) ProductRepo {
	return &mySql{conn: conn}
}
