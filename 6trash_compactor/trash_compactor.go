package trashcompactor

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func operation(op string, a, b int) (int, error) {
	switch op {
	case "*":
		if a == 0 {
			a = 1
		}
		a *= b
	case "+":
		a += b
	default:
		return 0, fmt.Errorf("invalid op %s", op)
	}
	return a, nil
}

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
			result[idx], err = operation(op, result[idx], num)
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil

}

func readColumnNumber(numbers []string, idx int) string {
	var num strings.Builder
	for _, line := range numbers {
		if idx > len(line)-1 {
			break
		}
		if line[idx] == ' ' {
			continue
		}
		num.WriteByte(line[idx])
	}
	return num.String()
}

func readAllColumnNumbersUntilBlank(numbers []string, idx int) ([]int, int, error) {
	currentNumbers := make([]int, len(numbers))
	for {
		numStr := readColumnNumber(numbers, idx)
		if numStr == "" {
			idx++
			break
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, 0, err
		}
		currentNumbers = append(currentNumbers, num)
		idx++
	}
	return currentNumbers, idx, nil
}

func processRighmostResult(lines []string) ([]int, error) {
	lastLine := len(lines) - 1
	operations := strings.Fields(lines[lastLine])
	result := make([]int, len(operations))
	numbers := lines[:lastLine]

	currentIdx := 0
	var currentNumbers []int
	var err error
	for idx, op := range operations {
		currentNumbers, currentIdx, err = readAllColumnNumbersUntilBlank(numbers, currentIdx)
		if err != nil {
			return nil, err
		}

		for _, num := range currentNumbers {
			finalValue, err := operation(op, result[idx], num)
			if err != nil {
				return nil, err
			}
			result[idx] = finalValue
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

	var result []int
	var err error
	result, err = processResult(lines)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, num := range result {
		sum += num
	}

	return sum
}

func applyRightMostOperations(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 4)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var result []int
	var err error
	result, err = processRighmostResult(lines)
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

func PartTwo() int {
	file, err := filehandling.OpenFile("6trash_compactor/input.txt")
	if err != nil {
		panic(err)
	}
	return applyRightMostOperations(file)
}
