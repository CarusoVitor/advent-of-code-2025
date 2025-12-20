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

// Each splitter causes an additional timeline to be created, since
// you can interpet that the beam that got split went to one of the two
// directions, and a new beam went to the other direction.
// The only problem here is that if there are two beams separated by only
// one empty space, the beams can stack. For example:
//
// .......S.......
// .......|.......
// .......^.......
// .......|.......
// .....|^X^|..... in X two beams are stacked
//
// This means that we have to keep track of stacked beams,
// since they represent different timelines.
// Here, this is done through the value (count) of the beams map, so,
// everytime a split is found, we create a new timeline for each stacked beam,
// deleting the unsplitted stacked beams
func totalTimelines(grid [][]byte, start point) int {
	beams := make(map[point]int, len(grid))
	beams[start] = 1
	timelines := 1
	for yIdx := start.y + 1; yIdx < len(grid); yIdx++ {
		for beamPoint, count := range beams {
			if grid[yIdx][beamPoint.x] != '^' {
				continue
			}

			left := point{x: beamPoint.x + 1, y: yIdx}
			right := point{x: beamPoint.x - 1, y: yIdx}

			beams[left] += count
			beams[right] += count
			timelines += count

			delete(beams, beamPoint)
		}
	}
	return timelines
}

func PartTwo() int {
	file, err := filehandling.OpenFile("7laboratories/input.txt")
	if err != nil {
		panic(err)
	}
	grid := filehandling.ExtractGrid(file)

	start := getStartingLocation(grid)
	if start == (point{}) {
		panic("invalid starting location")
	}
	return totalTimelines(grid, start)
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
