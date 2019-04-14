package Statics

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Helpers"
)

// Controller ... Static files controller
type Controller struct {
}

// Init ... initialise controller
func (con *Controller) Init(router *mux.Router) {
	router.Handle("/files/{rest}", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))

	router.HandleFunc("/api/static", con.Upload).Methods("POST")
}

// Upload ... upload files handler
func (con *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	fmt.Println("Upload File")

	file, handle, err := r.FormFile("file")
	if err != nil {
		Helpers.Status400(w, err.Error())
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
	case "image/jpg":
	case "image/gif":
	case "image/png":
		saveFile(w, file, handle)
	default:
		Helpers.Status400(w, "Invalid file type")
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handle *multipart.FileHeader) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		Helpers.Status400(w, err.Error())
		return
	}

	filePath := RandomString(32) + handle.Filename

	err = ioutil.WriteFile("./files/"+filePath, data, 0666)
	if err != nil {
		Helpers.Status500(w, "Failed to write file")
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("https://go.reecerussell.com/files/" + filePath))
}

//RandomString - Generate a random string of A-Z chars with len = l
func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
