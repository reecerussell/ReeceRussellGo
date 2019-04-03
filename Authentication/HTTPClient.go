package Authentication

import (
	"net/http"
	"time"
)

// HTTPClient ... custom http.Client object with timeout
type HTTPClient struct {
	Client http.Client
}

// Init ... initialises HttpClient
func (client *HTTPClient) Init() {
	client.Client = http.Client{
		Timeout: time.Second * 10,
	}
}
