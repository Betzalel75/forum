package app

import (
	"database/sql"
	"net/http"
	"strings"

	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"forum/internal/tools"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// login
func Authentication(w http.ResponseWriter, r *http.Request, email, password string) string {
	email = strings.TrimSpace(email)
	user, exist := repo.GetUserByEmail(bd.GetDB(), email)
	if exist == sql.ErrNoRows {
		return "Email incorecte"
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil {
		session, err := repo.GetSessionByUserI(bd.GetDB(), user.UserID)
		if err == nil {
			repo.DeleteSession(bd.GetDB(), session.SessionID)
		}
		SetCookie(w, r, user.UserID)
		http.Redirect(w, r, "/forum?name=all", http.StatusPermanentRedirect)
	}
	return "Passwords incorrect"
}

// sign Un
func Registration(w http.ResponseWriter, r *http.Request, username, email, password string) string {
	email = strings.TrimSpace(email)
	if username == "" || email == "" {
		return "all fields are required"
	}
	valid := strings.Split(email, "@")
	if len(valid) != 2 || valid[0] == "" || valid[1] == "" {
		return "incorrect email format"
	}
	if len(password) < 4 {
		return "password too short"
	}
	_, exist := repo.GetUserByEmail(bd.GetDB(), email)
	if exist == sql.ErrNoRows {
		pwd, err := passwordCrypt(password)
		if err != nil {
			tools.LogErr(err)
			return ""
		}
		id, err := uuid.NewV4()
		if err != nil {
			tools.LogErr(err)
			return ""
		}
		user := model.User{
			UserID:   id.String(),
			Username: username,
			Email:    email,
			Password: pwd,
			Photo:    "defautl.jpg",
		}
		repo.CreateUser(bd.GetDB(), user)
		SetCookie(w, r, user.UserID)
		http.Redirect(w, r, "/forum?name=all", http.StatusPermanentRedirect)
		return ""
	}

	return "email already exists"
}

func passwordCrypt(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		tools.LogErr(err)
	}
	return string(pwd), nil
}
