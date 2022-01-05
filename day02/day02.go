package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed *.txt
var fs embed.FS

type Computer struct {
	initialMemory      []int
	memory             []int
	instructionPointer int
}

func (c *Computer) processOpcode() bool {
	instruction := c.memory[c.instructionPointer : c.instructionPointer+4]
	opcode := instruction[0]
	parameters := instruction[1:4]
	val1 := c.memory[parameters[0]]
	val2 := c.memory[parameters[1]]

	ret := true

	switch opcode {
	case 1:
		c.memory[parameters[2]] = val1 + val2
	case 2:
		c.memory[parameters[2]] = val1 * val2
	case 99:
		ret = false
	default:
		panic("unexpected opcode")
	}

	return ret
}

func (c *Computer) step() {
	c.instructionPointer += 4
}

func (c *Computer) Reset() {
	c.memory = make([]int, len(c.initialMemory))
	copy(c.memory, c.initialMemory)
	c.instructionPointer = 0
}

func (c *Computer) ExecuteProgram(noun int, verb int) {
	if noun != -1 && verb != -1 {
		c.memory[1] = noun
		c.memory[2] = verb
	}

	for {
		if !c.processOpcode() {
			break
		}

		c.step()
	}
}

func NewComputer(program []byte) *Computer {
	programStrs := strings.Split(string(program), ",")
	programInts := make([]int, len(programStrs))

	for i, s := range programStrs {
		programInts[i], _ = strconv.Atoi(s)
	}
	computer := new(Computer)
	computer.initialMemory = programInts
	computer.Reset()

	return computer
}

func PartOne(puzzleInput []byte) int {
	computer := NewComputer(puzzleInput)
	computer.ExecuteProgram(12, 2)

	return computer.memory[0]
}

func PartTwo(puzzleInput []byte) int {
	targetValue := 19690720
	var result int

	computer := NewComputer(puzzleInput)

out:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			if noun > 0 || verb > 0 {
				computer.Reset()
			}
			computer.ExecuteProgram(noun, verb)
			if computer.memory[0] == targetValue {
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
