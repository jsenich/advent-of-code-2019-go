package main

import (
	"testing"
)

func TestFlattenImage(t *testing.T) {
	data := "0222112222120000"
	want := `01
10`

	got := FlattenImage(data, 2, 2)

	if got != want {
		t.Errorf("PartTwo() \ngot:\n%s\nwant:\n%s", got, want)
	}
}
