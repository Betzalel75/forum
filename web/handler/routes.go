package handler

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"forum/internal/app"
	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"forum/internal/tools"
)

type Erreur struct {
	Code    int
	Message string
}

type loginfo struct {
	Status  string
	Message string
}

type PageInfo struct {
	Name      string
	Photo     string
	Posts     []model.Post
	AllUser   []Connection
	Connected bool
}

// declancherles erreurs
func StatusError(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.WriteHeader(code)
	err := Erreur{code, message}
	RenderTemplate(w, err)
}

func RenderTemplate(w http.ResponseWriter, er Erreur) {
	tmpl, err := template.ParseFiles("./web/template/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	tmpl.ExecuteTemplate(w, "error", er)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/template/index.html")
	if err != nil {
		tools.LogErr(err)
		return
	}
	err = tmpl.ExecuteTemplate(w, "index", "")
	if err != nil {
		tools.LogErr(err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/template/login.html")
	if err != nil {
		tools.LogErr(err)
		return
	}
	message, status := "", ""

	if r.Method == "POST" {
		if r.FormValue("signup") == "SIGN UP" {
			username := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")
			message = app.Registration(w, r, username, email, password)
			status = "nosucces"
		} else if r.FormValue("signin") == "SIGN IN" {
			email := r.FormValue("email")
			password := r.FormValue("password")
			message = app.Authentication(w, r, email, password)
			status = "nosucces"
		} else {
			// log.Panicln("formulaire inconnue")
			message = "unknown form"
			status = "nosucces"
			return
		}
	}
	infolog := loginfo{
		Status:  status,
		Message: message,
	}
	tmpl.ExecuteTemplate(w, "login", infolog)
	if err != nil {
		tools.LogErr(err)
		return
	}
}

func forum(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./web/template/forum.html")
	if err != nil {
		tools.LogErr(err)
		return
	}
	filtre := r.URL.Query().Get("name")
	if filtre != "Event" && filtre != "General" && filtre != "Issue" {
		filtre = "all"
	}

	posts, err := repo.GetPostsByCategoryName(bd.GetDB(), filtre)
	if err != nil {
		tools.LogErr(err)
		return
	}

	tmpl.Execute(w, RecoverPageInfo(w, r, posts))
	if err != nil {
		tools.LogErr(err)
		return
	}
}

func profil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	if app.IsConnected(w, r) {
		tmpl, err := template.ParseFiles("./web/template/profil.html")
		if err != nil {
			tools.LogErr(err)
			return
		}
		cookie, err := app.GetCookie(w, r)
		if err != nil {
			return
		}

		userID, err := repo.GetUserIDBySession(bd.GetDB(), cookie)
		if err != nil {
			return
		}

		filtre := r.URL.Query().Get("name")
		if filtre != "Event" && filtre != "General" && filtre != "Issue" && filtre != "Liked" {
			filtre = "all"
		}

		posts, err := repo.GetPostsByCategoryAndUser(bd.GetDB(), filtre, userID)
		if filtre == "Liked" {
			posts, err = repo.GetLikedPostsByUser(bd.GetDB(), userID)
		}
		if err != nil {
			return
		}

		tmpl.ExecuteTemplate(w, "profil", RecoverPageInfo(w, r, posts))
		if err != nil {
			tools.LogErr(err)
			return
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := app.GetCookie(w, r)
	repo.DeleteSession(bd.GetDB(), cookie)
	app.DeleteCookie(w, r)
	http.Redirect(w, r, "/forum?name=all", http.StatusSeeOther)
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		app.AddPost(w, r)
		http.Redirect(w, r, "/profil?name=all", http.StatusPermanentRedirect)
	} else {
		http.Redirect(w, r, "/login?name=all", http.StatusPermanentRedirect)
	}
}

func like(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	if app.IsConnected(w, r) {
		var request app.FeedbackRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			StatusError(w, r, http.StatusBadRequest, "bad request")
			return
		}

		cookie, err := app.GetCookie(w, r)
		if err != nil {
			tools.LogErr(err)
			return
		}

		// Perform action (like or dislike) and update counts
		response := app.PerformAction(cookie, request.PostID, r.URL.Query().Get("name"), request.Action)
		// Return the updated counts as JSON response

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			tools.LogErr(err)
			return
		}
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	if !app.IsConnected(w, r) {
		http.Redirect(w, r, "/forum", http.StatusPermanentRedirect)
	}
	if r.Method == "POST" {
		postID := r.FormValue("postID")
		comment := r.FormValue("comment")
		submit := r.FormValue("add")
		cookie, err := app.GetCookie(w, r)
		if err != nil {
			tools.LogErr(err)
			return
		}
		app.AddComment(postID, comment, submit, cookie)
		url := r.URL.Query().Get("page")
		if url == "forum" || url == "profil" {
			http.Redirect(w, r, "/"+url, http.StatusPermanentRedirect)
		} else {
			tools.LogErr(err)
			return
		}
	} else {
		StatusError(w, r, http.StatusBadRequest, "bad request")
	}
}

type Connection struct {
	Name   string
	Photo  string
	Status string
}

// function to retrieve and sort all users and and check users connected and deconnected to fill the Connection structure table if the user is connected Ceonnection.Status ="" or if the user is deconnected Ceonnection.Status ="offline"
func GetConnection() []Connection {
	var connections []Connection
	users, err := repo.GetUsers(bd.GetDB())
	if err != nil {
		tools.LogErr(err)
	}
	for _, user := range users {
		connect := Connection{}
		connect = Connection{
			Name:   user.Username,
			Photo:  user.Photo,
			Status: "offline",
		}
		_, err := repo.GetSessionByUserI(bd.GetDB(), user.UserID)
		if err == nil {
			connect.Status = ""
		}
		connections = append(connections, connect)
	}
	return connections
}

func RecoverPageInfo(w http.ResponseWriter, r *http.Request, posts []model.Post) PageInfo {
	var connected = true
	cookie, err := app.GetCookie(w, r)
	if err != nil {
		// return PageInfo{}
		connected = false
	}
	userID, err := repo.GetUserIDBySession(bd.GetDB(), cookie)
	if err != nil {
		if err == sql.ErrNoRows {
			connected = false
		}
	}
	user, err := repo.GetUserByID(bd.GetDB(), userID)
	if err != nil {
		connected = false
	}
	pageinfo := PageInfo{
		Name:      user.Username,
		Photo:     user.Photo,
		Posts:     app.InitPostAndCommentLikeAndDislike(posts),
		AllUser:   GetConnection(),
		Connected: connected,
	}
	return pageinfo
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		root(w, r)
	case "/forum":
		forum(w, r)
	case "/login":
		if app.IsConnected(w, r) {
			http.Redirect(w, r, "/profil", http.StatusPermanentRedirect)
		}
		login(w, r)
	case "/post":
		post(w, r)
	case "/comment":
		comment(w, r)
	case "/like":
		like(w, r)
	case "/profil":
		profil(w, r)
	case "/settings":
		if !app.IsConnected(w, r) {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		}
		setting(w, r)
	case "/logout":
		logout(w, r)
	default:
		fmt.Println(http.StatusNotFound)
		StatusError(w, r, http.StatusNotFound, "Not found")
	}
}
