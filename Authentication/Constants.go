package Authentication

var (
	getAuthTokenURLValue      = "https://singleauth.co.uk/api/auth/token"
	validateAuthTokenURLValue = "https://singleauth.co.uk/api/auth/check"
	tenantIDValue             = "298cd37c-f712-4d77-bd59-dc32d6e96db7"
	appKeyValue               = "-----BEGIN CERTIFICATE----- MIIDhzCCAm6gAwIBAgIBADANBgkqhkiG9w0BAQ0FADBdMQswCQYDVQQGEwJ1azEY MBYGA1UECAwPQnVja2luZ2hhbXNoaXJlMRYwFAYDVQQKDA1SZWVjZSBSdXNzZWxs MRwwGgYDVQQDDBNnby5yZWVjZXJ1c3NlbGwuY29tMB4XDTE5MDQwODEyNTQzNFoX DTIwMDQwNzEyNTQzNFowXTELMAkGA1UEBhMCdWsxGDAWBgNVBAgMD0J1Y2tpbmdo YW1zaGlyZTEWMBQGA1UECgwNUmVlY2UgUnVzc2VsbDEcMBoGA1UEAwwTZ28ucmVl Y2VydXNzZWxsLmNvbTCCASMwDQYJKoZIhvcNAQEBBQADggEQADCCAQsCggECAMyi wCkX6xkQeRlcxChEO0xigvHQQ0c5yg3XpPbNjDksTi84tpYAycBZOgJUjcuT3rq0 LRFiKdzfY8gTXh6rgp/EQB2jccfVEJ2UvssCaFH+xHl+A9LVBWDdUM+3oe2pmHtV pn4NtGmXMv4SFkEzQy8chjzlSZx9fVV/9ORLJ4M3Wigr+ZfKZxqjuj0Pm66/+fdK ylO6vmmhIoadYS1Vg/MpG5Av2TPOVZhT9YuzWEoDy+lcmdT4uuFGFDobJTM2mycG ObfYvOzk2NFKhMeIw3l2r2SeMAMh3lErr0tVNnEi5mom2MWD20b0o7o/oOdLETun NIkFsC/Kux1h+X/+uzQNAgMBAAGjUDBOMB0GA1UdDgQWBBTQPpPbgq0C7YqL7VoI JO9m6iDZtTAfBgNVHSMEGDAWgBTQPpPbgq0C7YqL7VoIJO9m6iDZtTAMBgNVHRME BTADAQH/MA0GCSqGSIb3DQEBDQUAA4IBAgBHyE4MY9KT85Rb5ckyD7jYCFQfOMz5 EIuLTKcufILDdwzz+nmyyBuAMQ7hGj6RLXBFgZw3SQHL4VU+Y7n4a0hGpanQev2G Us4Cu7jC2/F3NNDDtCAoJrq3+o/rvu7PwHAZnWZXC2ZxV1lQWl3X7KsZH6dblsdD K8otPr9LxCYflJ9T6mse8bq51XMPxUpNfZY8vu+FC8latV/HXB56ahHoX9UAGTST pcCqvdO3lWB0W7NnIIvMo0hmXeX9El3TyXW7X5wPaNm9F0CVZ70puPqeaSiVe3tJ kyvycJ7v3R2IWwCWdN81rcOXsBO/M55cqGrWM1bK9BP5uRWh74CcGY1/Cg== -----END CERTIFICATE-----"
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
