package Authentication

// AccessToken ... OAuth data
type AccessToken struct {
	Token   string `json:"access_token"`
	Expires string `json:"expires"`
}
