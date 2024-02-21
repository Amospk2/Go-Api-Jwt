package controllers

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
	user string
}

func (c *UserController) GetUsers() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"users": []string{c.user}})
		},
	)
}

func NewController(user string) *UserController {
	return &UserController{
		user: user,
	}
}
