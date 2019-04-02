package Experience

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

// Controller ... A collect of functions for experience API
type Controller struct {
	DataStore ExperienceDataStore
}

// Init ... Initializes controller
func (con *Controller) Init(db Database.Database, router *mux.Router) {
	dataStore := ExperienceDataStore{}
	dataStore.Init(db)

	router.HandleFunc("/api/experience", Authentication.Middleware(con.GetAll)).Methods("GET")
	router.HandleFunc("/api/experience/{id}", Authentication.Middleware(con.GetByID)).Methods("GET")
	router.HandleFunc("/api/experience", Authentication.Middleware(con.Add)).Methods("POST")
	router.HandleFunc("/api/experience/{id}", Authentication.Middleware(con.Update)).Methods("PUT")
	router.HandleFunc("/api/experience/{id}", Authentication.Middleware(con.Delete)).Methods("DELETE")
}

// GetByID ... Gets an individual experience
func (con *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	experience, err := con.DataStore.GetByID(id)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	if (experience == Experience{}) {
		Helpers.Status404(w)
		return
	}

	json.NewEncoder(w).Encode(&experience)
}

// GetAll ... Get all experience
func (con *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	experience, err := con.DataStore.Get()
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&experience)
}

// Add ... Add experience to database
func (con *Controller) Add(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	var experience Experience
	_ = json.NewDecoder(r.Body).Decode(&experience)

	id, err := con.DataStore.Add(experience)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	experience.ID = int(id)
	json.NewEncoder(w).Encode(experience)
}

// Update ... Update experience
func (con *Controller) Update(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	var experience Experience
	_ = json.NewDecoder(r.Body).Decode(&experience)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	if id != experience.ID {
		Helpers.Status400(w, "Id Mismatch")
		return
	}

	err = con.DataStore.Update(experience)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(experience)
}

// Delete ... Delete experience
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
