package repo

import (
	"database/sql"
	"e-commerce/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepo interface {
	// Confirmations
	ExistsEmail(email string) bool
	ExistsId(userId string) bool
	ExistsPhone(phone string) bool

	// Real work
	Add(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId string) (*models.User, error)
	Edit(user *models.User, userId string) (*models.User, error)
	Delete(userId string) error
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

func (u *userSql) ExistsPhone(phone string) bool {
	var id string

	query := `SELECT id FROM users WHERE phone_number = ?`

	err := u.conn.QueryRow(query, phone).Scan(&id)

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

func (u *userSql) Edit(user *models.User, userId string) (*models.User, error) {
	query := `UPDATE users SET email = ?, first_name = ?, last_name = ?, phone_number = ?, date_updated = CURRENT_TIMESTAMP() WHERE id = ?`

	_, err := u.conn.Exec(query, user.Email, user.FirstName, user.LastName, user.PhoneNumber, userId)
	if err != nil {
		return nil, err
	}

	return u.GetById(userId)
}

func (u *userSql) Delete(userId string) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := u.conn.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepo(conn *sql.DB) UserRepo {
	return &userSql{conn: conn}
}
