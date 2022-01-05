package intcode

import (
	"strconv"
	"strings"
)

type Computer struct {
	initialMemory      []int
	instructionPointer int
	Memory             []int
}

func (c *Computer) processOpcode() bool {
	instruction := c.Memory[c.instructionPointer : c.instructionPointer+4]
	opcode := instruction[0]
	parameters := instruction[1:4]
	val1 := c.Memory[parameters[0]]
	val2 := c.Memory[parameters[1]]

	ret := true

	switch opcode {
	case 1:
		c.Memory[parameters[2]] = val1 + val2
	case 2:
		c.Memory[parameters[2]] = val1 * val2
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
	c.Memory = make([]int, len(c.initialMemory))
	copy(c.Memory, c.initialMemory)
	c.instructionPointer = 0
}

func (c *Computer) ExecuteProgram(noun int, verb int) {
	if noun != -1 && verb != -1 {
		c.Memory[1] = noun
		c.Memory[2] = verb
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
