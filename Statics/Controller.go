package Statics

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/reecerussell/ReeceRussellGo/Helpers"

	"github.com/gorilla/mux"
)

// Controller ... Static files controller
type Controller struct {
}

// Init ... initialise controller
func (con *Controller) Init(router *mux.Router) {

	router.HandleFunc("/api/static", con.Upload).Methods("POST")

	fs := http.FileServer(http.Dir("./files"))
	router.Handle("/files/", http.StripPrefix("/files", fs))
}

// Upload ... upload files handler
func (con *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, 2048)
	if err := r.ParseMultipartForm(2048); err != nil {
		Helpers.Status400(w, "FILE_TOO_BIG")
		return
	}

	// parse and validate file and post parameters
	fileType := r.PostFormValue("type")
	file, _, err := r.FormFile("uploadFile")
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
	fileName := RandomString(12)
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		Helpers.Status500(w, "CANT_READ_FILE_TYPE")
		return
	}
	newPath := filepath.Join("./files", fileName+fileEndings[0])
	fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

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
	w.Write([]byte("SUCCESS"))
}

//RandomString - Generate a random string of A-Z chars with len = l
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
