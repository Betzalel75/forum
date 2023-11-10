package repo

import (
	"database/sql"

	model "forum/internal/models"
)

// Create a new category
func CreateCategory(db *sql.DB, category model.Category) error {
	insertSQL := `
			INSERT INTO categories (category_id, post_id, name) VALUES (?, ?, ?);`
	_, err := db.Exec(insertSQL, category.CategoryID, category.PostID, category.Name)
	return err
}

// Read all categories
func GetCategories(db *sql.DB) ([]model.Category, error) {
	querySQL := `SELECT category_id,post_id, name FROM categories;`
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.CategoryID, &c.PostID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// Read a category by ID
func GetCategoryByID(db *sql.DB, categoryID string) (model.Category, error) {
	querySQL := `SELECT category_id, post_id, name FROM categories WHERE category_id = ?;`
	var c model.Category
	err := db.QueryRow(querySQL, categoryID).Scan(&c.CategoryID, c.PostID, &c.Name)
	if err != nil {
		return model.Category{}, err
	}
	return c, nil
}

// Read all categories
func GetCategoriesByName(db *sql.DB, name string) ([]model.Category, error) {
	querySQL := `SELECT category_id,post_id, name FROM categories WHERE name = ?;`
	rows, err := db.Query(querySQL, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.CategoryID, &c.PostID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// Update category name
func UpdateCategoryName(db *sql.DB, categoryID string, newName string) error {
	updateSQL := `UPDATE categories SET name = ? WHERE category_id = ?;`
	_, err := db.Exec(updateSQL, newName, categoryID)
	return err
}

// Delete a category by ID
func DeleteCategory(db *sql.DB, categoryID string) error {
	deleteSQL := `DELETE FROM categories WHERE category_id = ?;`
	_, err := db.Exec(deleteSQL, categoryID)
	return err
}
