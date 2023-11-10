package bd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbConn *sql.DB
)

// InitDB initializes the database connection and performs setup.
func InitDB() error {
	// Open a database connection.
	db, err := sql.Open("sqlite3", "./internal/database/forum.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	dbConn = db

	CreateUsersTable(db)
	CreatePostsTable(db)
	CreateLikesTable(db)
	CreateCommentsTable(db)
	CreateCategoriesTable(db)
	CreateSessionsTable(db)
	// CreateLikesCommentTable(db)

	return nil
}

// GetDB returns the active database connection.
func GetDB() *sql.DB {
	return dbConn
}
