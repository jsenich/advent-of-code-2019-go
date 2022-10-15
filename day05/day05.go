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
	computer.ExecuteProgram(1)

	return computer.GetDiagnosticCode().(int)
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 9938601
}
