package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func (db *UserRepository) Get() ([]Users, error) {

	users := make([]Users, 0)

	rows, err := db.pool.Query(context.Background(),
		"select id, name, email, age, password from public.users",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user Users

		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.Password,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (db *UserRepository) GetById(id string) (Users, error) {

	var user Users

	err := db.pool.QueryRow(
		context.Background(),
		"select id, name, email, age, password from public.users where id=$1", id,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.Password,
	)

	if err != nil {
		return Users{}, err
	}

	return user, nil
}

func (db *UserRepository) GetByEmail(email string) (Users, error) {

	var user Users

	err := db.pool.QueryRow(
		context.Background(),
		"select id, name, email, age, password from public.users where email=$1", email,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.Password,
	)

	if err != nil {
		return Users{}, err
	}

	return user, nil
}

func (db *UserRepository) Update(user Users) error {
	_, err := db.pool.Exec(
		context.Background(),
		"UPDATE USERS SET name = $1, email = $2, age = $3, password = $4 WHERE id = $5",
		user.Name, user.Email, user.Age, user.Password, user.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *UserRepository) Create(user Users) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO USERS VALUES($1,$2,$3,$4, $5)",
		user.Id, user.Name, user.Email, user.Age, user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db *UserRepository) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM USERS WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
