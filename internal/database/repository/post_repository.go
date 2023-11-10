package repo

import (
	"database/sql"
	model "forum/internal/models"
	"forum/internal/tools"
	"time"
)

// Create a new Post
func CreatePost(db *sql.DB, post model.Post) error {
	stmt, err := db.Prepare("INSERT INTO posts (post_id, user_id, title, content, image,likeCount, dislikeCount, commentCount,created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);")

	if err != nil {
		return err
	}
	defer stmt.Close()
	post.CreatedAt = time.Now()
	_, err = stmt.Exec(post.PostID, post.UserID, post.Title, post.Content, post.Image, post.LikeCount, post.DislikeCount, post.CommentCount, post.CreatedAt)

	return err
}

// Read all Posts
func GetPosts(db *sql.DB) ([]model.Post, error) {
	querySQL := `SELECT post_id, title, content, image,likeCount, dislikeCount, commentCount, created_at FROM posts;`
	rows, err := db.Query(querySQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.PostID, &post.Title, &post.Content, &post.Image, &post.LikeCount, &post.DislikeCount, &post.CommentCount, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Read a Post by ID
func GetPostByID(db *sql.DB, postID string) (model.Post, error) {
	querySQL := `SELECT post_id, user_id, title, content, image,likeCount, dislikeCount, commentCount, created_at FROM posts WHERE post_id = ?;`
	var post model.Post
	err := db.QueryRow(querySQL, postID).Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.LikeCount, &post.DislikeCount, &post.CommentCount, &post.CreatedAt)
	if err != nil {
		return model.Post{}, err
	}
	return post, nil
}

// Read all Posts
func GetPostsByUser(db *sql.DB, userID string) ([]model.Post, error) {
	querySQL := `SELECT post_id, user_id, title, content, image,likeCount, dislikeCount, commentCount, created_at FROM posts WHERE user_id= ?;`
	rows, err := db.Query(querySQL, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.LikeCount, &post.DislikeCount, &post.CommentCount, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostsByCategoryAndUser retrieves posts with a specific category for a given user.
func GetPostsByCategoryName(db *sql.DB, categoryName string) ([]model.Post, error) {
	querySQL := `
		SELECT p.post_id, p.user_id, p.title, p.content, p.image, p.likeCount, p.dislikeCount, p.commentCount, p.created_at
		FROM posts p
		JOIN categories c ON (p.post_id = c.post_id)
		WHERE (c.name = ?)
		ORDER BY p.created_at DESC;
	`
	rows, err := db.Query(querySQL, categoryName)
	if err != nil {
		tools.LogErr(err)
		return nil, err
	}
	defer rows.Close()
	var posts []model.Post
	for rows.Next() {
		var p model.Post
		if err := rows.Scan(
			&p.PostID, &p.UserID, &p.Title, &p.Content, &p.Image,
			&p.LikeCount, &p.DislikeCount, &p.CommentCount, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// GetPostsByCategoryAndUser retrieves posts with a specific category for a given user.
func GetPostsByCategoryAndUser(db *sql.DB, categoryName string, userID string) ([]model.Post, error) {
	querySQL := `
		SELECT p.post_id, p.user_id, p.title, p.content, p.image, p.likeCount, p.dislikeCount, p.commentCount, p.created_at
		FROM posts p
		JOIN categories c ON (p.post_id = c.post_id)
		WHERE (c.name = ? AND p.user_id = ?)
		ORDER BY p.created_at DESC;
	`
	rows, err := db.Query(querySQL, categoryName, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []model.Post
	for rows.Next() {
		var p model.Post
		if err := rows.Scan(
			&p.PostID, &p.UserID, &p.Title, &p.Content, &p.Image,
			&p.LikeCount, &p.DislikeCount, &p.CommentCount, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// GetLikedPostsByUser retrieves posts liked by a specific user.
func GetLikedPostsByUser(db *sql.DB, userID string) ([]model.Post, error) {
	querySQL := `
		SELECT p.post_id, p.user_id, p.title, p.content, p.image, p.likeCount, p.dislikeCount, p.commentCount, p.created_at
		FROM posts p
		JOIN likes l ON p.post_id = l.post_id
		WHERE (l.user_id = ? AND l.type =1)
		ORDER BY p.created_at DESC;
	`
	rows, err := db.Query(querySQL, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []model.Post
	for rows.Next() {
		var p model.Post
		if err := rows.Scan(
			&p.PostID, &p.UserID, &p.Title, &p.Content, &p.Image,
			&p.LikeCount, &p.DislikeCount, &p.CommentCount, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// Update Post content
func UpdatePostContent(db *sql.DB, postID string, newContent, newTitle string) error {
	stmt, err := db.Prepare("UPDATE posts SET content = ?, title = ? WHERE post_id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newContent, newTitle, postID)
	return err
}

// Delete a Post by ID
func DeletePost(db *sql.DB, postID string) error {
	deleteSQL := `DELETE FROM posts WHERE post_id = ?;`
	_, err := db.Exec(deleteSQL, postID)
	return err
}
