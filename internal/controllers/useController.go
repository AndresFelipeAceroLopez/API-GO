package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AndresFelipeAceroLopez/API-GO/internal/models" // Asegúrate que el import coincida con tu go.mod
)

// GET: Obtener todos los usuarios
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    models.Users,
		"total":   len(models.Users),
	})
}

// GET: Obtener usuario por ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extraer ID de la URL (asumiendo formato /users/1)
	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, _ := strconv.Atoi(idStr)

	for _, user := range models.Users {
		if user.ID == id {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"data":    user,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   fmt.Sprintf("Usuario con ID %d no encontrado", id),
	})
}

// POST: Crear un nuevo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	// Validaciones básicas
	if err != nil || newUser.Name == "" || newUser.Email == "" || newUser.Age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Los campos name, email, age son obligatorios",
		})
		return
	}

	// Generar ID (Lógica simple como la de tu TS)
	maxID := 0
	for _, u := range models.Users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	newUser.ID = maxID + 1
	models.Users = append(models.Users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Usuario creado exitosamente",
		"data":    newUser,
	})
}

// PUT: Actualizar usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, _ := strconv.Atoi(idStr)

	var updateData models.User
	json.NewDecoder(r.Body).Decode(&updateData)

	for i, user := range models.Users {
		if user.ID == id {
			if updateData.Name != "" {
				models.Users[i].Name = updateData.Name
			}
			if updateData.Email != "" {
				models.Users[i].Email = updateData.Email
			}
			if updateData.Age != 0 {
				models.Users[i].Age = updateData.Age
			}

			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "Usuario actualizado exitosamente",
				"data":    models.Users[i],
			})
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   fmt.Sprintf("Usuario con ID %d no encontrado", id),
	})
}

// DELETE: Eliminar Usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	idStr := parts[len(parts)-1]
	id, _ := strconv.Atoi(idStr)

	for i, user := range models.Users {
		if user.ID == id {
			deletedUser := models.Users[i]
			models.Users = append(models.Users[:i], models.Users[i+1:]...)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "Usuario eliminado exitosamente",
				"data":    deletedUser,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   fmt.Sprintf("Usuario con ID %d no encontrado", id),
	})
}
