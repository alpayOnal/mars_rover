package golang

type Plateau struct {
	MaxX int
	MaxY int
	MinX int
	MinY int
}

func NewPlateau(maxX int, maxY int) Plateau {
	p := Plateau{maxX, maxY, 0, 0}

	return p
}
