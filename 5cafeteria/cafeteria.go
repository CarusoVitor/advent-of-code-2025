package cafeteria

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

func PartOne() int {
	file, err := filehandling.OpenFile("5cafeteria/input.txt")
	if err != nil {
		panic(err)
	}
	return availableIngredients(file)
}

type ingredientRange struct {
	lower int
	upper int
}

func (ir ingredientRange) contains(num int) bool {
	return num >= ir.lower && num <= ir.upper
}

func newIngredientRange(line string) (ingredientRange, error) {
	lineSlice := strings.Split(line, "-")
	lower, err := strconv.Atoi(lineSlice[0])
	if err != nil {
		return ingredientRange{}, err
	}
	upper, err := strconv.Atoi(lineSlice[1])
	if err != nil {
		return ingredientRange{}, err
	}
	return ingredientRange{lower: lower, upper: upper}, nil
}

func readRanges(scanner *bufio.Scanner) ([]ingredientRange, error) {
	ranges := make([]ingredientRange, 0, 512)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ingredientRangeObj, err := newIngredientRange(line)
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, ingredientRangeObj)
	}
	return ranges, nil
}

func isAvailable(ingredient int, ranges []ingredientRange) bool {
	for _, r := range ranges {
		if r.contains(ingredient) {
			return true
		}
	}
	return false
}

func availableIngredients(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	ranges, err := readRanges(scanner)
	if err != nil {
		panic(err)
	}
	sum := 0
	for scanner.Scan() {
		ingredient, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		if isAvailable(ingredient, ranges) {
			sum++
		}
	}
	return sum
}
