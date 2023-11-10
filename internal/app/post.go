package app

import (
	"fmt"
	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"forum/internal/tools"
	"net/http"
	"time"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	if IsConnected(w, r) {
		if isFieldPostValidd(r) {
			cookie, err := GetCookie(w, r)
			if err != nil {
				tools.LogErr(err)
				return
			}
			userID, err := repo.GetUserIDBySession(bd.GetDB(), cookie)
			if err != nil {
				tools.LogErr(err)
				return
			}
			user, err := repo.GetUserByID(bd.GetDB(), userID)
			if err != nil {
				tools.LogErr(err)
				return
			}
			postID := tools.NeewId()
			post := model.Post{
				PostID:       postID,
				UserID:       userID,
				Title:        r.PostFormValue("title"),
				Content:      r.PostFormValue("desc"),
				Photo:        user.Photo,
				Name:         user.Username,
				LikeCount:    0,
				DislikeCount: 0,
				CommentCount: 0,
				CreatedAt:    time.Now(),
			}

			like := model.Like{
				LikeID: tools.NeewId(),
				UserID: userID,
				PostID: postID,
				Type:   0,
			}

			err = repo.CreateLike(bd.GetDB(), like)

			if err != nil {
				tools.LogErr(err)
				return
			}

			if file, _, _ := r.FormFile("postimage"); file != nil {
				image, err := tools.UploadFile(w, r, "postimage", postID)
				if err == nil {
					post.Image = image
				}
			}

			err = repo.CreatePost(bd.GetDB(), post)
			if err != nil {
				tools.LogErr(err)
				return
			}
			categorie := model.Category{
				CategoryID: tools.NeewId(),
				PostID:     postID,
				Name:       "all",
			}
			r.ParseForm()
			for _, cat := range r.PostForm["cat"] {
				catID := tools.NeewId()
				categorie := model.Category{
					CategoryID: catID,
					PostID:     postID,
					Name:       cat,
				}
				repo.CreateCategory(bd.GetDB(), categorie)
			}

			repo.CreateCategory(bd.GetDB(), categorie)
			return
		} else {
			fmt.Println("Champs incomplet")
		}
	} else {
		fmt.Println("not connected")
	}

}

func isFieldPostValidd(r *http.Request) bool {
	if r.FormValue("publish") != "Publish" || r.
		FormValue("title") == "" || r.FormValue("desc") == "" {
		return false
	}

	r.ParseForm()
	tab := r.PostForm["cat"]
	if len(tab) < 1 {
		return false
	}
	for _, cat := range tab {
		valid := Contains(cat)
		if valid != 1 {
			return false
		}
	}
	return true
}

func Contains(filtre string) int {
	categories := []string{"Event", "General", "Issue"}
	cpt := 0
	for _, v := range categories {
		if v == filtre {
			cpt++
		}
	}
	return cpt
}
