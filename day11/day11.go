package main

import (
	"adventgo/internal/pkg/intcode"
	"embed"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed *.txt
var fs embed.FS

type Color int
type TurnDirection int

const (
	Black Color = iota
	White
)

const (
	Left TurnDirection = iota
	Right
)

type Point struct {
	x, y int
}

type PaintRobot struct {
	Input       Color
	PointColors map[Point]Color

	computer         *intcode.Computer
	currentColor     Color
	currentDirection string
	paintedPoints    map[Point]int
	x                int
	xValues          []int
	y                int
	yValues          []int
}

func NewRobot(program []byte, startingColor Color) *PaintRobot {
	r := &PaintRobot{
		Input:            startingColor,
		PointColors:      map[Point]Color{{50, 50}: startingColor},
		computer:         intcode.NewComputer(program, intcode.WithLoopMode()),
		currentDirection: "U",
		paintedPoints:    make(map[Point]int),
		x:                50,
		y:                50,
	}

	return r
}

func (r *PaintRobot) PaintHull() {
	counter := 1
	for !r.computer.IsComplete() {
		r.computer.ExecuteProgram(int(r.Input))

		out := r.computer.GetDiagnosticCode().(int)
		if counter%2 == 0 {
			r.turn(TurnDirection(out))
		} else {
			r.currentColor = Color(out)
			p := Point{r.x, r.y}
			r.PointColors[p] = r.currentColor
			r.paintedPoints[p]++
		}
		counter++
	}
}

func (r *PaintRobot) RenderHull() {
	numRows := 100
	numCols := 100

	hull := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]string, numCols)
		for j := 0; j < numCols; j++ {
			row[j] = " "
		}

		hull[i] = row
	}

	for p, c := range r.PointColors {
		t := " "
		if c == White {
			t = "|"
		}

		hull[p.y][p.x] = t
	}

	for _, row := range hull {
		fmt.Println(strings.Join(row, ""))
	}
}

func (r *PaintRobot) turn(d TurnDirection) {
	directions := []string{"U", "R", "D", "L"}
	i := slices.Index(directions, r.currentDirection)

	if d == 0 {
		i--
	} else {
		i++
	}

	if i < 0 {
		i = len(directions) + i
	} else if i >= len(directions) {
		i = i - len(directions)
	}

	r.currentDirection = directions[i]

	switch r.currentDirection {
	case "U":
		r.y--
	case "D":
		r.y++
	case "L":
		r.x--
	case "R":
		r.x++
	}

	if !slices.Contains(r.xValues, r.x) {
		r.xValues = append(r.xValues, r.x)
	}

	if !slices.Contains(r.yValues, r.y) {
		r.yValues = append(r.yValues, r.y)
	}

	r.Input = r.PointColors[Point{r.x, r.y}]
}

func PartOne(puzzleInput []byte) int {
	r := NewRobot(puzzleInput, Black)
	r.PaintHull()

	return len(r.paintedPoints)
}

func PartTwo(puzzleInput []byte) {
	r := NewRobot(puzzleInput, White)
	r.PaintHull()
	r.RenderHull()
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 2041
	fmt.Printf("Part Two:\n")
	PartTwo(puzzleInput) // ZRZPKEZR
}
