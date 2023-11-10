package app

import (
	bd "forum/internal/database"
	repo "forum/internal/database/repository"
	model "forum/internal/models"
	"forum/internal/tools"
)

func AddComment(postID, comment, submit, cookie string) {
	if postID != "" && comment != "" && submit == "Add comment" {
		userID, err := repo.GetUserIDBySession(bd.GetDB(), cookie)
		if err != nil {
			tools.LogErr(err)
			return
		}
		comments := model.Comment{
			CommentID: tools.NeewId(),
			UserID:    userID,
			PostID:    postID,
			Content:   comment,
		}
		er := repo.CreateComment(bd.GetDB(), comments)
		if er != nil {
			tools.LogErr(err)
			return
		}
	}
}
