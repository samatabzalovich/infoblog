package models

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) (error, int) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err, 0
	}

	stmt := `INSERT INTO users(name, email, hashed_password, created, status) VALUES($1, $2, $3, current_timestamp, 'user') returning id`

	var id int
	err = m.DB.QueryRow(context.Background(), stmt, name, email, string(hashedPassword)).Scan(&id)
	if err != nil {
		var pgxSQLError *mysql.MySQLError
		if errors.As(err, pgxSQLError) {
			if pgxSQLError.Number == 1062 && strings.Contains(pgxSQLError.Message, "users_uc_email") {
				return ErrDuplicateEmail, 0
			}
		}
		return err, 0
	}
	return nil, id
}
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `select id, hashed_password from users where email = $1`

	err := m.DB.QueryRow(context.Background(), stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil

}
func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool
	stmt := `SELECT EXISTS(SELECT true FROM users WHERE id = $1)`
	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&exists)
	return exists, err

}
