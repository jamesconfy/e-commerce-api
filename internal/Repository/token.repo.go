package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type TokenRepo interface {
	Confirm(userId string) (*models.Auth, error)
}

type tokenSql struct {
	conn *sql.DB
}

func (t *tokenSql) Confirm(userId string) (*models.Auth, error) {
	var token models.Auth

	stmt := fmt.Sprintf(`SELECT access_token, refresh_token FROM users WHERE id = "%s"`, userId)
	row := t.conn.QueryRow(stmt)

	if err := row.Scan(&token.AccessToken, &token.RefreshToken); err != nil {
		return nil, err
	}

	return &token, nil
}

func NewTokenRepo(conn *sql.DB) TokenRepo {
	return &tokenSql{conn: conn}
}
