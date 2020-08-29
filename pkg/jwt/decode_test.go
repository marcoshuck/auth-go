package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	body := JWT{
		UUID:      "1234",
		FirstName: "Marcos",
		LastName:  "Huck",
		Email:     "marcos@huck.com.ar",
	}

	token, _, err := Encode(body, []byte("changeme"))
	assert.NoError(t, err)

	result, err := Decode(*token, []byte("changeme"))
	assert.NoError(t, err)
	assert.Equal(t, body, *result)
}

func TestDecodeFailsWhenDifferentSecretIsUsed(t *testing.T) {
	body := JWT{
		UUID:      "1234",
		FirstName: "Marcos",
		LastName:  "Huck",
		Email:     "marcos@huck.com.ar",
	}

	token, _, err := Encode(body, []byte("changeme"))
	assert.NoError(t, err)

	_, err = Decode(*token, []byte("changeme2"))
	assert.Error(t, err)
}
