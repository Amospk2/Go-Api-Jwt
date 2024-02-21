package routes

import (
	"api/infra/controllers"

	"github.com/gorilla/mux"
)

type UserRouter struct {
	controller *controllers.UserController
}

func (p *UserRouter) Load(mux *mux.Router) {
	mux.HandleFunc("/users", p.controller.GetUsers())
}

func NewUserRouter(
	controller *controllers.UserController,
) *UserRouter {
	return &UserRouter{
		controller: controller,
	}
}
