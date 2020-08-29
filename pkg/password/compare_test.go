package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	hash, err := Generate("test")
	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.True(t, len(*hash) > 0)

	assert.True(t, Compare("test", *hash))
	assert.False(t, Compare("tes", *hash))
	assert.False(t, Compare("piaaopdjaiwjaodowaiojtao12412515", *hash))
	assert.False(t, Compare("testtest", *hash))
	assert.False(t, Compare("", *hash))
}
