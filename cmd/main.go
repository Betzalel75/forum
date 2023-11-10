package main

import (
	bd "forum/internal/database"
	"forum/internal/serveur"
)

func main() {
	bd.InitDB()
	// repo.DeleteLikeAll(bd.GetDB())
	defer bd.GetDB().Close()
	serveur.Start()
}
