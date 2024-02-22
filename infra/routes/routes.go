package routes

import (
	"api/infra/controllers"
	"api/infra/database"
	"api/infra/middleware"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func addRoutes(mux *mux.Router, pool *pgxpool.Pool) {
	NewUserRouter(controllers.NewUserController(pool)).Load(mux)
	NewAuthRouter(controllers.NewAuthController(pool)).Load(mux)
}

func NewServer(env map[string]string) *mux.Router {
	mux := mux.NewRouter()

	connect := database.NewConnect(env["DATABASE_URL"])

	addRoutes(mux, connect)

	mux.Use(middleware.ApplicationTypeSet)

	return mux
}
