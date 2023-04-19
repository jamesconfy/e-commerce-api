package service

import (
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/se"
	"fmt"
)

type cachedCartService struct {
	cartService CartService
	cache       repo.Cache
}

// Clear implements CartService
func (c *cachedCartService) Clear(userId string) *se.ServiceError {
	key := fmt.Sprintf("/api/v1/carts/%s", userId)
	c.cache.Delete(key)
	return c.cartService.Clear(userId)
}

// Get implements CartService
func (c *cachedCartService) Get(userId string) (*models.Cart, *se.ServiceError) {
	var cart *models.Cart

	key := fmt.Sprintf("/api/v1/carts/%s", userId)
	err := c.cache.Get(key, &cart)
	if err == nil {
		return cart, nil
	}

	cart, er := c.cartService.Get(userId)
	if er != nil {
		return cart, nil
	}

	c.cache.AddByTag(key, cart, userId)
	return cart, er
}

func NewCachedCartService(cartService CartService, cache repo.Cache) CartService {
	return &cachedCartService{cartService: cartService, cache: cache}
}
