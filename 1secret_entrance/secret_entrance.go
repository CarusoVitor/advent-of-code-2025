package secretentrance

import (
	"strconv"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func countZeroReset(rotations []string, start int) (int, error) {
	zeroCount := 0
	acc := start
	for _, rotation := range rotations {
		direction := rotation[0]
		valueStr := rotation[1:]

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return -1, err
		}

		if direction == 'L' {
			value = -value
		}
		acc = ((acc+value)%100 + 100) % 100
		if acc == 0 {
			zeroCount++
		}
	}
	return zeroCount, nil
}

func countZeroClicks(rotations []string, start int) (int, error) {
	zeroClicks := 0
	acc := start
	for _, rotation := range rotations {
		direction := rotation[0]
		valueStr := rotation[1:]

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return -1, err
		}

		if direction == 'L' {
			value = -value
		}
		sum := (acc + value)
		processedSum := sum
		if sum == 0 {
			zeroClicks++
		} else if sum < 0 {
			processedSum = -sum
			// e.g sum is -1, then it needs to count as 101/100 = 1, but
			// if sum is -200, then it needs to count as 200/100 = 2
			if acc > 0 {
				processedSum += 100
			}
		}
		zeroClicks += processedSum / 100
		acc = (sum%100 + 100) % 100
	}
	return zeroClicks, nil
}

func PartOne() int {
	file, err := filehandling.OpenFile("1secret_entrance/input.txt")
	if err != nil {
		panic(err)
	}
	input := filehandling.ExtractSliceNewLine(file)
	res, err := countZeroReset(input, 50)
	if err != nil {
		panic(err)
	}
	return res
}

func PartTwo() int {
	file, err := filehandling.OpenFile("1secret_entrance/input.txt")
	if err != nil {
		panic(err)
	}
	input := filehandling.ExtractSliceNewLine(file)
	res, err := countZeroClicks(input, 50)
	if err != nil {
		panic(err)
	}
	return res
}
