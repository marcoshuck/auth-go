package password

import "golang.org/x/crypto/bcrypt"

// Generate generates a bcrypt password from the given plain password.
func Generate(plain string) (*string, error) {
	bslice := []byte(plain)

	generated, err := bcrypt.GenerateFromPassword(bslice, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashed := string(generated)

	return &hashed, nil
}
