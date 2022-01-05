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
		opcode := c.Memory[c.instructionPointer]
		if opcode == 99 {
			break
		}

		parameters := c.Memory[c.instructionPointer+1 : c.instructionPointer+4]
		val1 := c.Memory[parameters[0]]
		val2 := c.Memory[parameters[1]]

		switch opcode {
		case 1:
			c.Memory[parameters[2]] = val1 + val2
		case 2:
			c.Memory[parameters[2]] = val1 * val2
		default:
			panic("unexpected opcode")
		}

		c.instructionPointer += 4
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
