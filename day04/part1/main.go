package main

import "fmt"

func isValidPassword(x int) bool {
	// No six-digit check
	// No within range check

	twoAdjDigitsSame := false
	digitsIncreasing := true
	prevDigit := -1

	digits := []int{}
	for ; x > 0; x = x / 10 {
		digits = append([]int{x % 10}, digits...)
	}
	for _, digit := range digits {
		if digit < prevDigit {
			digitsIncreasing = false
		}
		if digit == prevDigit {
			twoAdjDigitsSame = true
		}
		prevDigit = digit
	}

	return twoAdjDigitsSame && digitsIncreasing
}

func main() {
	lowerBound, upperBound := 197487, 673252
	count := 0
	for i := lowerBound; i < upperBound; i++ {
		if isValidPassword(i) {
			count++
		}
	}
	fmt.Println(count)
}
