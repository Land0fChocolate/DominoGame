package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// XXXX
// X23X
// XXXX

//using the assert library to compare the actual output with expected output
func Test_expandGrid(t *testing.T) {
	testDominoGrid := dominoGrid{grid: [][]string{{"X", "X", "X", "X"}, {"X", "2", "3", "X"}, {"X", "X", "X", "X"}}}
	assert.Equal(t,
		dominoGrid(dominoGrid{grid: [][]string{[]string{"X", "X", "X", "X"}, []string{"X", "X", "X", "X"}, []string{"X", "2", "3", "X"}, []string{"X", "X", "X", "X"}}}),
		expandGrid("top", testDominoGrid))

	testDominoGrid = dominoGrid{grid: [][]string{{"X", "X", "X", "X"}, {"X", "2", "3", "X"}, {"X", "X", "X", "X"}}}
	assert.Equal(t,
		dominoGrid(dominoGrid{grid: [][]string{[]string{"X", "X", "X", "X", "X"}, []string{"X", "2", "3", "X", "X"}, []string{"X", "X", "X", "X", "X"}}}),
		expandGrid("right", testDominoGrid))

	testDominoGrid = dominoGrid{grid: [][]string{{"X", "X", "X", "X"}, {"X", "2", "3", "X"}, {"X", "X", "X", "X"}}}
	assert.Equal(t,
		dominoGrid(dominoGrid{grid: [][]string{[]string{"X", "X", "X", "X"}, []string{"X", "2", "3", "X"}, []string{"X", "X", "X", "X"}, []string{"X", "X", "X", "X"}}}),
		expandGrid("bot", testDominoGrid))

	testDominoGrid = dominoGrid{grid: [][]string{{"X", "X", "X", "X"}, {"X", "2", "3", "X"}, {"X", "X", "X", "X"}}}
	assert.Equal(t,
		dominoGrid(dominoGrid{grid: [][]string{[]string{"X", "X", "X", "X", "X"}, []string{"X", "X", "2", "3", "X"}, []string{"X", "X", "X", "X", "X"}}}),
		expandGrid("left", testDominoGrid))
}

func Test_isSpaceAlreadyOccupied(t *testing.T) {
	testDominoGrid := dominoGrid{grid: [][]string{{"X", "X", "X", "X"}, {"X", "2", "3", "X"}, {"X", "X", "X", "X"}}}

	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 1, 1))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 2, 1))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 3, 1))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 4, 1))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 1, 2))
	assert.Equal(t, true, isSpaceAlreadyOccupied(testDominoGrid, 2, 2))
	assert.Equal(t, true, isSpaceAlreadyOccupied(testDominoGrid, 3, 2))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 4, 2))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 1, 3))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 2, 3))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 3, 3))
	assert.Equal(t, false, isSpaceAlreadyOccupied(testDominoGrid, 4, 3))
}

//isSpaceNextToEquivalentEnd(newGrid dominoGrid, x, y int, end string) (spaceViable bool

func Test_isSpaceNextToEquivalentEnd(t *testing.T) {
	testDominoGrid := dominoGrid{grid: [][]string{{"X", "X", "X", "X"}, {"X", "X", "3", "X"}, {"X", "X", "X", "X"}}}
	
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 1, 1, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 1, 2, "3"))
	assert.Equal(t, true, isSpaceNextToEquivalentEnd(testDominoGrid, 1, 3, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 1, 4, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 2, 1, "3"))
	assert.Equal(t, true, isSpaceNextToEquivalentEnd(testDominoGrid, 2, 2, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 2, 3, "3"))
	assert.Equal(t, true, isSpaceNextToEquivalentEnd(testDominoGrid, 2, 4, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 3, 1, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 3, 2, "3"))
	assert.Equal(t, true, isSpaceNextToEquivalentEnd(testDominoGrid, 3, 3, "3"))
	assert.Equal(t, false, isSpaceNextToEquivalentEnd(testDominoGrid, 3, 4, "3"))
}
