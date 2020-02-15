package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Object struct {
	Name  string
	Links []Object
}

func readInput(filename string) map[string]*Object {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	objects := map[string]*Object{}

	inputText := string(data)

	lines := strings.Split(strings.TrimSpace(inputText), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ")")
		if len(parts) != 2 {
			fmt.Println("line=", line)
			log.Fatalln("unexpected line")
		}

		obj1, obj2 := parts[0], parts[1]
		if _, ok := objects[obj1]; !ok {
			objects[obj1] = &Object{
				Name:  obj1,
				Links: []Object{},
			}
		}
		if _, ok := objects[obj2]; !ok {
			objects[obj2] = &Object{
				Name:  obj2,
				Links: []Object{},
			}
		}

		objects[obj1].Links = append(objects[obj1].Links, *objects[obj2])
		objects[obj2].Links = append(objects[obj2].Links, *objects[obj1])
	}

	return objects
}

func parseCommandLineArguments() string {
	args := os.Args[1:]
	if len(args) != 1 {
		panic(errors.New("expected 1 argument"))
	}

	filename := args[0]
	return filename
}

type QueueItem struct {
	name     string
	distance int
}

func findSanta(objects map[string]*Object) (int, error) {
	queue := []QueueItem{{"YOU", 0}}
	visited := map[string]bool{}
	target := "SAN"

	for len(queue) > 0 {
		var current QueueItem
		current, queue = queue[0], queue[1:]

		if current.name == target {
			return current.distance - 2, nil
		}

		if _, ok := visited[current.name]; ok {
			continue
		}
		visited[current.name] = true

		for _, link := range objects[current.name].Links {
			queue = append(queue, QueueItem{link.Name, current.distance + 1})
		}
	}

	return 0, errors.New("unable to find santa")
}

func main() {
	filename := parseCommandLineArguments()
	objects := readInput(filename)
	result, err := findSanta(objects)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
