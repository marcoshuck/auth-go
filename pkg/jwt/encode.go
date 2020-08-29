package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// issuer is used to set an issuer to the json web tokens.
const issuer = "marcos.huck.com.ar"

// Encode receives a JWT body and retrieves the access token and the time for it to expire.
// If there is an error, an error is returned instead.
func Encode(body JWT, secretKey []byte) (*string, *int64, error) {
	expiresAt := time.Now().Add(time.Minute * 10).Unix()
	claims := Claims{
		Profile{
			Email:     body.Email,
			FirstName: body.FirstName,
			LastName:  body.LastName,
		},
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    issuer,
			Subject:   body.UUID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(secretKey)
	if err != nil {
		return nil, nil, err
	}
	return &result, &expiresAt, nil
}
