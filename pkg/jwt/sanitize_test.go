package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSanitize(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
		"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	assert.True(t, Sanitize(token))

	token = "a.b.c"

	assert.False(t, Sanitize(token))
}
