package foo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Path struct {
	Direction string
	Distance  int
}

func parseLine(s string) []Path {
	split := strings.Split(s, ",")
	path := make([]Path, len(split))
	for i, elem := range split {
		distance, _ := strconv.Atoi(string(elem[1:]))
		path[i] = Path{
			Direction: string(elem[0]),
			Distance:  distance,
		}
	}
	return path
}

func readInput(filename string) ([]Path, []Path) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputStrings := strings.Split(inputText, "\n")
	path1, path2 := parseLine(inputStrings[0]), parseLine(inputStrings[1])

	return path1, path2
}

func parseCommandLineArguments() string {
	args := os.Args[1:]
	if len(args) != 1 {
		panic(errors.New("expected 1 argument"))
	}

	filename := args[0]
	return filename
}

func calculateAllPositionsOnPath(p []Path) map[Point]int {
	currentPos := Point{0, 0}
	currentDistance := 0
	points := make(map[Point]int)

	speedX, speedY := 0, 0
	for _, pathElem := range p {
		switch pathElem.Direction {
		case "U":
			speedX, speedY = 0, 1
		case "D":
			speedX, speedY = 0, -1
		case "L":
			speedX, speedY = -1, 0
		case "R":
			speedX, speedY = 1, 0
		default:
			fmt.Println("error: unexpected direction")
		}

		for i := 0; i < pathElem.Distance; i++ {
			currentPos.x += speedX
			currentPos.y += speedY
			currentDistance++
			if _, ok := points[currentPos]; !ok {
				points[currentPos] = currentDistance
			}
		}
	}

	return points
}

func findMinimalCrossingPoints(points1, points2 map[Point]int) *int {
	var minDistance *int

	for point1 := range points1 {
		if _, ok := points2[point1]; ok {
			distance := points1[point1] + points2[point1]
			if minDistance == nil {
				minDistance = &distance
			}
			distance = min(distance, *minDistance)
			minDistance = &distance
		}
	}

	return minDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	filename := parseCommandLineArguments()
	line1, line2 := readInput(filename)
	points1, points2 := calculateAllPositionsOnPath(line1), calculateAllPositionsOnPath(line2)
	minimumDistance := findMinimalCrossingPoints(points1, points2)
	fmt.Println(*minimumDistance)
}
