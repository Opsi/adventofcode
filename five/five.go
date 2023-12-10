package five

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Opsi/adventofcode/util"
)

type Input struct {
	Seeds []int
	Maps  []Map
}

type Map []Range

type Range struct {
	DestStart int
	SrcStart  int
	Length    int
}

func (m Map) Apply(i int) int {
	for _, r := range m {
		if i >= r.SrcStart && i < r.SrcStart+r.Length {
			return r.DestStart + (i - r.SrcStart)
		}
	}
	return i
}

func parseRange(s string) (Range, error) {
	var r Range
	var err error
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return r, fmt.Errorf("invalid range: %s", s)
	}
	r.DestStart, err = strconv.Atoi(parts[0])
	if err != nil {
		return r, fmt.Errorf("invalid dest start: %v", err)
	}
	r.SrcStart, err = strconv.Atoi(parts[1])
	if err != nil {
		return r, fmt.Errorf("invalid src start: %v", err)
	}
	r.Length, err = strconv.Atoi(parts[2])
	if err != nil {
		return r, fmt.Errorf("invalid length: %v", err)
	}
	return r, nil
}

func parseMap(lines []string) (Map, error) {
	var m Map
	for _, line := range lines {
		newRange, err := parseRange(line)
		if err != nil {
			return m, fmt.Errorf("parse range: %v", err)
		}
		m = append(m, newRange)
	}
	return m, nil
}

const seedsPrefix = "seeds: "
const mapSuffix = "map:"

func parseInput(lines []string) (Input, error) {
	var input Input
	var err error
	if len(lines) < 2 {
		return input, fmt.Errorf("invalid input: %v", lines)
	}
	if !strings.HasPrefix(lines[0], seedsPrefix) {
		return input, fmt.Errorf("invalid seeds: %s", lines[0])
	}
	input.Seeds, err = util.ParseSpaceSeparatedInts(strings.TrimPrefix(lines[0], seedsPrefix))
	if err != nil {
		return input, fmt.Errorf("parse seeds: %v", err)
	}
	var currMapLines []string
	for i, line := range lines[1:] {
		if !strings.HasSuffix(line, mapSuffix) {
			currMapLines = append(currMapLines, line)
			continue
		}
		if len(currMapLines) == 0 {
			continue
		}
		newMap, err := parseMap(currMapLines)
		if err != nil {
			return input, fmt.Errorf("parse map %d: %v", i, err)
		}
		input.Maps = append(input.Maps, newMap)
		currMapLines = nil
	}
	if len(currMapLines) == 0 {
		return input, fmt.Errorf("last map had no lines")
	}
	newMap, err := parseMap(currMapLines)
	if err != nil {
		return input, fmt.Errorf("parse last map: %v", err)
	}
	input.Maps = append(input.Maps, newMap)
	return input, nil
}

func One(lines []string) (int, error) {
	input, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}
	// we start with the seeds and iterate over the maps
	locations := input.Seeds
	for _, m := range input.Maps {
		for i := 0; i < len(locations); i++ {
			locations[i] = m.Apply(locations[i])
		}
	}

	// now we return the smallest location
	min := locations[0]
	for _, loc := range locations[1:] {
		if loc < min {
			min = loc
		}
	}
	return min, nil
}
