package main

import (
	"bufio"
	"embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed data/day01_input.txt
var data embed.FS

func part_one(data string) int {
	var fuelTotal int
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		requirement := math.Floor(float64(mass)/3) - 2
		fuelTotal += int(requirement)
	}
	return fuelTotal
}

func part_two(data string) int {
	var fuelTotal int
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		var requirement float64
		for mass > 0 {
			requirement = math.Floor(float64(mass)/3) - 2
			if requirement > 0 {
				fuelTotal += int(requirement)
			}
			mass = int(requirement)
		}
		// fuelTotal += int(requirement)
	}
	return fuelTotal
}

func main() {
	puzzleInput, _ := data.ReadFile("data/day01_input.txt")

	// var fuelTotal int
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	mass, _ := strconv.Atoi(scanner.Text())
	// 	requirement := math.Floor(float64(mass)/3) - 2
	// 	fuelTotal += int(requirement)
	// }
	// fmt.Printf("Part One: %d", part_one(string(puzzleInput)))
	fmt.Printf("Part Two: %d", part_two(string(puzzleInput)))
}
