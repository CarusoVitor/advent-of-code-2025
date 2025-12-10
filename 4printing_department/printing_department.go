package printingdepartment

import (
	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func isAcessable(grid [][]byte, target byte, x, y, threshold int) bool {
	sum := 0
	width := len(grid[0])
	heigth := len(grid)
	for xIdx := x - 1; xIdx <= x+1; xIdx++ {
		for yIdx := y - 1; yIdx <= y+1; yIdx++ {
			if xIdx == x && yIdx == y || xIdx < 0 || xIdx > width-1 || yIdx < 0 || yIdx > heigth-1 {
				continue
			}
			if grid[xIdx][yIdx] == target {
				sum++
				if sum == threshold {
					return false
				}
			}
		}
	}
	return true
}

func allAcessableRolls(paperGrid [][]byte) int {
	totalSum := 0
	currentSum := -1
	for currentSum != 0 {
		currentSum = acessableRolls(paperGrid, true)
		totalSum += currentSum
	}
	return totalSum
}

func acessableRolls(paperGrid [][]byte, removeRoll bool) int {
	sum := 0
	for x, line := range paperGrid {
		for y, position := range line {
			if position == '@' && isAcessable(paperGrid, '@', x, y, 4) {
				sum++
				if removeRoll {
					paperGrid[x][y] = '.'
				}
			}
		}
	}
	return sum
}

func PartOne() int {
	file, err := filehandling.OpenFile("4printing_department/input.txt")
	if err != nil {
		panic(err)
	}
	grid := filehandling.ExtractGrid(file)
	return acessableRolls(grid, false)
}

func PartTwo() int {
	file, err := filehandling.OpenFile("4printing_department/input.txt")
	if err != nil {
		panic(err)
	}
	grid := filehandling.ExtractGrid(file)
	return allAcessableRolls(grid)
}
