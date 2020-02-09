package golang

import (
	"fmt"
)

func main() {
	plateau := NewPlateau(5, 5)
	rover := NewRover(1, 2, "S", plateau)
	rover.Run("LMLMLMLMM")
	fmt.Println(rover.position())

	rover = NewRover(3, 3, "E", plateau)
	rover.Run("MMRMMRMRRM")
	fmt.Println(rover.position())
}
