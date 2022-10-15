package intcode

import (
	"strconv"
	"strings"
)

var opcodeParameterCounts map[int]int = map[int]int{
	1: 3,
	2: 3,
	3: 1,
	4: 1,
}

func intLength(n int) (length int) {
	for n > 0 {
		n = n / 10
		length++
	}

	return length
}

type Computer struct {
	program            []int
	instructionPointer int
	Memory             []int
	idInput            int
	diagnosticCode     interface{}
}

func (c *Computer) Reset() {
	c.Memory = make([]int, len(c.program))
	copy(c.Memory, c.program)
	c.instructionPointer = 0
}

func (c *Computer) GetDiagnosticCode() interface{} {
	return c.diagnosticCode
}

func (c *Computer) evaluateParam(mode int, offset int) int {
	pos := c.instructionPointer + offset
	if mode == 0 {
		return c.Memory[pos]
	} else {
		return pos
	}
}

func (c *Computer) ExecuteProgram(args ...int) {
	if len(args) == 2 {
		c.Memory[1] = args[0] // noun
		c.Memory[2] = args[1] // verb
	} else if len(args) == 1 {
		c.idInput = args[0]
	}

	for {
		opcode := c.Memory[c.instructionPointer]
		if opcode == 99 {
			break
		}

		modes := make([]int, 3)
		if intLength(opcode) > 1 {
			strOpcode := strconv.Itoa(opcode)
			opcode, _ = strconv.Atoi(string(strOpcode[len(strOpcode)-1]))
			var modePos int = 0
			for i := len(strOpcode) - 3; i >= 0; i-- {
				mode, _ := strconv.Atoi(string(strOpcode[i]))
				modes[modePos] = mode
				modePos++
			}
		}

		numParams := opcodeParameterCounts[opcode]
		parameters := make([]int, numParams)
		for i := 0; i < numParams; i++ {
			parameters[i] = c.evaluateParam(modes[i], i+1)
		}

		switch opcode {
		case 1:
			c.Memory[parameters[2]] = c.Memory[parameters[0]] + c.Memory[parameters[1]]
		case 2:
			c.Memory[parameters[2]] = c.Memory[parameters[0]] * c.Memory[parameters[1]]
		case 3:
			c.Memory[parameters[0]] = c.idInput
		case 4:
			c.diagnosticCode = c.Memory[parameters[0]]
		default:
			panic("unexpected opcode")
		}

		c.instructionPointer += opcodeParameterCounts[opcode] + 1
	}
}

func NewComputer(program []byte) *Computer {
	programStrs := strings.Split(string(program), ",")
	programInts := make([]int, len(programStrs))

	for i, s := range programStrs {
		programInts[i], _ = strconv.Atoi(s)
	}
	c := new(Computer)
	c.program = programInts
	c.idInput = -1
	c.Reset()

	return c
}
