package intcode

import (
	"testing"
)

func TestComputer_ExecuteProgram(t *testing.T) {
	tests := []struct {
		program    string
		noun, verb int
		want       int
	}{
		{"1,0,0,0,99", -1, -1, 2},
		{"2,3,0,3,99", -1, -1, 2},
		{"2,4,4,5,99,0", -1, -1, 2},
		{"1,1,1,4,99,5,6,0,99", -1, -1, 30},
	}
	for _, tt := range tests {
		t.Run(tt.program, func(t *testing.T) {
			c := NewComputer([]byte(tt.program))
			c.ExecuteProgram(tt.noun, tt.verb)
			got := c.Memory[0]
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
