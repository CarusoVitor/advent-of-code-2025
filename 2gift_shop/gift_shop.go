package giftshop

import (
	"fmt"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

type idRange struct {
	lower int
	upper int
}

func maxNumWithLen(value int) int {
	if value < 1 {
		return 0
	}
	result := 1
	for range value {
		result = result * 10
	}
	return result - 1
}

func minNumWithLen(value int) int {
	if value < 1 {
		return 0
	}
	res := 1
	for i := 1; i < value; i++ {
		res *= 10
	}
	return res
}

func (id idRange) split() (*idRange, *idRange) {
	lowerMax := min(id.upper, maxNumWithLen(id.lowerSize()))
	upperMin := max(id.lower, minNumWithLen(id.upperSize()))
	if lowerMax == id.upper && upperMin == id.lower {
		return nil, nil
	}
	return &idRange{
			lower: id.lower,
			upper: lowerMax,
		}, &idRange{
			lower: upperMin,
			upper: id.upper,
		}
}

func (id idRange) String() string {
	return fmt.Sprintf("%d-%d", id.lower, id.upper)
}

func (id idRange) lowerSize() int {
	return len(fmt.Sprint(id.lower))
}

func (id idRange) upperSize() int {
	return len(fmt.Sprint(id.upper))
}

func isIdInvalid(partSize int, word string) bool {
	first := word[:partSize]
	var charIdx int
	for charIdx = partSize; charIdx < len(word); charIdx += partSize {
		if word[charIdx:charIdx+partSize] != first {
			break
		}
	}
	return charIdx == len(word)

}

func sumInvalidIdsFromRange(idRangeObj idRange, splits []int) int {
	sum := 0
	for i := idRangeObj.lower; i <= idRangeObj.upper; i++ {
		for _, split := range splits {
			partSize := idRangeObj.lowerSize() / split
			numStr := fmt.Sprint(i)
			if isIdInvalid(partSize, numStr) {
				sum += i
				break
			}
		}
	}
	return sum
}

func sumInvalidIdsPartOne(lower, upper int) int {
	idRangeObj := idRange{lower, upper}

	isLowerOdd := idRangeObj.lowerSize()%2 != 0
	isUpperOdd := idRangeObj.upperSize()%2 != 0
	if isLowerOdd && isUpperOdd {
		return 0
	}

	lowerRange, upperRange := idRangeObj.split()
	sum := 0
	if lowerRange != nil && upperRange != nil {
		if !isLowerOdd {
			sum += sumInvalidIdsFromRange(*lowerRange, []int{2})
		}
		if !isUpperOdd {
			sum += sumInvalidIdsFromRange(*upperRange, []int{2})
		}
	} else {
		sum += sumInvalidIdsFromRange(idRangeObj, []int{2})
	}
	return sum
}

func divisorsExceptOne(n int) []int {
	if n <= 1 {
		return []int{}
	}

	divs := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			// skip 1
			if i != 1 {
				divs = append(divs, i)
			}
			if i != n/i && n/i != 1 {
				divs = append(divs, n/i)
			}
		}
	}

	return divs
}

func sumInvalidIdsPartTwo(lower, upper int) int {
	idRangeObj := idRange{lower, upper}

	lowerRange, upperRange := idRangeObj.split()

	sum := 0
	if lowerRange != nil && upperRange != nil {
		sum += sumInvalidIdsFromRange(*lowerRange, divisorsExceptOne(idRangeObj.lowerSize()))
		sum += sumInvalidIdsFromRange(*upperRange, divisorsExceptOne(idRangeObj.upperSize()))
	} else {
		sum += sumInvalidIdsFromRange(idRangeObj, divisorsExceptOne(idRangeObj.upperSize()))
	}
	return sum
}

func sumInvalidIds(ranges []string, exactTwice bool) int {
	sum := 0
	for _, item := range ranges {
		values := strings.Split(item, `-`)
		lower, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		upper, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		if exactTwice {
			sum += sumInvalidIdsPartOne(lower, upper)
		} else {
			sum += sumInvalidIdsPartTwo(lower, upper)
		}

	}
	return sum

}

func PartOne() int {
	file, err := filehandling.OpenFile("2gift_shop/input.txt")
	if err != nil {
		panic(err)
	}
	input, err := filehandling.ExtractSliceSep(file, ',', true)
	if err != nil {
		panic(err)
	}
	res := sumInvalidIds(input, true)
	return res
}

func PartTwo() int {
	file, err := filehandling.OpenFile("2gift_shop/input.txt")
	if err != nil {
		panic(err)
	}
	input, err := filehandling.ExtractSliceSep(file, ',', true)
	if err != nil {
		panic(err)
	}
	res := sumInvalidIds(input, false)
	return res
}
