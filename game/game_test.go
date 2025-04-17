package game

import (
	"fmt"
	"testing"
)

func TestConvertPositionToArrayIndex(t *testing.T) {
	tileWidth := 10
	levelWidth := 10
	assertEquals(0, getArrayIndex(levelWidth, tileWidth, 2, 3))
	assertEquals(1, getArrayIndex(levelWidth, tileWidth, 12, 3))
	assertEquals(2, getArrayIndex(levelWidth, tileWidth, 22, 3))
	assertEquals(9, getArrayIndex(levelWidth, tileWidth, 92, 3))

	assertEquals(10, getArrayIndex(levelWidth, tileWidth, 2, 11))
	assertEquals(11, getArrayIndex(levelWidth, tileWidth, 12, 11))
}

// TODO Use real assertion
func assertEquals(expected int, actual int) {
	if expected != actual {
		fmt.Printf("Expected %d, got %d", expected, actual)
		panic("Failed")

	}
}
