package repo

import (
	"database/sql"
	"e-commerce/internal/models"

	_ "github.com/lib/pq"
)

type UserRepo interface {
	// Confirmations
	ExistsEmail(email string) (bool, error)
	ExistsId(email string) bool
	ExistsPhone(phone string) (bool, error)

	// Real work
	Add(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId string) (*models.User, error)
	GetAll(page int) ([]*models.User, error)
	Edit(user *models.User, userId string) (*models.User, error)
	Delete(userId string) error
}

type userSql struct {
	conn *sql.DB
}

func (u *userSql) ExistsEmail(email string) (bool, error) {
	var userId string

	query := `SELECT id FROM users WHERE email = $1;`

	err := u.conn.QueryRow(query, email).Scan(&userId)

	if err != nil {
		if err == sql.ErrNoRows {
			// Email does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Email already exists
	return true, nil
}

func (u *userSql) ExistsId(userId string) bool {
	var email string

	query := `SELECT email FROM users WHERE id = $1;`

	err := u.conn.QueryRow(query, userId).Scan(&email)

	return err != sql.ErrNoRows
}

func (u *userSql) ExistsPhone(phone string) (bool, error) {
	var id string

	query := `SELECT id FROM users WHERE phone_number = $1`

	err := u.conn.QueryRow(query, phone).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			// Phone does not exist
			return false, nil
		}
		// An error occurred while executing the query
		return true, err
	}

	// Phone already exists
	return true, nil
}

func (u *userSql) Add(user *models.User) (usr *models.User, err error) {
	usr = new(models.User)

	query := `INSERT INTO users(first_name, last_name, email, phone_number, password) VALUES($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, email, phone_number, date_created, date_updated`

	err = u.conn.QueryRow(query, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.PhoneNumber, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (u *userSql) GetByEmail(email string) (*models.User, error) {
	var user models.User

	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE email = $1`

	err := u.conn.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.DateCreated, &user.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userSql) GetById(userId string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated FROM users WHERE id = $1`

	err := u.conn.QueryRow(query, userId).Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.DateCreated, &user.DateUpdated)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userSql) GetAll(page int) ([]*models.User, error) {
	limit := 20
	offset := limit * (page - 1)

	query := `SELECT id, email, password, first_name, last_name, phone_number, date_created, date_updated
			FROM users LIMIT $1 OFFSET $2;`

	rows, err := u.conn.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.DateCreated, &user.DateUpdated)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (u *userSql) Edit(user *models.User, userId string) (usr *models.User, err error) {
	usr = new(models.User)

	query := `UPDATE users SET email = $1, first_name = $2, last_name = $3, phone_number = $4, date_updated = CURRENT_TIMESTAMP WHERE id = $5 RETURNING id, first_name, last_name, email, phone_number, date_created, date_updated`

	err = u.conn.QueryRow(query, user.Email, user.FirstName, user.LastName, user.PhoneNumber, userId).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.PhoneNumber, &usr.DateCreated, &usr.DateUpdated)
	if err != nil {
		return
	}

	return
}

func (u *userSql) Delete(userId string) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := u.conn.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepo(conn *sql.DB) UserRepo {
	return &userSql{conn: conn}
}
