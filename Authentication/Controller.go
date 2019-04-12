package Authentication

import (
	"encoding/json"
	"fmt"
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
	service := Service{}
	service.Init()

	router.HandleFunc("/api/auth", con.Login).Methods("POST")
}

// Login ... Login method for authentication
func (con *Controller) Login(w http.ResponseWriter, r *http.Request) {
	Helpers.Headers(w)

	fmt.Println("Login method, hit")

	var loginCredential LoginCredential
	_ = json.NewDecoder(r.Body).Decode(&loginCredential)

	fmt.Println(loginCredential.Email)
	consts := Constants{}
	fmt.Println(consts.GetAuthTokenURL())

	token, err := con.Service.GetAuthToken(loginCredential.Email, loginCredential.Password)
	if err != nil {
		fmt.Println("    Error: " + err.Error() + "     ")
		Helpers.Status500(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&token)
}
