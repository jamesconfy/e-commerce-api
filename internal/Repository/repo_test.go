package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"e-commerce/internal/service/timeService"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// var db *sql.DB

var db, _ = sql.Open("mysql", "root:password@tcp(localhost:3306)/e_commerce_api")

func TestAddUser(t *testing.T) {
	// Create a usersql instance
	userSql := &userSql{conn: db}

	// Create a new user object
	user := &models.User{
		Id:          "4567",
		FirstName:   "Confidence",
		LastName:    "James",
		Email:       "bobdence@gmail.com",
		PhoneNumber: "08149795370",
		Password:    "123456",
		DateCreated: timeService.New().CurrentTime(),
		DateUpdated: timeService.New().CurrentTime(),
	}

	// Create a new cart object
	cart := &models.Cart{
		Id:          "123",
		UserId:      "4567",
		DateCreated: timeService.New().CurrentTime(),
	}

	// Create a new user cart object
	usercart := &models.UserCart{
		User: user,
		Cart: cart,
	}

	// Cart the Register User function
	err := userSql.Register(usercart, "accesstoken", "refreshtoken")
	if err != nil {
		t.Errorf("Failed to save user: %v", err)
	}

	// Cart the Create Cart function
	err = userSql.CreateCart(usercart)
	if err != nil {
		t.Errorf("Failed to create cart: %v", err)
	}

	var count int
	// Query the database and check if a new user has been created
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", user.Id).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 row in users table, got %d", count)
	}
}

func TestAddProduct(t *testing.T) {
	// Create a new product to add.
	product := &models.Product{
		Id:          "123",
		UserId:      "4567",
		Name:        "Test Product",
		Description: "This is a test product.",
		Price:       9.99,
		DateCreated: timeService.New().CurrentTime(),
		DateUpdated: timeService.New().CurrentTime(),
		Image:       "test.png",
	}

	// Create a new productSql instance with the test database connection.
	productSql := &productSql{conn: db}

	// Call the Add function to add the product to the database.
	err := productSql.Add(product)
	if err != nil {
		t.Fatalf("failed to add product: %v", err)
	}

	// Check that the product was added correctly.
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", product.Id).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}

	if count == 0 {
		t.Errorf("expected at least 1 row in products table, got %d", count)
	}
}

func Test_ProductSql_Add(t *testing.T) {
	type fields struct {
		conn *sql.DB
	}

	type args struct {
		product *models.Product
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &productSql{
				conn: tt.fields.conn,
			}
			if err := p.Add(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("productSql.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
