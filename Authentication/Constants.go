package Authentication

var (
	getAuthTokenURLValue      = ""
	validateAuthTokenURLValue = ""
	tenantIDValue             = ""
	appKeyValue               = ""
)

// Constants ...
type Constants struct {
}

// GetAuthTokenURL ... provides auth service the url for the requests
func (constants *Constants) GetAuthTokenURL() string {
	return getAuthTokenURLValue
}

// ValidateAuthTokenURL ... provides auth service the url for the requests
func (constants *Constants) ValidateAuthTokenURL() string {
	return validateAuthTokenURLValue
}

// TenantID ... authentication id
func (constants *Constants) TenantID() string {
	return tenantIDValue
}

// AppKey ... authentication app key
func (constants *Constants) AppKey() string {
	return appKeyValue
}
