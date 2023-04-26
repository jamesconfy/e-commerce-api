package repo

import (
	"database/sql"
	"e-commerce/internal/models"
)

type AuthRepo interface {
	Add(auth *models.Auth) (*models.Auth, error)
	Get(userId string) (*models.Auth, error)
	Delete(userId, accessToken string) error
	Clear(userId, accessToken string) error
}

type authSql struct {
	conn *sql.DB
}

func (a *authSql) Add(auth *models.Auth) (auh *models.Auth, err error) {
	auh = new(models.Auth)

	query := `INSERT INTO auth (user_id, access_token, refresh_token) VALUES ($1, $2, $3) RETURNING id, user_id, access_token, refresh_token, expires_at, date_created, date_updated`

	err = a.conn.QueryRow(query, auth.UserId, auth.AccessToken, auth.RefreshToken).Scan(&auh.Id, &auh.UserId, &auh.AccessToken, &auh.RefreshToken, &auh.ExpiresAt, &auh.DateCreated, &auh.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (a *authSql) Get(userId string) (*models.Auth, error) {
	var auth models.Auth

	query := `SELECT id, user_id, access_token, refresh_token, expires_at, date_created, date_updated FROM auth WHERE user_id = $1`

	err := a.conn.QueryRow(query, userId).Scan(&auth.Id, &auth.UserId, &auth.AccessToken, &auth.RefreshToken, &auth.ExpiresAt, &auth.DateCreated, &auth.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (a *authSql) Delete(userId, accessToken string) error {
	query := `DELETE FROM auth WHERE user_id = $1 and access_token = $2`
	_, err := a.conn.Exec(query, userId, accessToken)
	if err != nil {
		return err
	}

	return nil
}

func (a *authSql) Clear(userId, accessToken string) error {
	query := `DELETE FROM auth WHERE user_id = $1 AND access_token != $2`
	_, err := a.conn.Exec(query, userId, accessToken)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthRepo(conn *sql.DB) AuthRepo {
	return &authSql{conn: conn}
}
