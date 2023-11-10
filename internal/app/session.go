package app

import (
	"net/http"
	"time"

	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"

	"github.com/gofrs/uuid"
)

func GetCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func IsConnected(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := GetCookie(w, r)
	if err != nil {
		return false
	}
	session, err := repo.GetSessionByID(bd.GetDB(), cookie)
	if err != nil {
		return false
	} else {
		if !session.Ttl.Add(time.Minute * 13).After(time.Now()) {
			repo.DeleteSession(bd.GetDB(), session.SessionID)
			DeleteCookie(w, r)
			return false
		}
	}

	return err == nil
}

func SetCookie(w http.ResponseWriter, r *http.Request, userID string) {
	sessionID, err := uuid.NewV4()
	if err != nil {
		return
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  sessionID.String(),
		Path:   "/",  // Chemin pour lequel le cookie est valide
		MaxAge: 3660, // Durée de vie en secondes (1 heure dans cet exemple)
	}
	session := model.Session{
		SessionID: sessionID.String(),
		UserID:    userID,
		Ttl:       time.Now(),
	}
	repo.CreateSession(bd.GetDB(), session)
	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",  // Valeur vide
		Path:   "/", // Chemin du cookie
		MaxAge: -1,  // MaxAge négatif pour la suppression
	}
	http.SetCookie(w, cookie)
}
