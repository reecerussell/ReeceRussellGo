package Project

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

// Controller ... A collection of functions for project api
type Controller struct {
	DataStore ProjectDataStore
}

// Init ... Initializes controller
func (con *Controller) Init(db Database.Database, router *mux.Router) {
	dataStore := ProjectDataStore{}
	dataStore.Init(db)

	router.HandleFunc("/api/projects", Authentication.Middleware(con.GetAll)).Methods("GET")
	router.HandleFunc("/api/projects/{id}", con.GetByID).Methods("GET")
	router.HandleFunc("/api/projects", Authentication.Middleware(con.Add)).Methods("POST")
	router.HandleFunc("/api/projects/{id}", Authentication.Middleware(con.Update)).Methods("PUT")
	router.HandleFunc("/api/projects/{id}", Authentication.Middleware(con.Delete)).Methods("DELETE")
}

// GetByID ... Gets an individual project
func (con *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	project, err := con.DataStore.GetByID(id)
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status500(w, err.Error())
		return
	}

	if (project == Project{}) {
		Helpers.Status404(w)
		return
	}

	json.NewEncoder(w).Encode(&project)
}

// GetAll ... Get all projects
func (con *Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	projects, err := con.DataStore.Get()
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&projects)
}

// Add ... Add project to database
func (con *Controller) Add(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	var project Project
	_ = json.NewDecoder(r.Body).Decode(&project)

	id, err := con.DataStore.Add(project)
	if err != nil {
		Helpers.Status500(w, err.Error())
		return
	}

	project.ID = int(id)
	json.NewEncoder(w).Encode(project)
}

// Update ... Update project
func (con *Controller) Update(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	var project Project
	_ = json.NewDecoder(r.Body).Decode(&project)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	if id != project.ID {
		fmt.Println(err.Error())
		Helpers.Status400(w, "Id Mismatch")
		return
	}

	err = con.DataStore.Update(project)
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(project)
}

// Delete ... Delete project
func (con *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status400(w, "Invalid Id format")
		return
	}

	err = con.DataStore.Delete(id)
	if err != nil {
		fmt.Println(err.Error())
		Helpers.Status500(w, "Could not connect to database")
		return
	}
}
