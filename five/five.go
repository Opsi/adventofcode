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

type Map []MapRange

type Range struct {
	Start  int
	Length int
}

// ExclusiveEnd returns the first index that is not part of the range
func (r Range) ExclusiveEnd() int {
	return r.Start + r.Length
}

func (r Range) InclusiveEnd() int {
	return r.ExclusiveEnd() - 1
}

func (r Range) Shift(delta int) Range {
	return Range{r.Start + delta, r.Length}
}

type MapRange struct {
	From  Range
	Delta int
}

func (r Range) IsInRange(i int) bool {
	return i >= r.Start && i < r.ExclusiveEnd()
}

func (m Map) Apply(i int) int {
	for _, r := range m {
		if r.From.IsInRange(i) {
			return i + r.Delta
		}
	}
	return i
}

func (mr MapRange) ApplyRange(r Range) (applied, unapplied []Range) {
	switch {
	// r is left of mr
	case r.InclusiveEnd() < mr.From.Start:
		unapplied = []Range{r}

	// r is right of mr
	case r.Start > mr.From.InclusiveEnd():
		unapplied = []Range{r}

	// r is completely in mr
	case mr.From.IsInRange(r.Start) && mr.From.IsInRange(r.InclusiveEnd()):
		applied = []Range{{r.Start + mr.Delta, r.Length}}

	// mr is completely in r
	case r.IsInRange(mr.From.Start) && r.IsInRange(mr.From.InclusiveEnd()):
		left := Range{r.Start, mr.From.Start - r.Start}
		if left.Length > 0 {
			unapplied = append(unapplied, left)
		}
		right := Range{mr.From.InclusiveEnd() + 1, r.InclusiveEnd() - mr.From.InclusiveEnd()}
		if right.Length > 0 {
			unapplied = append(unapplied, right)
		}
		applied = []Range{mr.From.Shift(mr.Delta)}

	// r has a left overlap
	case mr.From.IsInRange(r.InclusiveEnd()):
		left := Range{r.Start, mr.From.Start - r.Start}
		if left.Length > 0 {
			unapplied = append(unapplied, left)
		}
		applied = []Range{{mr.From.Start + mr.Delta, r.Length - left.Length}}

	// r has a right overlap
	case mr.From.IsInRange(r.Start):
		right := Range{mr.From.ExclusiveEnd(), r.ExclusiveEnd() - mr.From.ExclusiveEnd()}
		if right.Length > 0 {
			unapplied = append(unapplied, right)
		}
		applied = []Range{{r.Start + mr.Delta, r.Length - right.Length}}
	}
	return applied, unapplied
}

func parseRange(s string) (MapRange, error) {
	var r MapRange
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return r, fmt.Errorf("invalid range: %s", s)
	}
	destStart, err := strconv.Atoi(parts[0])
	if err != nil {
		return r, fmt.Errorf("invalid dest start: %v", err)
	}
	r.From.Start, err = strconv.Atoi(parts[1])
	if err != nil {
		return r, fmt.Errorf("invalid src start: %v", err)
	}
	r.From.Length, err = strconv.Atoi(parts[2])
	if err != nil {
		return r, fmt.Errorf("invalid length: %v", err)
	}
	r.Delta = destStart - r.From.Start
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

func seedsToRanges(seeds []int) ([]Range, error) {
	if len(seeds) == 0 {
		return nil, fmt.Errorf("no seeds")
	}
	if len(seeds)%2 != 0 {
		return nil, fmt.Errorf("odd number of seeds")
	}
	ranges := make([]Range, len(seeds)/2)
	for i := 0; i < len(ranges); i++ {
		ranges[i] = Range{seeds[2*i], seeds[2*i+1]}
	}
	return ranges, nil
}

func Two(lines []string) (int, error) {
	input, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}

	ranges, err := seedsToRanges(input.Seeds)
	if err != nil {
		return 0, fmt.Errorf("seeds to ranges: %v", err)
	}

	restRanges := ranges
	for _, m := range input.Maps {
		var mappedRanges []Range
		for _, mapRange := range m {
			var newRestRanges []Range
			for _, r := range restRanges {
				newApplied, newUnapplied := mapRange.ApplyRange(r)
				mappedRanges = append(mappedRanges, newApplied...)
				newRestRanges = append(newRestRanges, newUnapplied...)
			}
			restRanges = newRestRanges
		}
		// all rest ranges are now mapped (basically with delta 0)
		restRanges = append(restRanges, mappedRanges...)
	}

	// now we return the smallest location
	min := restRanges[0].Start
	for _, r := range restRanges[1:] {
		if r.Start < min {
			min = r.Start
		}
	}
	return min, nil
}
