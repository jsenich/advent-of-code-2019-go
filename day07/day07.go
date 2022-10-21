package main

import (
	"adventgo/internal/pkg/intcode"
	"embed"
	"fmt"
	"sort"

	"gonum.org/v1/gonum/stat/combin"
)

//go:embed *.txt
var fs embed.FS

func PartOne(puzzleInput []byte) int {
	var sequences []int
	amplifiers := make([]intcode.Computer, 5)
	p := combin.Permutations(5, 5)

	for _, sequence := range p {
		var inputSignal int

		for i := 0; i < 5; i++ {
			amplifiers[i] = *intcode.NewComputer(puzzleInput)
			amplifiers[i].ExecuteProgram(true, sequence[i], inputSignal)
			inputSignal = amplifiers[i].GetDiagnosticCode().(int)
			if i == 4 {
				sequences = append(sequences, inputSignal)
			}
		}
	}

	sort.Ints(sequences)

	return sequences[len(sequences)-1]
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 18812
}
