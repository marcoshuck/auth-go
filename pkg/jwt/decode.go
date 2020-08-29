package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// Decode tries to convert a string to a JWT.
// It returns an error if something goes wrong.
func Decode(token string, secretKey []byte) (*JWT, error) {
	parsed, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsed.Claims.(*Claims); ok && parsed.Valid {
		return &JWT{
			UUID:      claims.Subject,
			FirstName: claims.Profile.FirstName,
			LastName:  claims.Profile.LastName,
			Email:     claims.Profile.Email,
		}, nil
	}
	return nil, errors.New("unable to decode JWT")
}
