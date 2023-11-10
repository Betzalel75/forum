package app

import (
	"net/http"
	"os"

	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"forum/internal/tools"

	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(w http.ResponseWriter, r *http.Request, User model.User) (model.User, string, bool) {
	form := []string{r.FormValue("username"), r.FormValue("email"), r.FormValue("pass")}
	ok := false
	if form[1] != "" {
		return User, "modify unauthorized email", false
	}
	if len(r.FormValue("oldPassword")) != 0 {
		form[2] = r.FormValue("oldPassword")
		ok = true
	}
	if bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(form[2])) == nil {
		if ok && len(r.FormValue("newPassword")) < 4 {
			return User, "new passeword is ", false
		} else if ok && len(r.FormValue("newPassword")) > 3 {
			form[2] = r.FormValue("newPassword")
		}
		pwd, err := passwordCrypt(form[2])
		if err != nil {
			return User, "Unable to modify your new data", false
		}
		image := User.Photo
		if User.Photo != "default.jpg" {
			os.Remove("./web/static/media/" + User.Photo)
		}
		if file, _, _ := r.FormFile("photo"); file != nil {
			img, err := tools.UploadFile(w, r, "photo", User.UserID)
			if err != nil {
				return User, img, false
			}
			if img != "" {
				image = img
			}
		}
		repo.UpdateUser(bd.GetDB(), User.UserID, User.Email, form[0], pwd, image)
		user := User
		info, err := repo.GetUserByID(bd.GetDB(), User.UserID)
		if err != nil {
			return User, "Unable to display your new data", false
		}
		user.Username, user.Email, user.Photo = info.Username, info.Email, info.Photo
		return user, "successfully modified", true
	}
	return User, "passwords incorrect", false
}
