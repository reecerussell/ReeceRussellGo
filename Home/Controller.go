package Home

import (
	"encoding/json"
	"net/http"

	"github.com/reecerussell/ReeceRussellGo/Database"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Experience"
	"github.com/reecerussell/ReeceRussellGo/Helpers"
	"github.com/reecerussell/ReeceRussellGo/Project"
)

//
type HomeController struct {
	ProjectDataStore    Project.ProjectDataStore
	ExperienceDataStore Experience.ExperienceDataStore
}

func (con *HomeController) Init(
	database Database.Database,

	router *mux.Router,
) {
	projectDataStore := Project.ProjectDataStore{}
	projectDataStore.Init(database)
	con.ProjectDataStore = projectDataStore

	experienceDataStore := Experience.ExperienceDataStore{}
	experienceDataStore.Init(database)
	con.ExperienceDataStore = experienceDataStore

	router.HandleFunc("/", con.GetViewData).Methods("GET")
}

// GetViewData ... gets data required for home page
func (con *HomeController) GetViewData(w http.ResponseWriter, r *http.Request) {

	reqHeader := r.Header.Get("Requested-By")
	if reqHeader != "reecerussell.com" {
		w.WriteHeader(404)
		return
	}

	Helpers.Headers(w)

	projects, err := con.ProjectDataStore.Get()
	if err != nil {
		projects = nil
	}

	experience, err := con.ExperienceDataStore.Get()
	if err != nil {
		experience = nil
	}

	data := ViewData{
		Projects:   projects,
		Experience: experience,
	}

	json.NewEncoder(w).Encode(&data)
}
