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

func (a *authSql) Add(auth *models.Auth) (*models.Auth, error) {
	query := `INSERT INTO auth (id, user_id, access_token, refresh_token, expires_at)
	VALUES (?, ?, ?, ?, DATE_ADD(NOW(), INTERVAL 2 HOUR))`

	_, err := a.conn.Exec(query, auth.Id, auth.UserId, auth.AccessToken, auth.RefreshToken)
	if err != nil {
		return nil, err
	}

	return a.Get(auth.UserId)
}

func (a *authSql) Get(userId string) (*models.Auth, error) {
	var auth models.Auth

	query := `SELECT id, user_id, access_token, refresh_token, expires_at, date_created, date_updated FROM auth WHERE user_id = ?`

	err := a.conn.QueryRow(query, userId).Scan(&auth.Id, &auth.UserId, &auth.AccessToken, &auth.RefreshToken, &auth.ExpiresAt, &auth.DateCreated, &auth.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func (a *authSql) Delete(userId, accessToken string) error {
	query := `DELETE FROM auth WHERE user_id = ? and access_token = ?`
	_, err := a.conn.Exec(query, userId, accessToken)
	if err != nil {
		return err
	}

	return nil
}

func (a *authSql) Clear(userId, accessToken string) error {
	query := `DELETE FROM auth WHERE user_id = ? AND access_token != ?`
	_, err := a.conn.Exec(query, userId, accessToken)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthRepo(conn *sql.DB) AuthRepo {
	return &authSql{conn: conn}
}
