package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ParameterMode int
type Opcode int

const (
	Position  = 0
	Immediate = 1
)

const (
	Add         = 1
	Multiply    = 2
	GetInput    = 3
	Output      = 4
	JumpIfTrue  = 5
	JumpIfFalse = 6
	LessThan    = 7
	Equals      = 8
	Halt        = 99
)

type Instruction struct {
	opcode                             Opcode
	param1Mode, param2Mode, param3Mode ParameterMode
}

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

func parseInstruction(x int) Instruction {
	// 01234
	// ABCDE
	//  1002

	digits := []int{0, 0, 0, 0, 0}
	digitCounter := 4
	for ; x > 0; x /= 10 {
		digits[digitCounter] = x % 10
		digitCounter--
	}

	return Instruction{
		opcode:     Opcode(digits[3]*10 + digits[4]),
		param1Mode: ParameterMode(digits[2]),
		param2Mode: ParameterMode(digits[1]),
		param3Mode: ParameterMode(digits[0]),
	}
}

func getParams(i Instruction, nums []int, instructionPtr int) (int, int, int) {
	var param1Val, param2Val, param3Val int

	if i.param1Mode == Position {
		param1Val = nums[nums[instructionPtr+1]]
	} else if i.param1Mode == Immediate {
		param1Val = nums[instructionPtr+1]
	}

	if i.param2Mode == Position {
		param2Val = nums[nums[instructionPtr+2]]
	} else if i.param2Mode == Immediate {
		param2Val = nums[instructionPtr+2]
	}

	param3Val = nums[instructionPtr+3]

	return param1Val, param2Val, param3Val
}

func process(nums []int) (int, []int) {
	inputs, outputs := []int{5}, []int{}
	inputCounter := 0

	for instructionPtr := 0; instructionPtr < len(nums); {
		i := parseInstruction(nums[instructionPtr])
		switch i.opcode {
		case Add:
			param1Val, param2Val, param3Val := getParams(i, nums, instructionPtr)
			nums[param3Val] = param1Val + param2Val
			instructionPtr += 4

		case Multiply:
			param1Val, param2Val, param3Val := getParams(i, nums, instructionPtr)
			nums[param3Val] = param1Val * param2Val
			instructionPtr += 4

		case GetInput:
			val := inputs[inputCounter]
			inputCounter++

			writePos := nums[instructionPtr+1]
			nums[writePos] = val

			instructionPtr += 2

		case Output:
			readPos := nums[instructionPtr+1]
			outputs = append(outputs, nums[readPos])
			instructionPtr += 2

		case JumpIfTrue:
			param1Val, param2Val, _ := getParams(i, nums, instructionPtr)

			if param1Val != 0 {
				instructionPtr = param2Val
			} else {
				instructionPtr += 3
			}

		case JumpIfFalse:
			param1Val, param2Val, _ := getParams(i, nums, instructionPtr)

			if param1Val == 0 {
				instructionPtr = param2Val
			} else {
				instructionPtr += 3
			}

		case LessThan:
			param1Val, param2Val, param3Val := getParams(i, nums, instructionPtr)

			if param1Val < param2Val {
				nums[param3Val] = 1
			} else {
				nums[param3Val] = 0
			}

			instructionPtr += 4

		case Equals:
			param1Val, param2Val, param3Val := getParams(i, nums, instructionPtr)

			if param1Val == param2Val {
				nums[param3Val] = 1
			} else {
				nums[param3Val] = 0
			}

			instructionPtr += 4

		case Halt:
			return nums[0], outputs

		default:
			return -1, []int{}
		}
	}

	return -1, []int{}
}

func main() {
	filename := parseCommandLineArguments()
	nums := readInput(filename)
	result, outputs := process(nums)
	fmt.Println("result:", result, "outputs:", outputs)
}
