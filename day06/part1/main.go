package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"aoc/utils"
)

type Object struct {
	Name     string
	Children []Object
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
				Name:     obj1,
				Children: []Object{},
			}
		}
		if _, ok := objects[obj2]; !ok {
			objects[obj2] = &Object{
				Name:     obj2,
				Children: []Object{},
			}
		}

		objects[obj1].Children = append(objects[obj1].Children, *objects[obj2])
	}

	return objects
}

func countOrbitsInternal(objects map[string]*Object, current string, height int) int {
	fmt.Println("countOrbitsInternal", "current=", current, "height=", height)
	count := 0
	for _, child := range objects[current].Children {
		count += countOrbitsInternal(objects, child.Name, height+1)
	}
	fmt.Println("count=", count)
	return height + count
}

func countOrbits(objects map[string]*Object) int {
	return countOrbitsInternal(objects, "COM", 0)
}

func main() {
	filename := utils.ParseCommandLineArguments()
	objects := readInput(filename)
	for name, object := range objects {
		fmt.Println(name, object)
	}
	result := countOrbits(objects)
	fmt.Println(result)
}
