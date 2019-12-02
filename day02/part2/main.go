package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func parseCommandLineArguments() string {
	args := os.Args[1:]
	if len(args) != 1 {
		panic(errors.New("expected 1 argument"))
	}

	filename := args[0]
	return filename
}

func process(nums []int, noun, verb int) int {
	nums[1] = noun
	nums[2] = verb

	for instructionPtr := 0; (instructionPtr + 3) < len(nums); instructionPtr += 4 {
		opcode := nums[instructionPtr]
		switch opcode {
		case 1:
			readPos1, readPos2, writePos := nums[instructionPtr+1], nums[instructionPtr+2], nums[instructionPtr+3]
			nums[writePos] = nums[readPos1] + nums[readPos2]
		case 2:
			readPos1, readPos2, writePos := nums[instructionPtr+1], nums[instructionPtr+2], nums[instructionPtr+3]
			nums[writePos] = nums[readPos1] * nums[readPos2]
		case 99:
			return nums[0]
		default:
			return -1
		}
	}
	return -1
}

func calculateNounAndVerb(nums []int, target int) (noun, verb int) {
	for noun = 0; noun < 99; noun++ {
		for verb = 0; verb < 99; verb++ {
			numsCopy := make([]int, len(nums))
			copy(numsCopy, nums)
			result := process(numsCopy, noun, verb)
			if result == target {
				return
			}
		}
	}
	return -1, -1
}

func main() {
	filename := parseCommandLineArguments()
	nums := readInput(filename)
	target := 19690720
	noun, verb := calculateNounAndVerb(nums, target)
	result := 100*noun + verb
	fmt.Println(result)
}
