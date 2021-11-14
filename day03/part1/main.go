package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"aoc/utils"
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

func calculateAllPositionsOnPath(p []Path) map[Point]bool {
	currentPos := Point{0, 0}
	points := make(map[Point]bool)

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
			points[currentPos] = true
		}
	}

	return points
}

func findCrossingPoints(points1, points2 map[Point]bool) []Point {
	crossingPoints := []Point{}

	for point1 := range points1 {
		if _, ok := points2[point1]; ok {
			crossingPoints = append(crossingPoints, point1)
		}
	}

	return crossingPoints
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findMinimumDistance(points []Point) int {
	if len(points) == 0 {
		return -1
	}

	minimumDistance := abs(points[0].x) + abs(points[0].y)
	for _, point := range points {
		distance := abs(point.x) + abs(point.y)
		minimumDistance = min(minimumDistance, distance)
	}

	return minimumDistance
}

func main() {
	filename := utils.ParseCommandLineArguments()
	line1, line2 := readInput(filename)
	points1, points2 := calculateAllPositionsOnPath(line1), calculateAllPositionsOnPath(line2)
	crossingPoints := findCrossingPoints(points1, points2)
	minimumDistance := findMinimumDistance(crossingPoints)
	fmt.Println(minimumDistance)
}
