package lobby

import (
	"strconv"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func getJoltage(bank string, size int) (int, error) {
	joltage := make([]byte, size)
	nextIdx := 0

	for batteryIdx := 0; batteryIdx < len(bank); batteryIdx++ {
		battery := bank[batteryIdx]
		batteriesLeft := min(len(bank)-batteryIdx, size)
		joltageIdx := size - batteriesLeft

		for joltageIdx < nextIdx {
			if battery > joltage[joltageIdx] {
				joltage[joltageIdx] = battery
				nextIdx = joltageIdx + 1
				break
			}
			joltageIdx++
		}
		if joltageIdx == nextIdx && nextIdx < size {
			joltage[nextIdx] = battery
			nextIdx++
		}
	}

	return strconv.Atoi(string(joltage))
}

func outputJoltage(batteries []string, size int) int {
	joltage := 0
	for _, bank := range batteries {
		bankJoltage, err := getJoltage(bank, size)

		if err != nil {
			panic(err)
		}
		joltage += bankJoltage
	}
	return joltage
}

func PartOne() int {
	file, err := filehandling.OpenFile("3lobby/input.txt")
	if err != nil {
		panic(err)
	}
	input := filehandling.ExtractSliceNewLine(file)
	return outputJoltage(input, 2)
}

func PartTwo() int {
	file, err := filehandling.OpenFile("3lobby/input.txt")
	if err != nil {
		panic(err)
	}
	input := filehandling.ExtractSliceNewLine(file)
	return outputJoltage(input, 12)
}
