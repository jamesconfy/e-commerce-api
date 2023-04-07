package repo

import (
	"database/sql"
	"e-commerce/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepo interface {
	ExistsEmail(email string) bool
	ExistsId(userId string) bool
	Add(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId string) (*models.User, error)
	// CreateToken(token *userModels.ResetPasswordRes) error
	// ValidateToken(userId, tokenId string) (*userModels.ValidateTokenRes, error)
	// ChangePassword(userId, newPassword string) error
}

type userSql struct {
	conn *sql.DB
}

func (u *userSql) ExistsEmail(email string) bool {
	var userId string

	query := `SELECT id FROM users WHERE email = ?`

	err := u.conn.QueryRow(query, email).Scan(&userId)

	return err != sql.ErrNoRows
}

func (u *userSql) ExistsId(userId string) bool {
	var email string

	query := `SELECT email FROM users WHERE id = ?`

	err := u.conn.QueryRow(query, userId).Scan(&email)

	return err != sql.ErrNoRows
}

func (u *userSql) Add(user *models.User) (*models.User, error) {
	query := `INSERT INTO users(id, first_name, last_name, email, phone_number, password) 
				VALUES(?, ?, ?, ?, ?, ?)`

	_, err := u.conn.Exec(query,
		user.Id, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password)
	if err != nil {
		return nil, err
	}

	return u.GetById(user.Id)
}

func (u *userSql) GetByEmail(email string) (*models.User, error) {
	var user models.User

	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE email = ?`

	err := u.conn.QueryRow(query, email).Scan(
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
	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE id = ?`

	err := u.conn.QueryRow(query, userId).Scan(
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
