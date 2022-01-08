package main

import (
	"embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed *.txt
var fs embed.FS

type Point struct {
	X, Y int
}

func pointSliceToMap(points []Point) (set map[Point]struct{}) {
	set = make(map[Point]struct{}, len(points))

	for _, p := range points {
		set[p] = struct{}{}
	}

	return set
}

func getWirePoints(wire []string) (points []Point) {
	var x, y int // points = make(map[Point]struct{}, len(wire))
	points = make([]Point, 0, len(wire))
	for _, move := range wire {
		direction := string(move[0])

		distance, _ := strconv.Atoi(string(move[1:]))
		for i := 1; i <= distance; i++ {
			switch direction {
			case "R":
				x += 1
			case "L":
				x -= 1
			case "U":
				y += 1
			case "D":
				y -= 1
			}

			points = append(points, Point{x, y})
		}
	}

	return points
}

func getDistance(p1 Point, p2 Point) int {
	distance := math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y))

	return int(distance)
}

func parsePuzzleInput(puzzleInput string) (wire1 []string, wire2 []string) {
	wires := strings.Split(strings.TrimSpace(string(puzzleInput)), "\n")
	wire1 = strings.Split(wires[0], ",")
	wire2 = strings.Split(wires[1], ",")

	return wire1, wire2
}

func PartOne(puzzleInput string) int {
	w1, w2 := parsePuzzleInput(puzzleInput)
	w1Points := pointSliceToMap(getWirePoints(w1))
	w2Points := pointSliceToMap(getWirePoints(w2))

	intersectionDistances := []int{}
	for p := range w2Points {
		if _, ok := w1Points[p]; ok {
			intersectionDistances = append(intersectionDistances, getDistance(Point{0, 0}, p))
		}
	}

	sort.SliceStable(intersectionDistances, func(i, j int) bool {
		return intersectionDistances[i] < intersectionDistances[j]
	})

	return intersectionDistances[0]
}

func PartTwo(puzzleInput string) int {
	return 0
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(strings.TrimSpace(string(puzzleInput)))) // 806
	// fmt.Printf("Part Two: %d\n", PartTwo(puzzleInput))
}
