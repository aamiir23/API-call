package routes

import (
	"encoding/json"
	"instagram-api/models"
	"net/http"

	"instagram-api/db"

	"github.com/gorilla/mux"
)

// var db *sql.DB

// func SetDB(database *sql.DB) {
// 	db = database
// }

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store user in the database (ensure password is hashed before storing)
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
