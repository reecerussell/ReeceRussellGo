package Authentication

// Credential ... provides body for auth requests
type Credential struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	TenantID string `json:"TenantId"`
	AppKey   string `json:"AppKey"`
}
