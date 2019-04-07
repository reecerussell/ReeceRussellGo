package Skills

import (
	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Database"
)

// Controller ... a collection of methods used by the router
type Controller struct {
	DataStore DataStore
}

// Init
func (con *Controller) Init(db Database.Database, router *mux.Router) {
	dataStore := DataStore{}
	dataStore.Init(db)
}
