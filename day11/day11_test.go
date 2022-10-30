package main

import "testing"

func TestGetNewDirection(t *testing.T) {
	type args struct {
		currentDirection string
		deg              int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"left from up",
			args{currentDirection: "U", deg: 0},
			"L",
		},
		{
			"right from up",
			args{currentDirection: "U", deg: 1},
			"R",
		},
		{
			"right from left",
			args{currentDirection: "l", deg: 1},
			"U",
		},
		{
			"right from down",
			args{currentDirection: "D", deg: 1},
			"L",
		},
		{
			"left from down",
			args{currentDirection: "D", deg: 0},
			"R",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNewDirection(tt.args.currentDirection, tt.args.deg); got != tt.want {
				t.Errorf("GetNewDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
