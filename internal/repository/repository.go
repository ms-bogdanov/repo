package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"repo/config"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(cfg config.PgConfig) *UserStorage {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil
	}

	return &UserStorage{
		db: db,
	}
}

type Conditions struct {
}

type UserRepository interface {
	Create(ctx context.Context, user User) error
	GetById(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, c Conditions) ([]User, error)
}

const (
	queryCreateUser = `insert into users("name", age) values ($1, $2)`

	querySelectUser = `select id, name, age from users where id = $1`

	queryUpdateUser = `update users set "name"= $2, age = $3 where id=$1;`

	queryDeleteUser = `delete from users where id = $1;`

	queryListUser = `select id, "name", age from users`
)

func (us *UserStorage) Create(ctx context.Context, user User) error {
	_, err := us.db.ExecContext(ctx, queryCreateUser, user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserStorage) GetById(ctx context.Context, id int64) (User, error) {
	var user User

	err := us.db.QueryRowContext(ctx, querySelectUser, id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (us *UserStorage) Update(ctx context.Context, user User) error {
	_, err := us.db.ExecContext(ctx, queryUpdateUser, user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserStorage) Delete(ctx context.Context, id int64) error {
	_, err := us.db.ExecContext(ctx, queryDeleteUser, id)
	if err != nil {
		return err
	}

	return err
}

func (us *UserStorage) List(ctx context.Context, c Conditions) ([]User, error) {
	rows, err := us.db.QueryContext(ctx, queryListUser)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
