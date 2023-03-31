package repo

import (
	"database/sql"
	"e-commerce/internal/models"
	"fmt"
)

type UserRepo interface {
	ExistsEmail(email string) bool
	ExistsId(userId string) bool
	Register(user *models.UserCart, accessToken, refreshToken string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId string) (*models.User, error)
	UpdateTokens(auth *models.Auth) error
	// CreateToken(token *userModels.ResetPasswordRes) error
	// ValidateToken(userId, tokenId string) (*userModels.ValidateTokenRes, error)
	// ChangePassword(userId, newPassword string) error
}

type userSql struct {
	conn *sql.DB
}

func (u *userSql) ExistsEmail(email string) bool {
	var userId string

	stmt := fmt.Sprintf(`SELECT id FROM users WHERE email = '%[1]v'`, email)

	err := u.conn.QueryRow(stmt).Scan(&userId)

	return err != sql.ErrNoRows
}

func (u *userSql) ExistsId(userId string) bool {
	var email string

	stmt := fmt.Sprintf(`SELECT email FROM users WHERE id = '%[1]v'`, userId)

	err := u.conn.QueryRow(stmt).Scan(&email)

	return err != sql.ErrNoRows
}

func (u *userSql) Register(user *models.UserCart, accessToken, refreshToken string) (*models.User, error) {
	query := `INSERT INTO users(id, first_name, last_name, email, phone_number, password, date_created, access_token, refresh_token) VALUES ('%[1]v', '%[2]v', '%[3]v', '%[4]v', '%[5]v', '%[6]v', '%[7]v', '%[8]v', '%[9]v')`

	stmt := fmt.Sprintf(query,
		user.User.Id, user.User.FirstName, user.User.LastName, user.User.Email, user.User.PhoneNumber, user.User.Password, user.User.DateCreated, accessToken, refreshToken)

	_, err := u.conn.Exec(stmt)
	if err != nil {
		return nil, err
	}

	return u.GetById(user.User.Id)
}

func (u *userSql) GetByEmail(email string) (*models.User, error) {
	var user models.User

	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE email = '%s'`

	stmt := fmt.Sprintf(query, email)

	err := u.conn.QueryRow(stmt).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.DateCreated,
		&user.DateUpdated,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userSql) GetById(userId string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_created FROM users WHERE id = '%[1]s'`

	stmt := fmt.Sprintf(query, userId)

	err := u.conn.QueryRow(stmt).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.DateCreated,
		&user.DateUpdated,
	)

	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return &models.User{}, err
		// }
		return nil, err
	}

	return &user, nil
}

func (u *userSql) UpdateTokens(auth *models.Auth) error {
	query := `UPDATE users SET access_token = '%[1]v', refresh_token = '%[2]v', date_updated = '%[3]v' WHERE id = '%[4]v'`

	stmt := fmt.Sprintf(query, auth.AccessToken, auth.RefreshToken, auth.DateUpdated, auth.UserId)

	_, err := u.conn.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

// func (m *userSql) CreateToken(token *userModels.ResetPasswordRes) error {
// 	ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*30)
// 	defer cancelFunc()

// 	tx, err := m.conn.BeginTx(ctx, nil)
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 			return
// 		}
// 		tx.Commit()
// 		// return
// 	}()

// 	deleteQuery := fmt.Sprintf(`DELETE FROM token WHERE user_id = '%v'`, token.UserId)
// 	_, err = tx.ExecContext(ctx, deleteQuery)
// 	if err != nil {
// 		return err
// 	}

// 	addQuery := fmt.Sprintf(`INSERT INTO token(user_id, token_id, token, expiry) VALUES('%v','%v','%v','%v')`, token.UserId, token.TokenId, token.Token, token.Expiry)
// 	_, err = tx.ExecContext(ctx, addQuery)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m *userSql) ValidateToken(userId, tokenId string) (*userModels.ValidateTokenRes, error) {
// 	stmt := fmt.Sprintf(`SELECT user_id, token_id, token, expiry
// 						FROM token
// 						WHERE user_id = '%v' AND token = '%v'`, userId, tokenId)

// 	var token userModels.ValidateTokenRes

// 	err := m.conn.QueryRow(stmt).Scan(
// 		&token.UserId,
// 		&token.TokenId,
// 		&token.Token,
// 		&token.Expiry,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &token, nil
// }

// func (m *userSql) ChangePassword(user_id, newPassword string) error {
// 	query := fmt.Sprintf(`UPDATE users SET password = '%v' WHERE user_id = '%v'`, newPassword, user_id)
// 	_, err := m.conn.Exec(query)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	return nil
// }

func NewUserRepo(conn *sql.DB) UserRepo {
	return &userSql{conn: conn}
}
