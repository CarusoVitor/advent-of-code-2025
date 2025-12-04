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

func sumInvalidIdsFromRange(idRangeObj idRange, splits int) int {
	sum := 0
	partSize := idRangeObj.lowerSize() / splits

	for i := idRangeObj.lower; i <= idRangeObj.upper; i++ {
		numStr := fmt.Sprint(i)

		first := numStr[:partSize]
		var charIdx int
		for charIdx = partSize; charIdx < len(numStr); charIdx += partSize {
			if numStr[charIdx:charIdx+partSize] != first {
				break
			}
		}
		if charIdx == len(numStr) {
			sum += i
		}
	}
	return sum
}

func sumInvalidIdsFromLimits(lower, upper int) int {
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
			sum += sumInvalidIdsFromRange(*lowerRange, 2)
		}
		if !isUpperOdd {
			sum += sumInvalidIdsFromRange(*upperRange, 2)
		}
	} else {
		sum += sumInvalidIdsFromRange(idRangeObj, 2)
	}
	return sum
}

func sumInvalidIds(ranges []string) int {
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
		sum += sumInvalidIdsFromLimits(lower, upper)

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
	res := sumInvalidIds(input)
	return res
}
