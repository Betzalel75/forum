package tools

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
)

func NeewId() string {
	id, err := uuid.NewV4()
	if err != nil {
		log.Panicln(err)
	}
	return id.String()
}

func UploadFile(w http.ResponseWriter, r *http.Request, name, id string) (string, error) {
	// Vérifiez si la taille du fichier est inférieure à 800KB
	if r.ContentLength > 800*1024 {
		return "Image too large", errors.New("error")
	}
	// Récupérez le fichier à partir de la requête
	file, header, err := r.FormFile(name)
	if err != nil {
		return "", nil
	}
	defer file.Close()
	// Vérifiez si le fichier est une image
	ext := filepath.Ext(header.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".svg" && ext != ".gif" && ext != ".jpeg" {
		return "Non-authorized file format. Ex: .jpg, .svg, .gif, .jpeg", errors.New("error")
	}
	// Créez un nouveau nom pour le fichier
	newFilename := id + ext
	// Stockez le fichier dans un dossier avec le nouveau nom
	dst, err := os.Create("./web/static/media/" + newFilename)
	if err != nil {
		return "unable to download file", errors.New("error")
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return "unable to download file", errors.New("error")
	}
	return newFilename, nil
}

var (
	Error *log.Logger
)

func Init() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Error = log.New(file,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func LogErr(err error) {
	if err != nil {
		Error.Println(err.Error())
	}
}
