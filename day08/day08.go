package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed *.txt
var fs embed.FS

func PartOne(puzzleInput string, width, height int) int {
	layers := make([]string, 0)
	layerSize := width * height

	layerCounts := map[int]int{}
	var zerosIndex int
	var minZerosCount int = -1

	sb := strings.Builder{}
	var i int
	for _, d := range strings.Split(puzzleInput, "") {
		sb.WriteString(d)
		if sb.Len() == layerSize {
			s := sb.String()
			layers = append(layers, s)
			layerCounts[i] = strings.Count(s, "0")
			i++
			sb.Reset()
		}
	}

	for k, v := range layerCounts {
		if minZerosCount == -1 {
			minZerosCount = v
			zerosIndex = k
			continue
		}
		if v < minZerosCount {
			zerosIndex = k
			minZerosCount = v
		}
	}

	l := layers[zerosIndex]

	return strings.Count(l, "1") * strings.Count(l, "2")
}

func main() {
	f, _ := fs.ReadFile("input.txt")
	puzzleInput := strings.Trim(string(f), "\n")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput, 25, 6)) // 1360
}
