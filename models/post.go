package models

type Post struct {
    ID              int     `json:"id"`
    UserID          int     `json:"user_id"` // Foreign key reference to the user
    Caption         string  `json:"caption"`
    ImageURL        string  `json:"image_url"`
    PostedTimestamp string  `json:"posted_timestamp"`
}