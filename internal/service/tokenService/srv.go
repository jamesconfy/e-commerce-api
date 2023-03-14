package tokenService

import (
	"e-commerce/internal/service/loggerService"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Email  string
	Id     string
	Status string
	jwt.RegisteredClaims
}

type TokenSrv interface {
	CreateToken(id, email string) (string, string, error)
	ValidateToken(token string) (*Token, error)
}

type tokenSrv struct {
	SecretKey string
	logSrv    loggerService.LogSrv
}

func (t *tokenSrv) CreateToken(id, email string) (string, string, error) {
	tokenDetails := &Token{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(60))),
		},
	}

	refreshTokenDetails := &Token{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenDetails).SignedString([]byte(t.SecretKey))
	if err != nil {
		t.logSrv.Warning(fmt.Sprintf("Could not create access token created for User: %s with Email: %s", id, email))
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenDetails).SignedString([]byte(t.SecretKey))
	if err != nil {
		t.logSrv.Warning(fmt.Sprintf("Could not create refresh token created for User: %s with Email: %s", id, email))
		return "", "", err
	}

	t.logSrv.Info(fmt.Sprintf("Access and Refresh token created for User: %s with Email: %s", id, email))
	return token, refreshToken, err
}

func (t *tokenSrv) ValidateToken(tokenUrl string) (*Token, error) {
	token, err := jwt.ParseWithClaims(
		tokenUrl,
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

	if !claims.ExpiresAt.Time.Before(time.Now().Local()) {
		return nil, err
	}

	return claims, err

}

func NewTokenSrv(secret string, logSrv loggerService.LogSrv) TokenSrv {
	return &tokenSrv{SecretKey: secret, logSrv: logSrv}
}
