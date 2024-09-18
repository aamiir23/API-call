package routes

import (
	"encoding/json"
	"instagram-api/db"
	"instagram-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

// var db *sql.DB

// func SetDB(database *sql.DB) {
// 	db = database
// }

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Store post in the database (make sure to include user_id when inserting)
	_, err := db.Exec("INSERT INTO posts (user_id, caption, image_url) VALUES ($1, $2, $3)", post.UserID, post.Caption, post.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post models.Post
	err := db.QueryRow("SELECT id, user_id, caption, image_url FROM posts WHERE id = $1", id).Scan(&post.ID, &post.UserID, &post.Caption, &post.ImageURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func listUserPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	rows, err := db.Query("SELECT id, user_id, caption, image_url FROM posts WHERE user_id = $1", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Caption, &post.ImageURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}
