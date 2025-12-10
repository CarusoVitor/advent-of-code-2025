package cafeteria

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

type ingredientRange struct {
	lower int
	upper int
}

func (ir ingredientRange) contains(num int) bool {
	return num >= ir.lower && num <= ir.upper
}

func (ir ingredientRange) overlaps(other ingredientRange) bool {
	ours := (ir.lower >= other.lower && ir.lower <= other.upper) ||
		(ir.upper >= other.lower && ir.upper <= other.lower)
	theirs := (other.lower >= ir.lower && other.lower <= ir.upper) ||
		(other.upper >= ir.lower && other.upper <= ir.lower)
	return ours || theirs
}

func (ir ingredientRange) merge(other ingredientRange) ingredientRange {
	return ingredientRange{
		min(ir.lower, other.lower),
		max(ir.upper, other.upper),
	}
}

func (ir ingredientRange) total() int {
	return ir.upper - ir.lower + 1
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

func rangeQuantity(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	ranges, err := readRanges(scanner)
	if err != nil {
		panic(err)
	}

	sum := 0
	for idx := 0; idx < len(ranges); idx++ {
		sum += ranges[idx].total()
		for j := idx + 1; j < len(ranges); j++ {
			if !ranges[idx].overlaps(ranges[j]) {
				continue
			}
			ranges[j] = ranges[j].merge(ranges[idx])
			sum -= ranges[idx].total()
			break
		}
	}

	return sum
}

func PartOne() int {
	file, err := filehandling.OpenFile("5cafeteria/input.txt")
	if err != nil {
		panic(err)
	}
	return availableIngredients(file)
}
func PartTwo() int {
	file, err := filehandling.OpenFile("5cafeteria/input.txt")
	if err != nil {
		panic(err)
	}
	return rangeQuantity(file)
}
