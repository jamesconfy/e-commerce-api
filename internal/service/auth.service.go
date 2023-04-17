package service

import (
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
	Validate(token string) (*Token, error)
}

type authSrv struct {
	authRepo  repo.AuthRepo
	SecretKey string
	logSrv    LogSrv
}

func (t *authSrv) Create(id, email string) (string, string, error) {
	accessTokenDetails := &Token{
		// User Email
		Email: email,
		// User Id
		Id: id,
		// Registered Claims
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}

	refreshTokenDetails := &Token{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(72))),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenDetails).SignedString([]byte(t.SecretKey))
	if err != nil {
		// t.logSrv.Warning(fmt.Sprintf("Could not create access token created for User: %s with Email: %s", id, email))
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenDetails).SignedString([]byte(t.SecretKey))
	if err != nil {
		//t.logSrv.Warning(fmt.Sprintf("Could not create refresh token created for User: %s with Email: %s", id, email))
		return "", "", err
	}

	// t.logSrv.Info(fmt.Sprintf("Access and Refresh token created for UserId: %s with Email: %s", id, email))
	// cookie := http.Cookie{
	// 	Name:     "access_token",
	// 	Value:    token,
	// 	HttpOnly: true,
	// 	Path:     "/",
	// 	Secure:   false,
	// }

	// http.SetCookie(*c, cookie)
	return accessToken, refreshToken, err
}

func (t *authSrv) Validate(tokenUrl string) (*Token, error) {
	token, err := jwt.ParseWithClaims(
		tokenUrl,
		&Token{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.SecretKey), nil
		},
	)

	if token == nil {
		// t.logSrv.Debug(fmt.Sprintf("Token is empty after parsing || Token String: %s", tokenUrl))
		return nil, errors.New("check the provided token")
	}

	claims, ok := token.Claims.(*Token)
	if !ok {
		// t.logSrv.Debug(fmt.Sprintf("Token claims is not ok || Claims: %s", claims))
		return nil, err
	}

	if err := claims.Valid(); err != nil {
		// t.logSrv.Debug(fmt.Sprintf("Token claims is not valid || Claims: %s", claims))
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now().Local()) {
		// t.logSrv.Debug(fmt.Sprintf("Token is expired || Expired Time: %s", claims.ExpiresAt))
		return nil, fmt.Errorf("expired token, please login again || expired time: %s", claims.ExpiresAt.Time)
	}

	row, err := t.authRepo.Get(claims.Id)
	if err != nil {
		// t.logSrv.Fatal(fmt.Sprintf("Internal server error, when trying to get token associated with user || UserId: %s", claims.Id))
		return nil, err
	}

	if row.AccessToken != tokenUrl {
		// t.logSrv.Info(fmt.Sprintf("User tried to use old access token to login || UserId: %s || Latest Access Token: %s || Provided Access Token: %s", claims.Id, row.AccessToken, tokenUrl))
		return nil, fmt.Errorf("outdated token")
	}

	if row.ExpiresAt.Before(time.Now().Local()) {
		return nil, fmt.Errorf("token is expired")
	}

	// t.logSrv.Info(fmt.Sprintf("Access token validated for UserId: %s with Email: %s and Expires At: %s", claims.Id, claims.Email, claims.ExpiresAt))
	return claims, err
}

func NewAuthService(repo repo.AuthRepo, secret string, logSrv LogSrv) AuthService {
	return &authSrv{authRepo: repo, SecretKey: secret, logSrv: logSrv}
}
