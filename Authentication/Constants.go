package Authentication

var (
	getAuthTokenURLValue      = "https://localhost:44306/api/auth/token"
	validateAuthTokenURLValue = "https://localhost:44306/api/auth/check"
	tenantIDValue             = "7d2e5c4a-7e8e-48bb-8e56-d82c339adae8"
	appKeyValue               = "-----BEGIN CERTIFICATE----- MIIDYzCCAkqgAwIBAgIBADANBgkqhkiG9w0BAQ0FADBLMQswCQYDVQQGEwJ1azEY MBYGA1UECAwPQnVja2luZ2hhbXNoaXJlMQ0wCwYDVQQKDARUZXN0MRMwEQYDVQQD DAphdXRoLmxvY2FsMB4XDTE5MDQwMTEzMjczOVoXDTIwMDMzMTEzMjczOVowSzEL MAkGA1UEBhMCdWsxGDAWBgNVBAgMD0J1Y2tpbmdoYW1zaGlyZTENMAsGA1UECgwE VGVzdDETMBEGA1UEAwwKYXV0aC5sb2NhbDCCASMwDQYJKoZIhvcNAQEBBQADggEQ ADCCAQsCggECALAK2fZfReaOVk/vEJebeG2n6cvys3yDu891hnBDEKLVdc/5qxBW aj8/OgS8CUcB+JB6EmqaUI7YV4/df+n3/zjPNUQ7c3RKBVZsNlOitAAELXW+T97G S0GJnSLtsSR6v/MsUYMGZ7siNrdHiYrYgO7pxBSSTAgq7IFJ3vJ4ivuw1EgISOov Ea8VT0PlDjbHFgfjaUWmI9wsVI6yzgapkU7PED5LH/4OaiYPYs6pfs7I/ed3QCPv SX9uICHnZe/bTnRLTnUWmP1IJ40dsFQ9NC6CPXGBHBKuufo4OFw00BDRHze1Du2U W1IurFpoipAf2WMsSgeXkhzqDBj20rvvDbq3AgMBAAGjUDBOMB0GA1UdDgQWBBSz YU9kWeRuOifiH9FcI1L9lmOGLjAfBgNVHSMEGDAWgBSzYU9kWeRuOifiH9FcI1L9 lmOGLjAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBDQUAA4IBAgAG7wWQPHuha0v1 Lj6n44KkFtYYAQr5R1BeYVr04Gy6AZ8j1kEsrG9gAS3pnIc7q+DWm/Blu9LuJPk5 B3W9A9Hw25XFQPvgng1GRDZyvV92EQ1UowWjUBtQRmgBpB/WmtUqmbo5Zn47uLvj hycz2JF3dp9difRgJKrHySRkrikF5KkJiKi2pmTKnDAwBYsN0b8wd/WxnulfnHse sI1CxQisKHGHFx6l9D/ftoJputgHhLfVFIUCtEQVCF3N27xkyfI9kl1YT5qKmWKy Od1SIh6ud/lZ0C58BxUSS3dbigJIRK6O1vJLaoIAaea1q6HZICJZ6rMrNFZbSN5Z emHSeeUwAw== -----END CERTIFICATE-----"
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
