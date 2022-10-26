package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed *.txt
var fs embed.FS

func FlattenImage(imageData string, width, height int) string {
	layers := make([][][]string, 0)

	var layer [][]string
	var row []string
	var i, j int
	for _, d := range strings.Split(imageData, "") {
		if i == 0 {
			layer = make([][]string, height)
		}
		if j == 0 {
			row = make([]string, width)
		}
		row[j] = d
		if j+1 == width {
			layer[i] = row
			j = 0
			if i+1 == height {
				layers = append(layers, layer)
				i = 0
			} else {
				i++
			}
		} else {
			j++
		}
	}

	image := strings.Builder{}
	flattened := make([][]string, height)
	for i := range flattened {
		flattened[i] = make([]string, width)
	}

	for i := len(layers) - 1; i >= 0; i-- {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if i == len(layers)-1 {
					flattened[y][x] = layers[i][y][x]
				} else {
					if layers[i][y][x] == "2" {
						continue
					} else {
						flattened[y][x] = layers[i][y][x]
					}
				}
			}
		}
	}

	for _, row := range flattened {
		image.WriteString(strings.Join(row, "") + "\n")
	}

	return strings.TrimRight(image.String(), "\n")
}

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

func PartTwo(puzzleInput string, width, height int) string {
	output := FlattenImage(puzzleInput, width, height)
	output = strings.ReplaceAll(strings.ReplaceAll(output, "0", " "), "1", "|")

	return output
}

func main() {
	f, _ := fs.ReadFile("input.txt")
	puzzleInput := strings.Trim(string(f), "\n")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput, 25, 6))  // 1360
	fmt.Printf("Part Two:\n%s\n", PartTwo(puzzleInput, 25, 6)) // FPUAR
}
