package handler

import (
	"forum/internal/app"
	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"log"
	"net/http"
	"text/template"
)

type Set struct {
	Status  string
	Message string
	User    model.User
}

func setting(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./web/template/setting.html")
	if err != nil {
		log.Panic(err)
	}
	set := Set{}
	session_Id, _ := app.GetCookie(w, r)
	user_Id, err := repo.GetUserIDBySession(bd.GetDB(), session_Id)
	if err != nil {
		log.Print(err)
		// return models.User{}, "impossible de modifier vos informatio", false
	}
	User, err := repo.GetUserByID(bd.GetDB(), user_Id)
	if err != nil {
		log.Print(err)
		// return models.User{}, "impossible de modifier vos informatio", false
	}
	tmpUser := model.User{
		Username: User.Username,
		Email:    User.Email,
		Photo:    User.Photo,
	}
	set.User = tmpUser
	if r.Method == "POST" {
		newUser, message, status := app.UpdateUser(w, r, User)
		if status {
			set = Set{
				Status:  "succes",
				Message: message,
				User:    newUser,
			}
		} else {
			set = Set{
				Status:  "nosucces",
				Message: message,
				User:    tmpUser,
			}
		}

		// fmt.Println(newUser)
	}
	tmpl.Execute(w, set)

}
