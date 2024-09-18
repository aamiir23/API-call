package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"instagram-api/db"
	"instagram-api/routes"
)

func TestCreateUser(t *testing.T) {
	router := mux.NewRouter()
	db.InitDB("user=username dbname=instagram_db sslmode=disable") // Update with your credentials
	defer db.CloseDB()

	routes.SetDB(db.DB)
	router.HandleFunc("/users", routes.CreateUser).Methods("POST")

	user := map[string]string{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "password123",
	}
	jsonUser, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code %v, got %v", http.StatusCreated, res.Code)
	}
}

func TestGetUser(t *testing.T) {
	router := mux.NewRouter()
	db.InitDB("user=username dbname=instagram_db sslmode=disable") // Update with your credentials
	defer db.CloseDB()

	routes.SetDB(db.DB)
	router.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")

	req, _ := http.NewRequest("GET", "/users/1", nil) // Assuming user with ID 1 exists

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, res.Code)
	}
}