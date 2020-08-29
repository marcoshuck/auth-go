package password

import "golang.org/x/crypto/bcrypt"

// Compare receives the plain and the hashed password,
// encodes the first one and, tries to compare both of them.
// It returns true if they are equal. Returns false in any other case.
func Compare(plain, hashed string) bool {
	bPlain := []byte(plain)
	bHashed := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bHashed, bPlain)
	if err != nil {
		return false
	}
	return true
}
