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
		result = result*10 + 9
	}
	return result
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
	lowerMax := min(id.upper, maxNumWithLen(len(fmt.Sprint(id.lower))))
	upperMin := max(id.lower, minNumWithLen(len(fmt.Sprint(id.upper))))
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

func sumInvalidIdsFromRange(idRangeObj idRange) int {
	sum := 0
	halfSize := idRangeObj.lowerSize() / 2
	for i := idRangeObj.lower; i <= idRangeObj.upper; i++ {
		left := fmt.Sprint(i)[:halfSize]
		right := fmt.Sprint(i)[halfSize:]
		if left == right {
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
			sum += sumInvalidIdsFromRange(*lowerRange)
		}
		if !isUpperOdd {
			sum += sumInvalidIdsFromRange(*upperRange)
		}
	} else {
		sum += sumInvalidIdsFromRange(idRangeObj)
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
