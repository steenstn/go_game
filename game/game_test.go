package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPositionToArrayIndex(t *testing.T) {
	tileWidth := 10
	levelWidth := 10
	assert.Equal(t, 0, getArrayIndex(levelWidth, tileWidth, 2, 3))
	assert.Equal(t, 1, getArrayIndex(levelWidth, tileWidth, 12, 3))
	assert.Equal(t, 2, getArrayIndex(levelWidth, tileWidth, 22, 3))
	assert.Equal(t, 9, getArrayIndex(levelWidth, tileWidth, 92, 3))

	assert.Equal(t, 10, getArrayIndex(levelWidth, tileWidth, 2, 11))
	assert.Equal(t, 11, getArrayIndex(levelWidth, tileWidth, 12, 11))
}
