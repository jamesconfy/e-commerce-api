package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	"e-commerce/internal/se"
	"fmt"
)

type cachedUserService struct {
	userService UserService
	cache       repo.Cache
}

func (c *cachedUserService) Validate(req any) error {
	return c.userService.Validate(req)
}

// Add implements UserService
func (c *cachedUserService) Add(req *forms.Signup) (*models.UserCart, *se.ServiceError) {
	c.cache.DeleteByTag("users:all")
	return c.userService.Add(req)
}

// Delete implements UserService
func (c *cachedUserService) Delete(userId string) *se.ServiceError {
	c.cache.DeleteByTag("users:all", userId)
	return c.userService.Delete(userId)
}

// DeleteToken implements UserService
func (c *cachedUserService) DeleteToken(userId string) *se.ServiceError {
	c.cache.DeleteByTag(userId)
	return c.userService.DeleteToken(userId)
}

// Edit implements UserService
func (c *cachedUserService) Edit(req *forms.EditUser, userId string) (*models.User, *se.ServiceError) {
	user, err := c.userService.Edit(req, userId)
	if err == nil {
		key := fmt.Sprintf("users:%s", userId)

		c.cache.DeleteByTag("users:all", key)
	}

	return user, err
}

// GetAll implements UserService
func (c *cachedUserService) GetAll(pageI int) ([]*models.User, *se.ServiceError) {
	var users []*models.User

	key := fmt.Sprintf("users:all:%d", pageI)
	err := c.cache.Get(key, &users)
	if err == nil {
		return users, nil
	}

	users, er := c.userService.GetAll(pageI)
	if er != nil {
		return nil, er
	}

	c.cache.AddByTag(key, users, "users:all")
	return users, nil
}

// GetById implements UserService
func (c *cachedUserService) GetById(userId string) (*models.User, *se.ServiceError) {
	var user *models.User

	key := fmt.Sprintf("users:%s", userId)
	err := c.cache.Get(key, &user)
	if err == nil {
		return user, nil
	}

	user, er := c.userService.GetById(userId)
	if er != nil {
		return nil, er
	}

	c.cache.AddByTag(key, user, userId)
	return user, er
}

// Login implements UserService
func (c *cachedUserService) Login(req *forms.Login) (*models.Auth, *se.ServiceError) {
	return c.userService.Login(req)
}

func NewCachedUserService(userService UserService, cache repo.Cache) UserService {
	return &cachedUserService{userService: userService, cache: cache}
}
