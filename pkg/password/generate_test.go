package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	hash, err := Generate("test")
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.True(t, len(*hash) > 0)
}
