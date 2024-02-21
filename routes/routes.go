package routes

import (
	"api/controllers"
	"api/middleware"

	"github.com/gorilla/mux"
)

func addRoutes(mux *mux.Router) {
	controller := controllers.NewController("user")
	r := NewUserRouter(controller)
	r.Load(mux)

}

func NewServer() *mux.Router {
	mux := mux.NewRouter()
	addRoutes(mux)
	mux.Use(middleware.ApplicationTypeSet)
	return mux
}
