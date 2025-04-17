package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadLevelMultipleLines(t *testing.T) {

	result := LoadLevel("testlevel.txt")

	assert.Equal(t, 1, result[0], "pos 0")
	assert.Equal(t, 1, result[1], "pos 1")
	assert.Equal(t, 1, result[2], "pos 2")
	assert.Equal(t, 0, result[3], "pos 3")
	assert.Equal(t, 0, result[4], "pos 4")
	assert.Equal(t, 0, result[5], "pos 5")
	assert.Equal(t, 1, result[6], "pos 6")
	assert.Equal(t, 1, result[7], "pos 7")
	assert.Equal(t, 1, result[8], "pos 8")
}

func TestLoadLevelSingleLine(t *testing.T) {

	result := LoadLevel("testlevel_oneline.txt")

	assert.Equal(t, 1, result[0], "pos 0")
	assert.Equal(t, 1, result[1], "pos 1")
	assert.Equal(t, 1, result[2], "pos 2")
	assert.Equal(t, 0, result[3], "pos 3")
	assert.Equal(t, 0, result[4], "pos 4")
	assert.Equal(t, 0, result[5], "pos 5")
	assert.Equal(t, 1, result[6], "pos 6")
	assert.Equal(t, 1, result[7], "pos 7")
	assert.Equal(t, 1, result[8], "pos 8")
}
