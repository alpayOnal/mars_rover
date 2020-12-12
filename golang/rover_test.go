package main

import (
	"fmt"
	"testing"
)

var plateau Plateau

func TestMain(m *testing.M) {
	plateau = NewPlateau(4, 4)
	m.Run()
}

func TestRoverCase5(t *testing.T) {
	rover1 := NewRover(1, 1, "E", plateau)
	rover2 := NewRover(3, 1, "E", plateau)

	rover1.Run("MM")
	//rover2.Run("MM")
	fmt.Println(rover2)

	//assert.Equal(t, "E", rover.Direction)
	//assert.Equal(t, 4, rover.XCoordinates)
	//assert.Equal(t, 3, rover.YCoordinates)
}

//
//func TestRoverTurnLeft(t *testing.T) {
//	rover := NewRover(1, 2, "E", plateau)
//	rover.turnLeft()
//	assert.Equal(t, "N", rover.Direction)
//}
//
//func TestRoverTurnRight(t *testing.T) {
//	rover := NewRover(1, 2, "E", plateau)
//	rover.turnRight()
//	assert.Equal(t, "S", rover.Direction)
//}
//
//func TestRoverMoveEast(t *testing.T) {
//	rover := NewRover(1, 2, "E", plateau)
//	rover.move()
//	assert.Equal(t, 2, rover.XCoordinates)
//}
//
//func TestRoverMovWest(t *testing.T) {
//	rover := NewRover(1, 2, "W", plateau)
//	rover.move()
//	assert.Equal(t, 0, rover.XCoordinates)
//}
//
//func TestRoverMoveNorth(t *testing.T) {
//	rover := NewRover(1, 2, "N", plateau)
//	rover.move()
//	assert.Equal(t, 3, rover.YCoordinates)
//}
//
//func TestRoverMoveSouth(t *testing.T) {
//	rover := NewRover(1, 2, "S", plateau)
//	rover.move()
//	assert.Equal(t, 1, rover.YCoordinates)
//}
//
//func TestRoverCase1(t *testing.T) {
//	rover := NewRover(1, 2, "N", plateau)
//	rover.Run("LMLMLMLMM")
//	assert.Equal(t, "N", rover.Direction)
//	assert.Equal(t, 1, rover.XCoordinates)
//	assert.Equal(t, 3, rover.YCoordinates)
//}
//
//func TestRoverCase2(t *testing.T) {
//	rover := NewRover(3, 3, "E", plateau)
//	rover.Run("MMRMMRMRRM")
//	assert.Equal(t, "E", rover.Direction)
//	assert.Equal(t, 5, rover.XCoordinates)
//	assert.Equal(t, 1, rover.YCoordinates)
//}
