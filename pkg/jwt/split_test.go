package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	header, payload, signature := Split("a.b.c")
	assert.Equal(t, "a", *header)
	assert.Equal(t, "b", *payload)
	assert.Equal(t, "c", *signature)

	header, payload, signature = Split("a.b.c.d")
	assert.Nil(t, header)
	assert.Nil(t, payload)
	assert.Nil(t, signature)

	header, payload, signature = Split("a.b")
	assert.Nil(t, header)
	assert.Nil(t, payload)
	assert.Nil(t, signature)
}
