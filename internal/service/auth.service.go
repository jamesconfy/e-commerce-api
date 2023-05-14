package service

import (
	"database/sql"
	repo "e-commerce/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Email string
	Id    string
	jwt.RegisteredClaims
}

type AuthService interface {
	Create(id, email string) (string, string, error)
	Validate(url string) (*Token, error)
}

type authSrv struct {
	authRepo  repo.AuthRepo
	SecretKey string
	logSrv    LogSrv
}

func (t *authSrv) Create(id, email string) (accessToken, refreshToken string, err error) {
	accessTokenDetails := &Token{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(2))),
		},
	}

	refreshTokenDetails := &Token{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(72))),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenDetails).SignedString([]byte(t.SecretKey))
	if err != nil {
		return
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenDetails).SignedString([]byte(t.SecretKey))
	if err != nil {
		return
	}

	return
}

func (t *authSrv) Validate(url string) (*Token, error) {
	token, err := jwt.ParseWithClaims(
		url,
		&Token{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.SecretKey), nil
		},
	)

	if token == nil {
		return nil, errors.New("check the provided token")
	}

	claims, ok := token.Claims.(*Token)
	if !ok {
		return nil, err
	}

	if err := claims.Valid(); err != nil {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now().Local()) {
		return nil, fmt.Errorf("expired token, please login again || expired time: %s", claims.ExpiresAt.Time)
	}

	row, err := t.authRepo.Get(claims.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("outdated token")
		}
		return nil, err
	}

	if row.ExpiresAt.Before(time.Now().Local()) {
		return nil, fmt.Errorf("token is expired")
	}

	return claims, err
}

func NewAuthService(repo repo.AuthRepo, secret string, logSrv LogSrv) AuthService {
	return &authSrv{authRepo: repo, SecretKey: secret, logSrv: logSrv}
}
