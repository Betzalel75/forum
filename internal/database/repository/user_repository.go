package repo

import (
	"database/sql"
	model "forum/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

// Create a new user
func CreateUser(db *sql.DB, user model.User) error {
	stmt, err := db.Prepare("INSERT INTO users (user_id, username, email, password, photo) VALUES (?, ?, ?, ?, ?);")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.UserID, user.Username, user.Email, user.Password, user.Photo)

	return err
}

// Read all users
func GetUsers(db *sql.DB) ([]model.User, error) {
	querySQL := `SELECT user_id, username, email, password, photo FROM users;`
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.Password, &u.Photo); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Read a user by ID
func GetUserByID(db *sql.DB, userID string) (model.User, error) {
	querySQL := `SELECT user_id, username, email, password, photo FROM users WHERE user_id = ?;`
	var u model.User
	err := db.QueryRow(querySQL, userID).Scan(&u.UserID, &u.Username, &u.Email, &u.Password, &u.Photo)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

// Read a user by ID
func GetUserByEmail(db *sql.DB, email string) (model.User, error) {
	querySQL := `SELECT user_id, username, password, photo FROM users WHERE email = ?;`
	var u model.User
	err := db.QueryRow(querySQL, email).Scan(&u.UserID, &u.Username, &u.Password, &u.Photo)
	return u, err
}

// Update user email
func UpdateUser(db *sql.DB, userID string, newEmail, newUserName, newPassword, newPhoto string) error {
	stmt, err := db.Prepare("UPDATE users SET email = ?, username = ?, password = ?, photo = ? WHERE user_id = ?;")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newEmail, newUserName, newPassword, newPhoto, userID)

	return err
}

// Delete a user by ID
func DeleteUser(db *sql.DB, userID string) error {
	deleteSQL := `DELETE FROM users WHERE user_id = ?;`
	_, err := db.Exec(deleteSQL, userID)
	return err
}
