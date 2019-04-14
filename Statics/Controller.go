package Statics

import (
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/reecerussell/ReeceRussellGo/Helpers"

	"github.com/gorilla/mux"
)

// Controller ... Static files controller
type Controller struct {
}

// Init ... initialise controller
func (con *Controller) Init(router *mux.Router) {

	router.HandleFunc("/api/static", con.Upload).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./files"))))
}

// Upload ... upload files handler
func (con *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, 2560000)
	if err := r.ParseMultipartForm(2560000); err != nil {
		Helpers.Status400(w, "FILE_TOO_BIG")
		return
	}

	// parse and validate file and post parameters
	file, _, err := r.FormFile("file")
	if err != nil {
		Helpers.Status400(w, "INVALID_FILE")
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		Helpers.Status400(w, "INVALID_FILE")
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
		break
	default:
		Helpers.Status400(w, "INVALID_FILE_TYPE")
		return
	}
	fileName := uuid.New().String()
	fileEndings, err := mime.ExtensionsByType(filetype)
	newPath := filepath.Join("./files", fileName+fileEndings[0])
	if err != nil {
		Helpers.Status500(w, "CANT_READ_FILE_TYPE")
		return
	}

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		Helpers.Status500(w, "CANT_WRITE_FILE")
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		Helpers.Status500(w, "CANT_WRITE_FILE")
		return
	}
	w.Write([]byte("https://go.reecerussell.com/" + newPath))
}
