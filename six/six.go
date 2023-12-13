package six

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Opsi/adventofcode/util"
)

type Race struct {
	Time     int
	Distance int
}

func (r Race) WaysToWin() int {
	smallestBTToWin, _ := sort.Find(r.Time/2, func(buttonTime int) int {
		if distance(r.Time, buttonTime) > r.Distance {
			return -1
		}
		return 1
	})
	return (r.Time + 1) - 2*smallestBTToWin
}

func distance(raceTime, buttonTime int) int {
	return (raceTime - buttonTime) * buttonTime
}

func parseInput(lines []string) ([]Race, error) {
	if len(lines) != 2 {
		return nil, fmt.Errorf("input must have 2 lines but has %d", len(lines))
	}
	timePrefix := "Time:"
	if !strings.HasPrefix(lines[0], timePrefix) {
		return nil, fmt.Errorf("invalid time line: %s", lines[0])
	}
	timeLine := strings.TrimPrefix(lines[0], timePrefix)
	times, err := util.ParseSpaceSeparatedInts(timeLine)
	if err != nil {
		return nil, fmt.Errorf("parse times: %v", err)
	}
	distancePrefix := "Distance:"
	if !strings.HasPrefix(lines[1], distancePrefix) {
		return nil, fmt.Errorf("invalid distance line: %s", lines[1])
	}
	distanceLine := strings.TrimPrefix(lines[1], distancePrefix)
	distances, err := util.ParseSpaceSeparatedInts(distanceLine)
	if err != nil {
		return nil, fmt.Errorf("parse distances: %v", err)
	}
	if len(times) != len(distances) {
		return nil, fmt.Errorf("input has %d times and %d distances", len(times), len(distances))
	}
	races := make([]Race, len(times))
	for i := range times {
		races[i] = Race{
			Time:     times[i],
			Distance: distances[i],
		}
	}
	return races, nil
}

func One(lines []string) (int, error) {
	races, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}
	product := 1
	for _, race := range races {
		product *= race.WaysToWin()
	}
	return product, nil
}

func Two(lines []string) (int, error) {
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, " ", "")
	}
	races, err := parseInput(lines)
	if err != nil {
		return 0, fmt.Errorf("parse input: %v", err)
	}
	if len(races) != 1 {
		return 0, fmt.Errorf("there should only be one race in two")
	}
	return races[0].WaysToWin(), nil
}
