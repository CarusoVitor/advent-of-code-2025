package trashcompactor

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func processResult(lines [][]string) ([]int, error) {
	quantity := len(lines[0])
	lastLine := len(lines) - 1
	result := make([]int, quantity)
	for idx := range quantity {
		op := lines[lastLine][idx]
		for numIdx := range lastLine {
			numStr := lines[numIdx][idx]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			switch op {
			case "*":
				if result[idx] == 0 {
					result[idx] = 1
				}
				result[idx] *= num
			case "+":
				result[idx] += num
			default:
				return nil, fmt.Errorf("invalid op %s", op)
			}
		}
	}
	return result, nil

}

func applyOperations(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	lines := make([][]string, 0, 4)
	for scanner.Scan() {
		lines = append(lines, strings.Fields(scanner.Text()))
	}

	result, err := processResult(lines)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, num := range result {
		sum += num
	}

	return sum
}

func PartOne() int {
	file, err := filehandling.OpenFile("6trash_compactor/input.txt")
	if err != nil {
		panic(err)
	}
	return applyOperations(file)
}
