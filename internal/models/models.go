package model

import (
	"time"
)

// User represents a user in the database
type User struct {
	UserID   string // Primary Key
	Username string
	Email    string
	Photo    string
	Password string // encrypted
	Date     time.Time
}

// Category represents a category in the database
type Category struct {
	CategoryID string // Primary Key
	PostID     string // Primary Key
	Name       string
}

// Post represents a post in the database
type Post struct {
	PostID       string // Primary Key
	UserID       string // Foreign Key, references User entity
	Title        string
	Content      string
	Image        string
	Photo        string
	Name         string
	LikeCount    int
	DislikeCount int
	CommentCount int
	Comment      []Comment
	CreatedAt    time.Time // Timestamp
}

// Comment represents a comment in the database
type Comment struct {
	CommentID  string // Primary Key
	UserID     string // Foreign Key, references User entity
	PostID     string // Foreign Key, references Post entity
	Content    string
	Photo      string
	Name       string
	LikeNbr    int
	DislikeNbr int
}

// Like represents a "Like" or "Dislike" in the database
type Like struct {
	LikeID    string // Primary Key
	UserID    string // Foreign Key, references User entity
	PostID    string // Foreign Key, references Post entity
	CommentID string // Foreign Key, references Comment entity
	Type      int    // 1 for Like, 0 for Dislike
}

type Session struct {
	SessionID string    // Primary Key
	UserID    string    // Foreign Key, references User entity
	Ttl       time.Time // Foreign Key, references User entity
}
