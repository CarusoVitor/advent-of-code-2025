package laboratories

import (
	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

type point struct {
	x int
	y int
}

func getStartingLocation(grid [][]byte) point {
	for y, line := range grid {
		for x, char := range line {
			if char == 'S' {
				return point{x: x, y: y}
			}
		}
	}
	return point{}
}

func totalSplits(grid [][]byte, start point) int {
	beams := make(map[int]struct{}, len(grid))
	beams[start.x] = struct{}{}
	splits := 0
	for yIdx := start.y + 1; yIdx < len(grid); yIdx++ {
		newBeams := make([]int, 0, 32)
		for beamX := range beams {
			if grid[yIdx][beamX] == '^' {
				newBeams = append(newBeams, beamX+1, beamX-1)
				splits++
				delete(beams, beamX)
			}
		}
		for _, beamX := range newBeams {
			beams[beamX] = struct{}{}
		}
	}
	return splits
}

func PartOne() int {
	file, err := filehandling.OpenFile("7laboratories/input.txt")
	if err != nil {
		panic(err)
	}
	grid := filehandling.ExtractGrid(file)

	start := getStartingLocation(grid)
	if start == (point{}) {
		panic("invalid starting location")
	}
	return totalSplits(grid, start)
}
