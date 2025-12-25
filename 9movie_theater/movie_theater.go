package movietheater

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

type point struct {
	x int
	y int
}

func (p point) area(other point) int {
	x := p.x - other.x + 1
	if x < 0 {
		x = -x
	}
	y := p.y - other.y + 1
	if y < 0 {
		y = -y
	}
	return x * y
}

func newPoint(line string) (point, error) {
	coordinates := strings.Split(line, ",")
	if len(coordinates) != 2 {
		return point{}, fmt.Errorf(
			"error: got %d coordinates instead of 2",
			len(coordinates),
		)
	}
	vals := make([]int, 2)
	for i, coord := range coordinates {
		val, err := strconv.Atoi(coord)
		if err != nil {
			return point{}, err
		}
		vals[i] = val
	}
	return point{vals[0], vals[1]}, nil
}

func extractCoordinates(file io.Reader) ([]point, error) {
	scanner := bufio.NewScanner(file)
	points := make([]point, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		box, err := newPoint(line)
		if err != nil {
			return nil, err
		}
		points = append(points, box)
	}
	return points, nil
}
func findBiggestRectangleArea(points []point) int {
	area := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			currentArea := points[i].area(points[j])
			if currentArea > area {
				area = currentArea
			}
		}
	}
	return area
}
func PartOne() int {
	file, err := filehandling.OpenFile("9movie_theater/input.txt")
	if err != nil {
		panic(err)
	}
	coords, err := extractCoordinates(file)
	if err != nil {
		panic(err)
	}
	biggestArea := findBiggestRectangleArea(coords)
	return biggestArea
}
