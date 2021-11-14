package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"aoc/utils"
)

func readInput(filename string) []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputText = strings.TrimSuffix(inputText, "\n")
	inputStrings := strings.Split(inputText, ",")

	nums := make([]int, len(inputStrings))
	for i, inputString := range inputStrings {
		nums[i], _ = strconv.Atoi(inputString)
	}
	return nums
}

func process(nums []int) int {
	// restore the gravity assist program to the "1202 program alarm"
	nums[1] = 12
	nums[2] = 2

	for counter := 0; (counter + 3) < len(nums); counter += 4 {
		switch nums[counter] {
		case 1:
			readPos1, readPos2, writePos := nums[counter+1], nums[counter+2], nums[counter+3]
			nums[writePos] = nums[readPos1] + nums[readPos2]
		case 2:
			readPos1, readPos2, writePos := nums[counter+1], nums[counter+2], nums[counter+3]
			nums[writePos] = nums[readPos1] * nums[readPos2]
		case 99:
			return nums[0]
		default:
			fmt.Printf("error: unexpected opcode: %d\n", nums[counter])
		}
	}

	return -1
}

func main() {
	filename := utils.ParseCommandLineArguments()

	nums := readInput(filename)
	result := process(nums)
	fmt.Println(result)
}
