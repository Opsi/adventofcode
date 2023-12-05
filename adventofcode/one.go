package adventofcode

import (
	"fmt"
	"strings"
)

func firstDigit(line string) (int, error) {
	for _, c := range line {
		if c >= '0' && c <= '9' {
			return int(c - '0'), nil
		}
	}
	return 0, fmt.Errorf("no digit found")
}

func lastDigit(line string) (int, error) {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c >= '0' && c <= '9' {
			return int(c - '0'), nil
		}
	}
	return 0, fmt.Errorf("no digit found")
}

func lineValue(line string) (int, error) {
	first, err := firstDigit(line)
	if err != nil {
		return 0, fmt.Errorf("first digit: %v", err)
	}
	last, err := lastDigit(line)
	if err != nil {
		return 0, fmt.Errorf("last digit: %v", err)
	}
	return 10*first + last, nil
}

func OneOne(document string) (int, error) {
	trimmed := strings.TrimSpace(document)
	sum := 0
	for i, line := range strings.Split(trimmed, "\n") {
		value, err := lineValue(line)
		if err != nil {
			return 0, fmt.Errorf("line value of line %d: %v", i, err)
		}
		sum += value
	}
	return sum, nil
}
