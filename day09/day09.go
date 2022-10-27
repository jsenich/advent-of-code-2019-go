package main

import (
	"adventgo/internal/pkg/intcode"
	"embed"
	"fmt"
)

//go:embed *.txt
var fs embed.FS

func PartOne(puzzleInput []byte) int {
	c := intcode.NewComputer(puzzleInput)
	c.ExecuteProgram(1)

	return c.GetDiagnosticCode().(int)
}

func PartTwo(puzzleInput []byte) int {
	c := intcode.NewComputer(puzzleInput)
	c.ExecuteProgram(2)

	return c.GetDiagnosticCode().(int)
}

func main() {
	puzzleInput, _ := fs.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput)) // 2351176124
	fmt.Printf("Part Two: %d\n", PartTwo(puzzleInput)) // 73110
}
