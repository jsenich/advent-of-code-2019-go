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
	programInput := 1
	computer.ExecuteProgram(programInput)

	return computer.GetDiagnosticCode().(int)
}

func PartTwo(puzzleInput []byte) int {
	computer := intcode.NewComputer(puzzleInput)
	programInput := 5
	computer.ExecuteProgram(programInput)

	return computer.GetDiagnosticCode().(int)
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 9938601
	fmt.Printf("Part Two: %d\n", PartTwo(puzzleInput)) // 4283952
}
