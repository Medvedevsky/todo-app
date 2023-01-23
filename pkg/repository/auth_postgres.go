package repository

import (
	"fmt"

	"github.com/Medvedevsky/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a AuthPostgres) CreateUser(user todo.User) (int, error) {
	var user_id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING user_id", usersTable)

	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&user_id); err != nil {
		return 0, err
	}

	return user_id, nil
}

func (a AuthPostgres) GetUser(login string, password string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT user_id FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := a.db.Get(&user, query, login, password)

	return user, err
}
