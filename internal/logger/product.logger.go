package logger

// func (m Messages) AddProductValidationError(req *productModels.AddProductReq) (str string) {
// 	str = fmt.Sprintf("Error when validating add product request || UserId: %s || Product Name: %s || Product Description: %s", req.UserId, req.Name, req.Description)
// 	return
// }

// func (m Messages) AddProductSuccess(req *productModels.AddProductReq) (str string) {
// 	str = fmt.Sprintf("Product created successfully || ProductId: %s || ProductName: %s || ProductDescription: %s || DateCreated: %s", req.ProductId, req.Name, req.Description, req.DateCreated)
// 	return
// }

// func (m Messages) AddProductRepoError(req *productModels.AddProductReq, err *errorModels.ServiceError) (str string) {
// 	str = fmt.Sprintf("Error occured when adding product to database || ProductId: %s || Error: %s || DateCreated: %s", req.ProductId, err, req.DateCreated)
// 	return
// }

// // func (m Messages) GetProductsRepoError(err *errorModels.ServiceError) (str string) {
// // 	str = fmt.Sprintf("Error occured when getting all product || Error: %s", err)
// // 	return
// // }

// func (m Messages) GetProductsSuccess() (str string) {
// 	str = "Products successfully gotten"
// 	return
// }

// func (m Messages) GetProductNotFound(productId string, err *errorModels.ServiceError) (str string) {
// 	str = fmt.Sprintf("Error occured when getting product || ProductId: %s || Error: %s", productId, err)
// 	return
// }

// func (m Messages) GetProductSuccess(req *productModels.GetProductRes) (str string) {
// 	str = fmt.Sprintf("Product successfully gotten || ProductId: %s || ProductName: %s || ProductDescription: %s || DateCreated: %s", req.ProductId, req.Name, req.Description, req.DateCreated)
// 	return
// }

// func (m Messages) EditProductValidationError(req *productModels.EditProductReq) (str string) {
// 	str = fmt.Sprintf("Error when validating edit product request || ProductId: %s || UserId: %s || Name: %v || Description: %v || Price: %v || Image: %v || DateCreated: || %v", req.ProductId, req.UserId, req.Name, req.Description, req.Price, req.Image, req.DateUpdated)
// 	return
// }

// func (m Messages) EditProductSuccess(req *productModels.EditProductReq) (str string) {
// 	str = fmt.Sprintf("Product edited successfully || ProductId: %v || UserId: %v || Name: %v || Description: %v || Image: %v || DateUpdated: %v", req.ProductId, req.UserId, req.Name, req.Description, req.Image, req.DateUpdated)
// 	return
// }

// func (m Messages) DeleteProductRepoError(productId string, err *errorModels.ServiceError) (str string) {
// 	str = fmt.Sprintf("Error occured when deleting product || ProductId: %s || Error: %s", productId, err)
// 	return
// }

// func (m Messages) DeleteProductSuccess(productId string) (str string) {
// 	str = fmt.Sprintf("Product successfully deleted || ProductId: %s", productId)
// 	return
// }

// func (m Messages) AddRatingValidationError(req *productModels.AddRatingsReq) (str string) {
// 	str = fmt.Sprintf("Error when validating add rating request || RatingId: %s || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated)
// 	return
// }

// func (m Messages) AddRatingRepoError(req *productModels.AddRatingsReq) (str string) {
// 	str = fmt.Sprintf("Error when adding rating request to database || RatingId: %s || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated)
// 	return
// }

// func (m Messages) AddRatingSuccess(req *productModels.AddRatingsRes) (str string) {
// 	str = fmt.Sprintf("Rating successfully added || RatingId: %s || Rating: %v || ProductId: %s || UserId: %s || DateCreated: %s", req.RatingId, req.Rating, req.ProductId, req.UserId, req.DateCreated)
// 	return
// }

// func (m Messages) VerifyUserRatingsRepoError(userId, productId string) (str string) {
// 	str = fmt.Sprintf("User tried to re-rate a product || UserId: %s || ProductId: %s", userId, productId)
// 	return
// }

// func (m Messages) VerifyUserRatingsSucess(userId, productId string) (str string) {
// 	str = fmt.Sprintf("Product rated successfully || UserId: %s || ProductId: %s", userId, productId)
// 	return
// }
