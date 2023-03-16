package tokenRepo

import (
	"database/sql"
	tokenModels "e-commerce/internal/models/tokenModels"
	"fmt"
)

type TokenRepo interface {
	ConfirmToken(userId string) (*tokenModels.ConfirmToken, error)
}

type mySql struct {
	conn *sql.DB
}

func (m *mySql) ConfirmToken(userId string) (*tokenModels.ConfirmToken, error) {
	stmt := fmt.Sprintf(`SELECT access_token, refresh_token FROM users WHERE user_id = "%s"`, userId)
	row, err := m.conn.Query(stmt)
	if err != nil {
		return nil, err
	}

	var token tokenModels.ConfirmToken

	row.Next()
	if err = row.Scan(&token.AccessToken, &token.RefreshToken); err != nil {
		return nil, err
	}

	return &token, nil
}

func NewMySqlTokenRepo(conn *sql.DB) TokenRepo {
	return &mySql{conn: conn}
}
