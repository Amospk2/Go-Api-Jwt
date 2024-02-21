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

	rows, err := db.pool.Query(context.Background(), "select * from public.users")

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
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}
