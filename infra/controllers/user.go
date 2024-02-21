package controllers

import (
	"api/domain/user"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController struct {
	repository *user.UserRepository
}

func (c *UserController) GetUsers() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			users, err := c.repository.Get()
			if err != nil {
				log.Fatal(err)
			}
			json.NewEncoder(w).Encode(map[string]any{"users": users})
		},
	)
}

func NewController(pool *pgxpool.Pool) *UserController {
	return &UserController{
		repository: user.NewUserRepository(pool),
	}
}
