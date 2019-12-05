package main

import "fmt"

func isValidPassword(x int) bool {
	// No six-digit check
	// No within range check

	digitsIncreasing := true
	prevDigit := -1

	digits := []int{}
	m := map[int]int{}
	for ; x > 0; x = x / 10 {
		digits = append([]int{x % 10}, digits...)
	}
	for _, digit := range digits {
		if digit < prevDigit {
			digitsIncreasing = false
		}
		if digit == prevDigit {
			if _, ok := m[digit]; ok {
				m[digit]++
			} else {
				m[digit] = 2
			}
		}
		prevDigit = digit
	}

	atLeastOneDouble := false
	for _, v := range m {
		if v == 2 {
			atLeastOneDouble = true
		}
	}
	//fmt.Println(m)

	return atLeastOneDouble && digitsIncreasing
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
	/*fmt.Println(isValidPassword(112233))
	fmt.Println(isValidPassword(123444))
	fmt.Println(isValidPassword(111122))*/
}
