package main

import (
	"fmt"
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
	if r.Direction == "N" && r.Plateau.MaxX > r.YCoordinates {
		r.YCoordinates += 1
	}

	if r.Direction == "E" && r.Plateau.MaxX > r.YCoordinates {
		r.XCoordinates += 1
	}

	if r.Direction == "S" && r.Plateau.MaxY > r.YCoordinates && r.YCoordinates > r.Plateau.MinY {
		r.YCoordinates -= 1
	}

	if r.Direction == "W" && r.Plateau.MaxX > r.YCoordinates && r.YCoordinates > r.Plateau.MinY{
		r.XCoordinates -= 1
	}
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

	return r
}
