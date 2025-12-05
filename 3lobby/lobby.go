package lobby

import (
	"strconv"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func getJoltage(bank string) (int, error) {
	var left byte = 0
	var right byte = 0

	for i := 0; i < len(bank); i++ {
		num := bank[i]
		if num > left && i < len(bank)-1 {
			left = num
			right = bank[i+1]
		} else if num > right {
			right = num
		}
	}
	return strconv.Atoi(string([]byte{left, right}))
}

func outputJoltage(batteries []string) int {
	joltage := 0
	for _, bank := range batteries {
		bankJoltage, err := getJoltage(bank)

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
	return outputJoltage(input)
}
