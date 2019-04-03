package Authentication

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reecerussell/ReeceRussellGo/Helpers"
)

// Controller ... provides router a collect of auth methods
type Controller struct {
	Service Service
}

// Init ... sets up router routes
func (con *Controller) Init(router *mux.Router) {
	router.HandleFunc("/api/auth", con.Login)
}

// Login ... Login method for authentication
func (con *Controller) Login(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	var loginCredential LoginCredential
	_ = json.NewDecoder(r.Body).Decode(&loginCredential)

	token, err := con.Service.GetAuthToken(loginCredential.Email, loginCredential.Password)
	if err != nil {
		Helpers.Status500(err.Error())
		return
	}

	json.NewEncoder(w).Encode(&token)
}
