package adventofcode

import (
	"fmt"
	"strings"
)

func firstDigit(line string) (int, error) {
	return 0, nil
}

func lastDigit(line string) (int, error) {
	return 0, nil
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

func CalibrationValue(document string) (int, error) {
	trimmed := strings.TrimSpace(document)
	sum := 0
	for i, line := range strings.Split(trimmed, "\n") {
		value, err := lineValue(line)
		if err != nil {
			return 0, fmt.Errorf("line value of line %d: %v", i, err)
		}
		fmt.Printf("line %d: %d\n", i, value)
	}
	return sum, nil
}
