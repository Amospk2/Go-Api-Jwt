package routes

import (
	"api/infra/controllers"
	"api/infra/database"
	"api/infra/middleware"

	"github.com/gorilla/mux"
)

func addRoutes(mux *mux.Router, controller *controllers.UserController) {
	r := NewUserRouter(controller)
	r.Load(mux)
}

func NewServer(env map[string]string) *mux.Router {
	mux := mux.NewRouter()

	connect := database.NewConnect(env["DATABASE_URL"])

	controller := controllers.NewController(connect)
	addRoutes(mux, controller)

	mux.Use(middleware.ApplicationTypeSet)

	return mux
}
