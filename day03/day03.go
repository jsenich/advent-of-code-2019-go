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

func getWirePoints(wire []string) (points map[Point]struct{}) {
	var x, y int
	points = make(map[Point]struct{}, len(wire))
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
			points[Point{x, y}] = struct{}{}
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
	w1Points := getWirePoints(w1)
	w2Points := getWirePoints(w2)

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

func PartTwo(puzzleInput []byte) int {
	return 0
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(strings.TrimSpace(string(puzzleInput)))) // 806
	// fmt.Printf("Part Two: %d\n", PartTwo(puzzleInput))
}
