package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/se"
	"fmt"
)

type cachedCartItemService struct {
	cartItemSrv CartItemService
	cache       repo.Cache
}

// Add implements CartItemService
func (c *cachedCartItemService) Add(req *forms.CartItem, userId string) (*models.Item, *se.ServiceError) {
	return c.cartItemSrv.Add(req, userId)
}

// Delete implements CartItemService
func (c *cachedCartItemService) Delete(productId string, userId string) *se.ServiceError {
	key := fmt.Sprintf("/api/v1/items/%s", productId)
	c.cache.Delete(key)
	return c.cartItemSrv.Delete(productId, userId)
}

// Get implements CartItemService
func (c *cachedCartItemService) Get(productId string, userId string) (*models.Item, *se.ServiceError) {
	var item *models.Item

	key := fmt.Sprintf("/api/v1/items/%s", productId)
	err := c.cache.Get(key, &item)
	if err == nil {
		return item, nil
	}

	item, er := c.cartItemSrv.Get(productId, userId)
	if er != nil {
		return nil, er
	}

	c.cache.AddByTag(key, item, userId)
	return item, er
}

// GetItems implements CartItemService
func (c *cachedCartItemService) GetItems(userId string) (*models.CartItem, *se.ServiceError) {
	var items *models.CartItem

	key := "/api/v1/items"
	err := c.cache.Get(key, &items)
	if err == nil {
		return items, nil
	}

	items, er := c.cartItemSrv.GetItems(userId)
	if er != nil {
		return nil, er
	}

	c.cache.AddByTag(key, items, userId)
	return items, er
}

func NewCachedCartItemService(cartItemSrv CartItemService, cache repo.Cache) CartItemService {
	return &cachedCartItemService{cartItemSrv: cartItemSrv, cache: cache}
}
