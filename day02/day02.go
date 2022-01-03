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
	program         []int
	currentPosition int
}

func (c *Computer) processOpcode() bool {
	op := c.program[c.currentPosition : c.currentPosition+4]
	opcode := op[0]
	val1 := c.program[op[1]]
	val2 := c.program[op[2]]

	ret := true

	switch opcode {
	case 1:
		c.program[op[3]] = val1 + val2
	case 2:
		c.program[op[3]] = val1 * val2
	case 99:
		ret = false
	default:
		panic("unexpected opcode")
	}

	return ret
}

func (c *Computer) step() {
	c.currentPosition += 4
}

func (c *Computer) Initialize() {
	c.currentPosition = 0
	c.program[1] = 12
	c.program[2] = 2
}

func (c *Computer) Run() {
	for {
		if !c.processOpcode() {
			break
		}

		c.step()
	}
}

func NewComputer(program []byte) Computer {
	programStrs := strings.Split(string(program), ",")
	programInts := make([]int, len(programStrs))
	for i, s := range programStrs {
		programInts[i], _ = strconv.Atoi(s)
	}

	return Computer{program: programInts}
}

func part_one(puzzleInput []byte) int {
	computer := NewComputer(puzzleInput)
	computer.Initialize()
	computer.Run()

	return computer.program[0]
}

func main() {
	puzzleInput, _ := fs.ReadFile("day02_input.txt")

	fmt.Printf("Part One: %d\n", part_one(puzzleInput))
}
