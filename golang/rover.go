package main

import (
	"fmt"
	"log"
	"strconv"
)

type Rover struct {
	XCoordinates int
	YCoordinates int
	Direction    string
	Plateau      Plateau
}

func (r *Rover) turnRight() {
	directions := map[string]string{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}
	r.Direction = directions[r.Direction]
}

func (r *Rover) turnLeft() {
	directions := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}
	r.Direction = directions[r.Direction]
}

func (r *Rover) move() {
	nextX:=r.XCoordinates
	nextY:=r.YCoordinates
	if r.Direction == "N" && r.Plateau.MaxY > r.YCoordinates {
		nextX = r.YCoordinates+1
	}

//	N
//W	   E
//	S
	if r.Direction == "E" && r.Plateau.MaxX > r.XCoordinates {
		nextX = r.XCoordinates+1


	}

	if r.Direction == "S" && r.Plateau.MaxY > r.YCoordinates && r.YCoordinates > r.Plateau.MinY {
		nextY = r.YCoordinates -1
	}

	if r.Direction == "W" && r.Plateau.MaxX > r.YCoordinates && r.YCoordinates > r.Plateau.MinY{
		nextX = r.XCoordinates - 1
	}

	for _, r :=range  r.Plateau.Rover {
		if r.XCoordinates == nextX && r.YCoordinates == nextY {
			fmt.Println("baska bir rover")
			log.Fatal("log.")
		}
	}

	r.YCoordinates = nextY
	r.XCoordinates = nextX


}

func (r *Rover) Run(commands string) {
	for _, command := range commands {
		c := string(command)
		switch c {
		case "R":
			r.turnRight()
		case "L":
			r.turnLeft()
		case "M":
			r.move()
		}
	}
}

func (r *Rover) position() string {
	return fmt.Sprintf("%s %s %s ", strconv.Itoa(r.XCoordinates), strconv.Itoa(r.YCoordinates), r.Direction)
}
func NewRover(x int, y int, direction string, plateau Plateau) Rover {
	r := Rover{x, y, direction, plateau}
	plateau.Rover = append(plateau.Rover, r)
	return r
}
