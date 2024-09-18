package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"instagram-api/db"
	"instagram-api/routes"

	"github.com/gorilla/mux"
)

func TestCreatePost(t *testing.T) {
	router := mux.NewRouter()
	db.InitDB("user=username dbname=instagram_db sslmode=disable") // Update with your credentials
	defer db.CloseDB()

	routes.SetDB(db.DB)
	router.HandleFunc("/posts", routes.CreatePost).Methods("POST")

	post := map[string]interface{}{
		"user_id":   1,
		"caption":   "This is a test post",
		"image_url": "http://example.com/image.jpg",
	}
	jsonPost, _ := json.Marshal(post)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonPost))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code %v, got %v", http.StatusCreated, res.Code)
	}
}

func TestGetPost(t *testing.T) {
	router := mux.NewRouter()
	db.InitDB("user=username dbname=instagram_db sslmode=disable") // Update with your credentials
	defer db.CloseDB()

	routes.SetDB(db.DB)
	router.HandleFunc("/posts/{id}", routes.GetPost).Methods("GET")

	req, _ := http.NewRequest("GET", "/posts/1", nil) // Assuming post with ID 1 exists

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, res.Code)
	}
}
