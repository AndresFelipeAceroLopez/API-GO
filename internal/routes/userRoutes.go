package routes

import (
	"net/http"

	"github.com/AndresFelipeAceroLopez/API-GO/internal/controllers"

	"github.com/gorilla/mux"
)

func UsersRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}