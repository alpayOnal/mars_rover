package main

type Plateau struct {
	MaxX int
	MaxY int
	MinX int
	MinY int
	Rover []Rover
}

func NewPlateau(maxX int, maxY int) Plateau {
	p := Plateau{maxX, maxY, 0, 0,[]Rover{}}

	return p
}
