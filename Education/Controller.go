package Education

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Authentication"
	"github.com/reecerussell/ReeceRussellGo/Database"
	"github.com/reecerussell/ReeceRussellGo/Helpers"
)

// Controller ... A collect of functions for education API
type Controller struct {
	DataStore DataStore
}

// Init ... Initializes controller
func (con *Controller) Init(db Database.Database, router *mux.Router) {
	dataStore := DataStore{}
	dataStore.Init(db)

	router.HandleFunc("/api/education", Authentication.Middleware(con.GetAll)).Methods("GET")
	router.HandleFunc("/api/education/{id}", Authentication.Middleware(con.GetByID)).Methods("GET")
	router.HandleFunc("/api/education", Authentication.Middleware(con.Add)).Methods("POST")
	router.HandleFunc("/api/education/{id}", Authentication.Middleware(con.Update)).Methods("PUT")
	router.HandleFunc("/api/education/{id}", Authentication.Middleware(con.Delete)).Methods("DELETE")
}

// GetByID ... Gets an individual education
func (con *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	education, err := con.DataStore.GetByID(id)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	if (education == Education{}) {
		Helpers.Status404(w)
		return
	}

	json.NewEncoder(w).Encode(&education)
}

// GetAll ... Get all education
func (con *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	education, err := con.DataStore.Get()
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&education)
}

// Add ... Add education to database
func (con *Controller) Add(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	var education Education
	_ = json.NewDecoder(r.Body).Decode(&education)

	id, err := con.DataStore.Add(education)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	education.ID = int(id)
	json.NewEncoder(w).Encode(&education)
}

// Update ... Update education
func (con *Controller) Update(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	var education Education
	_ = json.NewDecoder(r.Body).Decode(&education)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	if id != education.ID {
		Helpers.Status400(w, "Id Mismatch")
		return
	}

	err = con.DataStore.Update(education)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&education)
}

// Delete ... Delete education
func (con *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	err = con.DataStore.Delete(id)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}
}
