package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	body := JWT{
		UUID:      "1234",
		FirstName: "Marcos",
		LastName:  "Huck",
		Email:     "marcos@huck.com.ar",
	}

	now := time.Now()

	token, expiresAt, err := Encode(body, []byte("changeme"))

	assert.NoError(t, err)

	assert.NotNil(t, token)
	assert.NotNil(t, expiresAt)

	assert.True(t, len(*token) > 0)
	assert.True(t, *expiresAt > now.Unix())
}
