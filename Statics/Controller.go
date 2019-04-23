package Statics

import (
	"net/http"

	"github.com/reecerussell/ReeceRussellGo/Authentication"

	"github.com/reecerussell/ReeceRussellGo/Helpers"

	"github.com/gorilla/mux"
)

// Controller ... Static files controller
type Controller struct {
	Service Service
}

// Init ... initialise controller
func (con *Controller) Init(router *mux.Router) {
	con.Service = Service{}

	router.HandleFunc("/api/static", Authentication.Middleware(con.Upload)).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./files"))))
}

// Upload ... upload files handler
func (con *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	fileName, status, err := con.Service.CreateFile(w, r)
	if err != nil {
		if status == 400 {
			Helpers.Status400(w, err.Error())
			return
		}
		if status == 500 {
			Helpers.Status500(w, err.Error())
			return
		}
	}

	w.Write([]byte("https://go.reecerussell.com/static" + fileName))
}
