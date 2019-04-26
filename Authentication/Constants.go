package Authentication

var (
	getAuthTokenURLValue      = "https://singleauth.co.uk/api/auth/token"
	validateAuthTokenURLValue = "https://singleauth.co.uk/api/auth/check"
	tenantIDValue             = "8b39d4e2-7ac7-4c67-811f-0b4675116a75"
	appKeyValue               = "-----BEGIN CERTIFICATE----- MIID0zCCArqgAwIBAgIBADANBgkqhkiG9w0BAQ0FADCBgjELMAkGA1UEBhMCdWsx GDAWBgNVBAgMD0J1Y2tpbmdoYW1zaGlyZTESMBAGA1UECgwJV2ViU3lzdGVtMRkw FwYDVQQDDBByZWVjZXJ1c3NlbGwuY29tMRYwFAYDVQQHDA1NaWx0b24gS2V5bmVz MRIwEAYDVQQLDAlXZWJTeXN0ZW0wHhcNMTkwNDI2MTYxOTU3WhcNMjAwNDI1MTYx OTU3WjCBgjELMAkGA1UEBhMCdWsxGDAWBgNVBAgMD0J1Y2tpbmdoYW1zaGlyZTES MBAGA1UECgwJV2ViU3lzdGVtMRkwFwYDVQQDDBByZWVjZXJ1c3NlbGwuY29tMRYw FAYDVQQHDA1NaWx0b24gS2V5bmVzMRIwEAYDVQQLDAlXZWJTeXN0ZW0wggEjMA0G CSqGSIb3DQEBAQUAA4IBEAAwggELAoIBAgCyZtVlY6yae4vRJoSB+PakyKtmuAtT QCjqLENDrkO+xXDZKhbtrpkpjsesjUdj6ksDZCo3kiqAjbJXzFf8LKVu4YT/FV3h XwoIMCm8n5TBkZufNXQGxq4TZ6zchPk8kSJxwtq7QwqLwBaXWKcFiVYQZdPXJthh BH1qEvdERkkZ/9HL1mAMhLYRQeBHHTom0mWzHr/DbaA5cRUGMFHd3FoBeFIKBQcq /Vci2B8+BV5mXskyfPBK2MS92GyM9fS73C0WsBb2beNwJpfBTLKGAyHyrKsps50K rlhK+ygr8JwUucT1mF0u0aGT0ydfZvicml0JBNvTjutbtivRGAlgJ2+byQIDAQAB o1AwTjAdBgNVHQ4EFgQUBVEaOQ/iBao21KzBFRXY7d0ryR0wHwYDVR0jBBgwFoAU BVEaOQ/iBao21KzBFRXY7d0ryR0wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQ0F AAOCAQIAk4JYIXMdhaR7EcAyzel98AugjPg/GLXdAC+wp5OJjZd8zsEX/dtv6uxi JLkh39R2Jk4SH/axHzhtr9cVcm6WXDBCi3p1+IsIDNhIaElWcGuSS0P7bWLBtSaQ U8VGOITJZEsMQGxeIpQllgoyRGcr/5PC7gmuPOqKUHE8zMRKTPcOLM2pssAdzJip AKiPjgaPDOWZ37nOFLINh7EBT6yau6pJ6vqJkKfqKq8F4b6lWESA8jAtbZeY7aL5 Xg8zUHCyWrymMahbY+ckS5DcNVPcpEYte6+NtrpVT6LHkeB2iE5TMQsz/uCgaMui F8qgqAjIxiotYKdAMdsjn6SJ1Ro3DFQ= -----END CERTIFICATE-----"
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
