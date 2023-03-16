package userRepo

import (
	"context"
	"database/sql"
	"e-commerce/internal/models/userModels"
	"fmt"
	"time"
)

type UserRepo interface {
	RegisterUser(req *userModels.CreateUserReq) error
	GetByEmail(email string) (*userModels.GetByEmailRes, error)
	GetById(email string) (*userModels.GetByIdRes, error)
	UpdateTokens(req *userModels.UpdateTokens) error
	CreateToken(token *userModels.ResetPasswordRes) error
	ValidateToken(userId, tokenId string) (*userModels.ValidateTokenRes, error)
	ChangePassword(userId, newPassword string) error
}

type mySql struct {
	conn *sql.DB
}

func (m *mySql) RegisterUser(req *userModels.CreateUserReq) error {
	stmt := fmt.Sprintf(`INSERT INTO users(
                   user_id,
                   first_name,
                   last_name,
                   email,
                   phone_number,
                   password,
				   date_created,
				   access_token,
				   refresh_token
                   ) VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')`,
		req.UserId, req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Password, req.DateCreated, req.AccessToken, req.RefreshToken)

	_, err := m.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

func (m *mySql) GetByEmail(email string) (*userModels.GetByEmailRes, error) {
	ctx := context.Background()
	stmt := fmt.Sprintf(`
		SELECT user_id, email, password, first_name, last_name, phone_number, date_created
		FROM users
		WHERE email = '%s'
	`, email)
	var user userModels.GetByEmailRes

	err := m.conn.QueryRowContext(ctx, stmt).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.DateCreated,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *mySql) GetById(userId string) (*userModels.GetByIdRes, error) {
	ctx := context.Background()
	stmt := fmt.Sprintf(`
		SELECT user_id, email, password, first_name, last_name, phone_number, date_created
		FROM users
		WHERE user_id = '%s'
	`, userId)
	var user userModels.GetByIdRes

	err := m.conn.QueryRowContext(ctx, stmt).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.DateCreated,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *mySql) UpdateTokens(req *userModels.UpdateTokens) error {
	query := fmt.Sprintf(`UPDATE users 
							SET access_token = '%v', refresh_token = '%v', date_updated = '%v' 
							WHERE user_id = '%v'`, req.AccessToken, req.RefreshToken, req.DateUpdated, req.UserId)

	_, err := m.conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (m *mySql) CreateToken(token *userModels.ResetPasswordRes) error {
	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancelFunc()

	tx, err := m.conn.BeginTx(ctx, nil)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		// return
	}()

	deleteQuery := fmt.Sprintf(`DELETE FROM token WHERE user_id = '%v'`, token.UserId)
	_, err = tx.ExecContext(ctx, deleteQuery)
	if err != nil {
		return err
	}

	addQuery := fmt.Sprintf(`INSERT INTO token(user_id, token_id, token, expiry) VALUES('%v','%v','%v','%v')`, token.UserId, token.TokenId, token.Token, token.Expiry)
	_, err = tx.ExecContext(ctx, addQuery)
	if err != nil {
		return err
	}

	return nil
}

func (m *mySql) ValidateToken(userId, tokenId string) (*userModels.ValidateTokenRes, error) {
	stmt := fmt.Sprintf(`SELECT user_id, token_id, token, expiry
						FROM token
						WHERE user_id = '%v' AND token = '%v'`, userId, tokenId)

	var token userModels.ValidateTokenRes

	err := m.conn.QueryRow(stmt).Scan(
		&token.UserId,
		&token.TokenId,
		&token.Token,
		&token.Expiry,
	)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (m *mySql) ChangePassword(user_id, newPassword string) error {
	query := fmt.Sprintf(`UPDATE users SET password = '%v' WHERE user_id = '%v'`, newPassword, user_id)
	_, err := m.conn.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewMySqlUserRepo(conn *sql.DB) UserRepo {
	return &mySql{conn: conn}
}
