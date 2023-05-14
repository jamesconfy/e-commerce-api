package service

import (
	repo "e-commerce/internal/repository"
	"fmt"
)

type cachedAuthService struct {
	authService AuthService
	cache       repo.Cache
}

// Create implements AuthService
func (a *cachedAuthService) Create(id string, email string) (accessToken, refreshToken string, err error) {
	return a.authService.Create(id, email)
}

// Validate implements AuthService
func (a *cachedAuthService) Validate(url string) (toke *Token, err error) {
	var tok *Token

	key := fmt.Sprintf("validate:%v", url)
	err = a.cache.Get(key, &tok)
	if err == nil {
		return
	}

	toke, er := a.authService.Validate(url)
	if er != nil {
		return
	}

	a.cache.AddByTag(key, toke, toke.Id)
	return
}

func NewCachedAuthService(authService AuthService, cache repo.Cache) AuthService {
	return &cachedAuthService{authService: authService, cache: cache}
}
