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

func calculate_fuel_requirement(mass int, include_additional_fuel bool) int {
	var requirement float64

	if !include_additional_fuel {
		requirement = math.Floor(float64(mass)/3) - 2
	} else {
		var totalRequirement float64
		for {
			requirement = math.Floor(float64(mass)/3) - 2
			if requirement <= 0 {
				break
			}
			totalRequirement += requirement
			mass = int(requirement)
		}
		requirement = totalRequirement
	}

	return int(requirement)
}

func part_one(data string) int {
	var fuelTotal int
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		requirement := calculate_fuel_requirement(mass, false)
		fuelTotal += int(requirement)
	}
	return fuelTotal
}

func part_two(data string) int {
	var fuelTotal int
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		requirement := calculate_fuel_requirement(mass, true)
		fuelTotal += int(requirement)
	}
	return fuelTotal
}

func main() {
	puzzleInput, _ := data.ReadFile("data/day01_input.txt")

	fmt.Printf("Part One: %d\n", part_one(string(puzzleInput))) // 3349352
	fmt.Printf("Part Two: %d\n", part_two(string(puzzleInput))) // 5021154
}
