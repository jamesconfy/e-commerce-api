package logger

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	"fmt"
)

func (m Messages) AddSuccess(product *models.Product) (str string) {
	str = fmt.Sprintf("Product created successfully || Id: %s || Name: %s || Description: %s || Price: %v || DateCreated: %s", product.Id, product.Name, product.Description, product.Price, product.DateCreated)
	return
}

func (m Messages) AddRepoError(product *models.Product, err error) (str string) {
	str = fmt.Sprintf("Error occured when adding product to database || Id: %s  || DateCreated: %s || Error: %s", product.Id, err, product.DateCreated)
	return
}

func (m Messages) GetAllRepoError(err error) (str string) {
	str = fmt.Sprintf("Error occured when getting all product || Error: %s", err)
	return
}

func (m Messages) GetAllSuccess(products []*models.ProductRating) (str string) {
	str = fmt.Sprintf("Products successfully gotten || Products: %v", products)
	return
}

func (m Messages) GetProductRepoError(productId string, err error) (str string) {
	str = fmt.Sprintf("Error occured when getting product || ProductId: %s || Error: %v", productId, err)
	return
}

func (m Messages) GetProductSuccess(product *models.Product) (str string) {
	str = fmt.Sprintf("Product successfully gotten || Id: %s || ProductName: %s || ProductDescription: %s || DateCreated: %s", product.Id, product.Name, product.Description, product.DateCreated)
	return
}

func (m Messages) EditCompareUser(productUserId, userId string) (str string) {
	str = fmt.Sprintf("User tried to change resource || Product UserId: %v || Logged In UserId: %v", productUserId, userId)
	return
}

func (m Messages) EditProductRepoError(product *models.Product, err error) (str string) {
	str = fmt.Sprintf("Error when editing product || Id: %s || UserId: %s || Name: %v || Description: %v || Price: %v || Image: %v || DateCreated: || %v || Error: %v", product.Id, product.UserId, product.Name, product.Description, product.Price, product.Image, product.DateUpdated, err)
	return
}

func (m Messages) EditProductSuccess(product *models.Product) (str string) {
	str = fmt.Sprintf("Product edited successfully || Id: %v || UserId: %v || Name: %v || Description: %v || Image: %v || DateUpdated: %v", product.Id, product.UserId, product.Name, product.Description, product.Image, product.DateUpdated)
	return
}

func (m Messages) DeleteProductRepoError(productId string, err error) (str string) {
	str = fmt.Sprintf("Error occured when deleting product || ProductId: %s || Error: %s", productId, err)
	return
}

func (m Messages) DeleteProductSuccess(productId string) (str string) {
	str = fmt.Sprintf("Product successfully deleted || ProductId: %s", productId)
	return
}

func (m Messages) AddRatingCompareUser(rating *forms.Rating, productUserId, userId string) (str string) {
	str = fmt.Sprintf("User need to be the different before you can add rating || Rating: %v || Product UserId: %s || UserId: %s", rating.Value, productUserId, userId)
	return
}

func (m Messages) AddRatingRepoError(req *models.Rating, err error) (str string) {
	str = fmt.Sprintf("Error when adding rating request to database || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s || Error: %v", req.Value, req.ProductId, req.UserId, req.DateCreated, err)
	return
}

func (m Messages) AddRatingSuccess(req *models.Rating) (str string) {
	str = fmt.Sprintf("Rating successfully added || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.Value, req.ProductId, req.UserId, req.DateCreated)
	return
}
