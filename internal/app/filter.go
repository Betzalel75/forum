package app

import (
	"net/http"

	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"forum/internal/tools"
)

func FitreCategorie(r *http.Request) []model.Post {
	filtre := r.URL.Query().Get("name")
	if filtre != "Event" && filtre != "General" && filtre != "Issue" {
		filtre = "all"
	}
	postIDs, err := repo.GetCategoriesByName(bd.GetDB(), filtre)
	if err != nil {
		tools.LogErr(err)
		return nil
	}
	Posts := []model.Post{}
	for _, post := range postIDs {
		post, err := repo.GetPostByID(bd.GetDB(), post.PostID)
		if err != nil {
			tools.LogErr(err)
			return nil
		}
		Posts = append(Posts, post)
	}
	return Posts
}
