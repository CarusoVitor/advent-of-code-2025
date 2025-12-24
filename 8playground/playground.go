package playground

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	filehandling "github.com/CarusoVitor/advent-of-code-2025/file_handling"
)

type junctionBox struct {
	x int
	y int
	z int
}

type boxPair struct {
	first  junctionBox
	second junctionBox
}

type circuits struct {
	boxes    map[junctionBox]int
	existing map[int]struct{}
}

func (c circuits) has(box junctionBox) (int, bool) {
	num, ok := c.boxes[box]
	return num, ok
}

func (c *circuits) addIntoNewCircuit(box junctionBox) int {
	index := c.updateNewIndex()
	c.boxes[box] = index
	return index
}

func (c *circuits) updateNewIndex() int {
	newIdx := 0
	for {
		if _, ok := c.existing[newIdx]; !ok {
			break
		}
		newIdx++
	}
	c.existing[newIdx] = struct{}{}
	return newIdx
}

func (c *circuits) merge(other int, another int) {
	for box, circuit := range c.boxes {
		if circuit == another {
			c.boxes[box] = other
		}
	}
	delete(c.existing, another)
}

func (c circuits) sizes() []int {
	values := make([]int, len(c.existing))
	i := 0
	for circuit := range c.existing {
		count := 0
		for _, boxCircuit := range c.boxes {
			if boxCircuit == circuit {
				count++
			}
		}
		values[i] = count
		i++
	}
	return values
}

func (c circuits) hasConnectedEverything(goal int) bool {
	return len(c.existing) == 1 && c.sizes()[0] == goal
}

func (c *circuits) addIntoExistingCircuit(box junctionBox, existingIdx int) {
	c.boxes[box] = existingIdx
}

func newCircuits() circuits {
	boxes := make(map[junctionBox]int, 1024)
	existing := make(map[int]struct{}, 8)

	return circuits{boxes: boxes, existing: existing}
}

func newJunctionBox(coordinatesLine string) (junctionBox, error) {
	coordinates := strings.Split(coordinatesLine, ",")
	if len(coordinates) != 3 {
		return junctionBox{}, fmt.Errorf(
			"error: got %d coordinates instead of 3",
			len(coordinates),
		)
	}
	vals := make([]int, 3)
	for i, coord := range coordinates {
		val, err := strconv.Atoi(coord)
		if err != nil {
			return junctionBox{}, err
		}
		vals[i] = val
	}
	return junctionBox{vals[0], vals[1], vals[2]}, nil
}

func (jb junctionBox) distance(other junctionBox) float64 {
	return math.Sqrt(
		math.Pow(float64(jb.x-other.x), 2) +
			math.Pow(float64(jb.y-other.y), 2) +
			math.Pow(float64(jb.z-other.z), 2),
	)
}

func extractJunctionBoxes(file io.Reader) ([]junctionBox, error) {
	scanner := bufio.NewScanner(file)
	boxes := make([]junctionBox, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		box, err := newJunctionBox(line)
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, box)
	}
	return boxes, nil
}

func findMaximumDistance(pairs map[boxPair]float64) boxPair {
	max := boxPair{}
	for pair, distance := range pairs {
		if distance > pairs[max] {
			max = pair
		}
	}
	return max
}

// findNShortestPairs find the pairs that have the shortest distance
// between them.
func findNShortestPairs(boxes []junctionBox, n int) []boxPair {
	shortestConnections := make(map[boxPair]float64, len(boxes))
	maxDistancePair := boxPair{}
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			distance := boxes[i].distance(boxes[j])
			if len(shortestConnections) < n {
				shortestConnections[boxPair{boxes[i], boxes[j]}] = distance
			} else {
				if maxDistancePair == (boxPair{}) {
					maxDistancePair = findMaximumDistance(shortestConnections)
				}
				if distance < shortestConnections[maxDistancePair] {
					delete(shortestConnections, maxDistancePair)
					shortestConnections[boxPair{boxes[i], boxes[j]}] = distance
					maxDistancePair = findMaximumDistance(shortestConnections)
				}
			}

		}
	}

	return mapKeys(shortestConnections)
}

func makeCircuits(pairs []boxPair, n int) (circuits, boxPair) {
	allCircuits := newCircuits()
	lastPairConnected := boxPair{}
	for _, pair := range pairs {
		pairCircuit, hasPair := allCircuits.has(pair.first)
		pairedCircuit, hasPaired := allCircuits.has(pair.second)
		// both are in the same circuit
		if hasPair && hasPaired {
			if pairCircuit == pairedCircuit {
				continue
			} else {
				// both exists in separate circuits: needs to merge
				allCircuits.merge(pairCircuit, pairedCircuit)
			}
		} else if hasPair {
			allCircuits.addIntoExistingCircuit(pair.second, pairCircuit)
		} else if hasPaired {
			allCircuits.addIntoExistingCircuit(pair.first, pairedCircuit)
		} else {
			newCircuit := allCircuits.addIntoNewCircuit(pair.first)
			allCircuits.addIntoExistingCircuit(pair.second, newCircuit)
		}
		if allCircuits.hasConnectedEverything(n) {
			lastPairConnected = pair
			break
		}
	}

	return allCircuits, lastPairConnected
}

func productTopN(values []int, n int) int {
	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] > sorted[j] })

	product := 1
	for i := range min(n, len(values)) {
		product *= sorted[i]
	}
	return product
}

func mapKeys[K comparable, V any](values map[K]V) []K {
	slice := make([]K, len(values))
	i := 0
	for key := range values {
		slice[i] = key
		i++
	}
	return slice
}

func PartOne() int {
	file, err := filehandling.OpenFile("8playground/input.txt")
	if err != nil {
		panic(err)
	}
	boxes, err := extractJunctionBoxes(file)
	if err != nil {
		panic(err)
	}
	conns := findNShortestPairs(boxes, len(boxes))
	allCircuits, _ := makeCircuits(conns, len(boxes))
	return productTopN(allCircuits.sizes(), 3)
}

func PartTwo() int {
	file, err := filehandling.OpenFile("8playground/input.txt")
	if err != nil {
		panic(err)
	}
	boxes, err := extractJunctionBoxes(file)
	if err != nil {
		panic(err)
	}
	conns := findNShortestPairs(boxes, len(boxes)*len(boxes)-1)
	sort.Slice(conns, func(i, j int) bool {
		return conns[i].first.distance(
			conns[i].second,
		) < conns[j].first.distance(
			conns[j].second,
		)
	})
	_, lastPairConnected := makeCircuits(conns, len(boxes))
	return lastPairConnected.first.x * lastPairConnected.second.x
}
