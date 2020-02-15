package aocutils

import (
	"errors"
	"os"
)

func ParseCommandLineArguments() string {
	args := os.Args[1:]
	if len(args) != 1 {
		panic(errors.New("expected 1 argument"))
	}

	filename := args[0]
	return filename
}
