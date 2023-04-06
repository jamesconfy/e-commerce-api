package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type AuthRepo interface {
	Create(auth *models.Auth) (*models.Auth, error)
	Confirm(userId string) (*models.Auth, error)
	Get(userId string) (*models.Auth, error)
}

type authSql struct {
	conn *sql.DB
}

func (a *authSql) Confirm(userId string) (*models.Auth, error) {
	var auth models.Auth

	stmt := fmt.Sprintf(`SELECT access_token, refresh_token, expires_at, date_created, date_updated FROM auth WHERE user_id = "%s"`, userId)
	row := a.conn.QueryRow(stmt)

	if err := row.Scan(&auth.AccessToken, &auth.RefreshToken, &auth.ExpiresAt, &auth.DateCreated, &auth.DateUpdated); err != nil {
		return nil, err
	}

	return &auth, nil
}

func (a *authSql) Create(auth *models.Auth) (*models.Auth, error) {
	query := `INSERT INTO auth (id, user_id, access_token, refresh_token, expires_at, date_created, date_updated) 
	VALUES ('%[1]v', '%[2]v', '%[3]v', '%[4]v', DATE_ADD(NOW(), INTERVAL 24 HOUR), CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP())
	ON DUPLICATE KEY UPDATE access_token = '%[3]v', refresh_token = '%[4]v', expires_at = DATE_ADD(NOW(), INTERVAL 24 HOUR), date_updated = CURRENT_TIMESTAMP()`

	stmt := fmt.Sprintf(query, auth.Id, auth.UserId, auth.AccessToken, auth.RefreshToken, auth.ExpiresAt)

	_, err := a.conn.Exec(stmt)
	if err != nil {
		return nil, err
	}

	return a.Get(auth.UserId)
}

func (a *authSql) Get(userId string) (*models.Auth, error) {
	var auth models.Auth

	query := `SELECT id, user_id, access_token, refresh_token, expires_at, date_created, date_updated FROM auth WHERE user_id = '%s'`
	stmt := fmt.Sprintf(query, userId)

	row := a.conn.QueryRow(stmt)
	err := row.Scan(&auth.Id, &auth.UserId, &auth.AccessToken, &auth.RefreshToken, &auth.ExpiresAt, &auth.DateCreated, &auth.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}

func NewAuthRepo(conn *sql.DB) AuthRepo {
	return &authSql{conn: conn}
}
