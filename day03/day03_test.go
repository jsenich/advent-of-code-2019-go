package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name        string
		puzzleInput string
		want        int
	}{
		{
			name: "6",
			puzzleInput: strings.Join([]string{
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			}, "\n"),
			want: 6,
		},
		{
			name: "159",
			puzzleInput: strings.Join([]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			}, "\n"),
			want: 159,
		},
		{
			name: "135",
			puzzleInput: strings.Join([]string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			}, "\n"),
			want: 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.puzzleInput); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name        string
		puzzleInput string
		want        int
	}{
		{
			name: "610",
			puzzleInput: strings.Join([]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			}, "\n"),
			want: 610,
		},
		{
			name: "410",
			puzzleInput: strings.Join([]string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			}, "\n"),
			want: 410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.puzzleInput); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
