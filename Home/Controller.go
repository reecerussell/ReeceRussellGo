package Home

import (
	"encoding/json"
	"net/http"

	"github.com/reecerussell/ReeceRussellGo/Database"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Helpers"
)

// Controller ... handles the retreival of data required for the home page
type Controller struct {
	DataStore DataStore
}

// Init ... initialises the controller
func (con *Controller) Init(
	database Database.Database,
	router *mux.Router,
) {
	dataStore := DataStore{}
	dataStore.Init(database)

	router.HandleFunc("/", con.GetViewData).Methods("GET")
}

// GetViewData ... gets data required for home page
func (con *Controller) GetViewData(w http.ResponseWriter, r *http.Request) {

	reqHeader := r.Header.Get("Requested-By")
	if reqHeader != "reecerussell.com" {
		w.WriteHeader(404)
		return
	}

	Helpers.Headers(w)

	data := con.DataStore.GetViewData()

	json.NewEncoder(w).Encode(&data)
}
