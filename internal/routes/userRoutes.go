package routes

import (
	"net/http"

	"github.com/AndresFelipeAceroLopez/API-GO/internal/controllers"

	"github.com/gorilla/mux"
)

func UsersRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}
