package main

import (
	"log"
	"net/http"
	"os"

	"API-GO/internal/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	// ruta raíz
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"message": "Bienvenido a la API Rest con Go!",
			"version": "1.0.0"
		}`))
	})

	// users
	mux.Handle("/api/users/", http.StripPrefix("/api/users", routes.UsersRouter()))

	// middleware CORS
	handler := corsMiddleware(mux)

	log.Println("Servidor en http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// ---- cors ----
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
