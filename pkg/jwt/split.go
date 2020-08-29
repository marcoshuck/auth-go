package jwt

import "strings"

// Split receives an string and returns the header, the payload and the signature of a JWT.
func Split(token string) (header, payload, signature *string) {
	slice := strings.Split(token, ".")
	if len(slice) != 3 {
		return nil, nil, nil
	}
	return &slice[0], &slice[1], &slice[2]
}
