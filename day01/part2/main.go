package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuelRequired(mass int) int {
	return mass/3 - 2
}

func calculateCumulativeFuelRequired(mass int) int {
	cumulativeFuelRequired := 0
	nextStep := mass

	for nextStep != 0 {
		nextStepFuel := calculateFuelRequired(nextStep)
		if nextStepFuel < 0.0 {
			nextStepFuel = 0.0
		}
		cumulativeFuelRequired += nextStepFuel
		nextStep = nextStepFuel
	}

	return cumulativeFuelRequired
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("expected 1 argument")
	}

	filename := args[0]
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	totalFuelRequired := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		fuelRequired := calculateCumulativeFuelRequired(mass)
		totalFuelRequired += fuelRequired
	}

	fmt.Println(totalFuelRequired)
}
