package intcode

import (
	"log"
	"strconv"
	"strings"
)

type Option func(*Computer)

type Opcode int

const (
	_ Opcode = iota
	Add
	Multiply
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	RelativeBase
	Halt = Opcode(99)
)

var opcodeParameterCounts map[Opcode]int = map[Opcode]int{
	Add:          3,
	Multiply:     3,
	Input:        1,
	Output:       1,
	JumpIfTrue:   2,
	JumpIfFalse:  2,
	LessThan:     3,
	Equals:       3,
	RelativeBase: 1,
}

func opcodeLength(n Opcode) (length int) {
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
	diagnosticCode     interface{}
	phaseMode          bool
	phaseUsed          bool
	loopMode           bool
	complete           bool
	relativeBase       int
	diagnosticOutputs  []int
}

func WithPhaseMode() Option {
	return func(c *Computer) {
		c.phaseMode = true
	}
}

func WithLoopMode() Option {
	return func(c *Computer) {
		c.loopMode = true
	}
}

func (c *Computer) SetInputs(noun, verb int) {
	c.Memory[1] = noun
	c.Memory[2] = verb
}

func (c *Computer) Reset() {
	c.Memory = make([]int, len(c.program))
	copy(c.Memory, c.program)
	c.instructionPointer = 0
}

func (c *Computer) GetDiagnosticCode() interface{} {
	return c.diagnosticCode
}

func (c *Computer) GetDiagnosticOutputs() []int {
	return c.diagnosticOutputs
}

func (c *Computer) IsComplete() bool {
	return c.complete
}

func (c *Computer) evaluateParam(mode int, offset int) int {
	var returnVal = -1

	switch mode {
	case 0:
		returnVal = c.Memory[c.instructionPointer+offset]
	case 1:
		returnVal = c.instructionPointer + offset
	case 2:
		returnVal = c.relativeBase + c.Memory[c.instructionPointer+offset]
	}

	for len(c.Memory)-1 < returnVal {
		c.Memory = append(c.Memory, 0)
	}

	return returnVal
}

func (c *Computer) ExecuteProgram(inputs ...int) {
	for {
		opcode := Opcode(c.Memory[c.instructionPointer])
		if opcode == Halt {
			c.complete = true
			break
		}

		modes := make([]int, 3)
		if opcodeLength(opcode) > 1 {
			strOpcode := strconv.Itoa(int(opcode))
			oc, _ := strconv.Atoi(string(strOpcode[len(strOpcode)-1]))
			opcode = Opcode(oc)
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
		case Add:
			c.Memory[parameters[2]] = c.Memory[parameters[0]] + c.Memory[parameters[1]]
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Multiply:
			c.Memory[parameters[2]] = c.Memory[parameters[0]] * c.Memory[parameters[1]]
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Input:
			if c.phaseUsed {
				c.Memory[parameters[0]] = inputs[1]
			} else {
				if len(inputs) > 0 {
					c.Memory[parameters[0]] = inputs[0]
					if c.phaseMode {
						c.phaseUsed = true
					}
				}
			}
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Output:
			c.diagnosticCode = c.Memory[parameters[0]]
			c.diagnosticOutputs = append(c.diagnosticOutputs, c.Memory[parameters[0]])
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
			if c.loopMode {
				return
			}
		case JumpIfTrue:
			if c.Memory[parameters[0]] != 0 {
				c.instructionPointer = c.Memory[parameters[1]]
			} else {
				c.instructionPointer += opcodeParameterCounts[opcode] + 1
			}
		case JumpIfFalse:
			if c.Memory[parameters[0]] == 0 {
				c.instructionPointer = c.Memory[parameters[1]]
			} else {
				c.instructionPointer += opcodeParameterCounts[opcode] + 1
			}
		case LessThan:
			if c.Memory[parameters[0]] < c.Memory[parameters[1]] {
				c.Memory[parameters[2]] = 1
			} else {
				c.Memory[parameters[2]] = 0
			}
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case Equals:
			if c.Memory[parameters[0]] == c.Memory[parameters[1]] {
				c.Memory[parameters[2]] = 1
			} else {
				c.Memory[parameters[2]] = 0
			}
			c.instructionPointer += opcodeParameterCounts[opcode] + 1
		case RelativeBase:
			c.relativeBase += c.Memory[parameters[0]]
			c.instructionPointer += opcodeParameterCounts[opcode] + 1

		default:
			log.Printf("unexpected opcode: %v", opcode)
			c.complete = true
			return
		}
	}
}

func NewComputer(program []byte, opts ...Option) *Computer {
	programStrs := strings.Split(string(program), ",")
	programInts := make([]int, len(programStrs))

	for i, s := range programStrs {
		programInts[i], _ = strconv.Atoi(s)
	}
	c := new(Computer)
	c.program = programInts

	for _, opt := range opts {
		opt(c)
	}

	c.Reset()

	return c
}
