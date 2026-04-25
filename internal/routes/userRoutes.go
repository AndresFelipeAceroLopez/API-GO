package routes

import (
    "github.com/gorilla/mux"
    "your_module/controllers"
)

func UserRoutes(router *mux.Router) {
    router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
    router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
    router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}package routes
