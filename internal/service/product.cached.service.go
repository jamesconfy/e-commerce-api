package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	se "e-commerce/internal/se"
	"fmt"
)

type cachedProductService struct {
	productService ProductService
	cache          repo.Cache
}

// Add implements ProductService
func (c *cachedProductService) Add(req *forms.Product, userId string) (*models.Product, *se.ServiceError) {
	c.cache.DeleteByTag("products")
	return c.productService.Add(req, userId)
}

// AddRating implements ProductService
func (c *cachedProductService) AddRating(req *forms.Rating, productId string, userId string) (*models.Rating, *se.ServiceError) {
	return c.productService.AddRating(req, productId, userId)
}

// Delete implements ProductService
func (c *cachedProductService) Delete(productId string, userId string) *se.ServiceError {
	c.cache.DeleteByTag(fmt.Sprintf("products:%s", productId), "products")
	return c.productService.Delete(productId, userId)
}

// Edit implements ProductService
func (c *cachedProductService) Edit(req *forms.EditProduct, productId string, userId string) (*models.Product, *se.ServiceError) {
	product, err := c.productService.Edit(req, productId, userId)
	if err == nil {
		c.cache.DeleteByTag("products", fmt.Sprintf("products:%s", productId))
	}

	return product, err
}

// Get implements ProductService
func (c *cachedProductService) Get(productId string) (*models.ProductRating, *se.ServiceError) {
	var product *models.ProductRating

	key := fmt.Sprintf("products:%s", productId)
	err := c.cache.Get(key, &product)
	if err == nil {
		return product, nil
	}

	product, er := c.productService.Get(productId)
	if er != nil {
		return nil, er
	}

	c.cache.AddByTag(key, product, productId)
	return product, nil
}

// GetAll implements ProductService
func (c *cachedProductService) GetAll(page int) ([]*models.ProductRating, *se.ServiceError) {
	var products []*models.ProductRating

	key := fmt.Sprintf("products:all:%d", page)
	err := c.cache.Get(key, &products)
	if err == nil {
		return products, nil
	}

	products, er := c.productService.GetAll(page)
	if er != nil {
		return nil, er
	}

	c.cache.AddByTag(key, products, "products")
	return products, nil
}

// Validate implements ProductService
func (c *cachedProductService) Validate(req any) error {
	return c.productService.Validate(req)
}

func NewCachedProductService(productService ProductService, cache repo.Cache) ProductService {
	return &cachedProductService{productService: productService, cache: cache}
}
