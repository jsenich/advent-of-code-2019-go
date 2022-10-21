package main

import (
	"adventgo/internal/pkg/intcode"
	"embed"
	"fmt"
	"sort"
)

//go:embed *.txt
var fs embed.FS

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func PartOne(puzzleInput []byte) int {
	var sequences []int
	amplifiers := make([]intcode.Computer, 5)
	p := permutation([]int{0, 1, 2, 3, 4})

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
