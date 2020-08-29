package jwt

// JWT represents the JWT body to be encoded.
type JWT struct {
	UUID      string
	FirstName string
	LastName  string
	Email     string
}
