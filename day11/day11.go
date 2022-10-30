package main

import (
	"adventgo/internal/pkg/intcode"
	"embed"
	"fmt"

	"golang.org/x/exp/slices"
)

//go:embed *.txt
var fs embed.FS

const (
	Black int = iota
	White
)

type Point struct {
	x, y int
}

func GetNewDirection(currentDirection string, deg int) string {
	directions := []string{"U", "R", "D", "L"}
	i := slices.Index(directions, currentDirection)

	if deg == 0 {
		i -= 1
	} else {
		i += 1
	}

	if i < 0 {
		i = len(directions) + i
	} else if i >= len(directions) {
		i = i - len(directions)
	}

	return directions[i]
}

func PartOne(puzzleInput []byte) int {
	painted := make(map[Point]int)
	pointColors := map[Point]int{{0, 0}: Black}
	var x, y int
	var input, color int
	counter := 1
	direction := "U"

	c := intcode.NewComputer(puzzleInput, intcode.WithLoopMode())

	for !c.IsComplete() {
		c.ExecuteProgram(input)
		out := c.GetDiagnosticCode().(int)

		if counter%2 == 0 {
			direction = GetNewDirection(direction, out)
			switch direction {
			case "U":
				y--
			case "D":
				y++
			case "L":
				x--
			case "R":
				x++
			}
			input = pointColors[Point{x, y}]
		} else {
			color = out
			pointColors[Point{x, y}] = color
			painted[Point{x, y}]++
		}

		counter++
	}

	return len(painted)
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput))
}
