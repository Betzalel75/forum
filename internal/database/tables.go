package bd

import (
	"database/sql"
	"log"
)

// Create the Categories table
func CreateCategoriesTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS categories (
            category_id TEXT PRIMARY KEY UNIQUE,
            post_id TEXT,
            name TEXT
        );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// Create the Comments table
func CreateCommentsTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS comments (
            comment_id TEXT PRIMARY KEY UNIQUE,
            user_id TEXT NOT NULL,
            post_id TEXT NOT NULL,
            content TEXT
        );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// Create the Likes table
func CreateLikesTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS likes (
            like_id TEXT PRIMARY KEY UNIQUE,
            user_id TEXT,
            post_id TEXT,
            comment_id TEXT,
            type INTEGER
        );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// Create the Posts table
func CreatePostsTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS posts (
            post_id TEXT PRIMARY KEY UNIQUE,
            user_id TEXT,
            title TEXT,
            content TEXT,
            image TEXT,
            likeCount INTEGER,
            dislikeCount INTEGER,
            commentCount INTEGER,
            created_at DATETIME
        );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// Create the Users table
func CreateUsersTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            user_id TEXT PRIMARY KEY UNIQUE,
            username TEXT,
            email TEXT,
            password TEXT,
            photo TEXT,
            date DATETIME,
            level INTEGER
        );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// Create the Users Session
func CreateSessionsTable(db *sql.DB) {
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS sessions (
            session_id TEXT PRIMARY KEY UNIQUE,
            user_id TEXT,
            ttl DATETIME
        );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
