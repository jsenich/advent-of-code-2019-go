package main

import (
	"adventgo/internal/pkg/intcode"
	"embed"
	"fmt"
)

//go:embed *.txt
var fs embed.FS

func PartOne(puzzleInput []byte) int {
	computer := intcode.NewComputer(puzzleInput)
	computer.ExecuteProgram(12, 2)

	return computer.Memory[0]
}

func PartTwo(puzzleInput []byte) int {
	targetValue := 19690720
	var result int

	computer := intcode.NewComputer(puzzleInput)

out:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if noun > 0 || verb > 0 {
				computer.Reset()
			}
			computer.ExecuteProgram(noun, verb)
			if computer.Memory[0] == targetValue {
				result = 100*noun + verb
				break out
			}
		}
	}

	return result
}

func main() {
	puzzleInput, _ := fs.ReadFile("day02_input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 9706670
	fmt.Printf("Part Two: %d\n", PartTwo(puzzleInput)) // 2552
}
