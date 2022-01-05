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
	memory             []int
	noun               int
	verb               int
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

func (c *Computer) Initialize() {
	c.instructionPointer = 0
	c.memory[1] = c.noun
	c.memory[2] = c.verb
}

func (c *Computer) Run() {
	for {
		if !c.processOpcode() {
			break
		}

		c.step()
	}
}

func NewComputer(program []byte, noun int, verb int) *Computer {
	programStrs := strings.Split(string(program), ",")
	programInts := make([]int, len(programStrs))
	for i, s := range programStrs {
		programInts[i], _ = strconv.Atoi(s)
	}
	computer := new(Computer)
	computer.memory = programInts
	computer.noun = noun
	computer.verb = verb
	computer.Initialize()

	return computer
}

func PartOne(puzzleInput []byte) int {
	computer := NewComputer(puzzleInput, 12, 2)
	// computer.Reset()
	computer.Run()

	return computer.memory[0]
}

func main() {
	puzzleInput, _ := fs.ReadFile("day02_input.txt")

	fmt.Printf("Part One: %d\n", PartOne(puzzleInput))
}
