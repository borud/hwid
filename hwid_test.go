package hwid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	// straightforward without params
	id, err := ID()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	// check for empty string param, which should behave like without params
	id2, err := ID("")
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.Equal(t, id, id2)

	id, err = ID("thisCannotPossiblyBeAValidInterface")
	assert.Error(t, err)
}
