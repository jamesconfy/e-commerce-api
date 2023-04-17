package service

import (
	repo "e-commerce/internal/repository"
)

type cachedAuthService struct {
	authService AuthService
	cache       repo.Cache
}

// Create implements AuthService
func (a *cachedAuthService) Create(id string, email string) (string, string, error) {
	return a.authService.Create(id, email)
}

// Validate implements AuthService
func (a *cachedAuthService) Validate(token string) (*Token, error) {
	var tok *Token

	err := a.cache.Get(token, &tok)
	if err == nil {
		return tok, nil
	}

	toke, er := a.authService.Validate(token)
	if er != nil {
		return nil, er
	}

	a.cache.Add(token, toke)
	return toke, er
}

func NewCachedAuthService(authService AuthService, cache repo.Cache) AuthService {
	return &cachedAuthService{authService: authService, cache: cache}
}
