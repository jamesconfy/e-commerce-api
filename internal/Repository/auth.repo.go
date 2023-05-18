package repo

import (
	"database/sql"
	"time"

	"e-commerce/internal/models"
	"e-commerce/utils"
)

type AuthRepo interface {
	Add(auth *models.Auth) (*models.Auth, error)
	Get(userId, url string) (*models.Auth, error)
	Delete(userId, accessToken string) error
	Clear(userId, accessToken string) error
}

type authSql struct {
	conn *sql.DB
}

func (a *authSql) Add(auth *models.Auth) (auh *models.Auth, err error) {
	auh = new(models.Auth)

	expires_at := utils.AppConfig.EXPIRES_AT
	if expires_at != 0 {
		query := `INSERT INTO auth (user_id, access_token, refresh_token, expires_at) VALUES ($1, $2, $3, $4) RETURNING id, user_id, access_token, refresh_token, expires_at, date_created, date_updated`

		err = a.conn.QueryRow(query, auth.UserId, auth.AccessToken, auth.RefreshToken, a.getExpiry(expires_at)).Scan(&auh.Id, &auh.UserId, &auh.AccessToken, &auh.RefreshToken, &auh.ExpiresAt, &auh.DateCreated, &auh.DateUpdated)
		if err != nil {
			return
		}

		return
	}

	query := `INSERT INTO auth (user_id, access_token, refresh_token) VALUES ($1, $2, $3) RETURNING id, user_id, access_token, refresh_token, expires_at, date_created, date_updated`

	err = a.conn.QueryRow(query, auth.UserId, auth.AccessToken, auth.RefreshToken).Scan(&auh.Id, &auh.UserId, &auh.AccessToken, &auh.RefreshToken, &auh.ExpiresAt, &auh.DateCreated, &auh.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (a *authSql) Get(userId, url string) (*models.Auth, error) {
	var auth models.Auth

	query := `SELECT id, user_id, access_token, refresh_token, expires_at, date_created, date_updated FROM auth WHERE user_id = $1 AND access_token = $2`

	err := a.conn.QueryRow(query, userId, url).Scan(&auth.Id, &auth.UserId, &auth.AccessToken, &auth.RefreshToken, &auth.ExpiresAt, &auth.DateCreated, &auth.DateUpdated)

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

func (a *authSql) getExpiry(expires_at int) time.Time {
	return time.Now().Add(time.Hour * time.Duration(int64(expires_at)))
}
