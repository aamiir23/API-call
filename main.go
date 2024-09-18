package main

import (
    "database/sql"
    "log"
    "net/http"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

var db *sql.DB

func initDB() {
    var err error
    connStr := "user=username dbname=instagram_db sslmode=disable" // Update with your credentials
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    initDB()
    defer db.Close()

    r := mux.NewRouter()

    // User routes
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users/{id}", getUser).Methods("GET")

    // Post routes
    r.HandleFunc("/posts", createPost).Methods("POST")
    r.HandleFunc("/posts/{id}", getPost).Methods("GET")
    r.HandleFunc("/posts/users/{userId}", listUserPosts).Methods("GET")

    // Start the server
    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}