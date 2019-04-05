package Authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Service ... a collection of authentication functions
type Service struct {
	HTTPClient HTTPClient
}

var (
	tenantID    = ""
	appKey      = ""
	getTokenURL = ""
	validateURL = ""
)

// Init ... Sets all values needed for auth
func (service *Service) Init() {
	httpClient := HTTPClient{}
	httpClient.Init()

	service.HTTPClient = httpClient

	constants := Constants{}
	tenantID = constants.TenantID()
	appKey = constants.AppKey()
	getTokenURL = constants.GetAuthTokenURL()
	validateURL = constants.ValidateAuthTokenURL()
}

// GetAuthToken ... makes request for auth token
func (service *Service) GetAuthToken(email string, password string) (token AccessToken, err error) {

	credential := Credential{
		Email:    email,
		Password: password,
		TenantID: tenantID,
		AppKey:   appKey,
	}

	fmt.Println("Making request")

	credBytes, err := json.Marshal(credential)
	if err != nil {
		return AccessToken{}, err
	}

	fmt.Println("Parsed json")

	fmt.Printf("Making request to: %s\n", getTokenURL)
	resp, err := service.HTTPClient.Client.Post(getTokenURL,
		"application/json",
		bytes.NewBuffer(credBytes))
	if err != nil {
		return AccessToken{}, errors.New("1: " + err.Error())
	}

	if resp.StatusCode != 200 {
		return AccessToken{}, errors.New("Unable to request access token")
	}

	json.NewDecoder(resp.Body).Decode(&token)

	return token, nil
}

// ValidateToken ... validates token
func (service *Service) ValidateToken(token string) (success bool, err error) {
	req, err := http.NewRequest("GET", validateURL, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := service.HTTPClient.Client.Do(req)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}
