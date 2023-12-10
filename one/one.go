package one

import (
	"fmt"
	"regexp"
)

func One(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		value, err := lineValueOne(line)
		if err != nil {
			return 0, fmt.Errorf("line value of line %d: %v", i, err)
		}
		sum += value
	}
	return sum, nil
}

func Two(lines []string) (int, error) {
	sum := 0
	for i, line := range lines {
		value, err := lineValueTwo(line)
		if err != nil {
			return 0, fmt.Errorf("line value of line %d: %v", i, err)
		}
		sum += value
	}
	return sum, nil
}

func lineValueOne(line string) (int, error) {
	first, err := firstDigitOne(line)
	if err != nil {
		return 0, fmt.Errorf("first digit: %v", err)
	}
	last, err := lastDigitOne(line)
	if err != nil {
		return 0, fmt.Errorf("last digit: %v", err)
	}
	return 10*first + last, nil
}

func firstDigitOne(line string) (int, error) {
	for _, c := range line {
		if c >= '0' && c <= '9' {
			return int(c - '0'), nil
		}
	}
	return 0, fmt.Errorf("no digit found")
}

func lastDigitOne(line string) (int, error) {
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c >= '0' && c <= '9' {
			return int(c - '0'), nil
		}
	}
	return 0, fmt.Errorf("no digit found")
}

var valueMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func lineValueTwo(line string) (int, error) {
	first, err := firstDigitTwo(line)
	if err != nil {
		return 0, fmt.Errorf("first digit: %v", err)
	}
	last, err := lastDigitTwo(line)
	if err != nil {
		return 0, fmt.Errorf("last digit: %v", err)
	}
	return 10*first + last, nil
}

var reFirst = regexp.MustCompile("(one|1|two|2|three|3|four|4|five|5|six|6|seven|7|eight|8|nine|9)")
var reLast = regexp.MustCompile("(9|enin|8|thgie|7|neves|6|xis|5|evif|4|ruof|3|eerht|2|owt|1|eno)")

func firstDigitTwo(line string) (int, error) {
	res := reFirst.FindString(line)
	switch len(res) {
	case 0:
		return 0, fmt.Errorf("no digit found")
	case 1:
		return int(res[0] - '0'), nil
	default:
		return valueMap[res], nil
	}
}

func lastDigitTwo(line string) (int, error) {
	res := reLast.FindString(reversed(line))
	switch len(res) {
	case 0:
		return 0, fmt.Errorf("no digit found")
	case 1:
		return int(res[0] - '0'), nil
	default:
		return valueMap[reversed(res)], nil
	}
}

func reversed(input string) string {
	// Convert the string to a slice of runes
	inputRunes := []rune(input)

	// Get the length of the slice
	length := len(inputRunes)

	// Reverse the slice
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		inputRunes[i], inputRunes[j] = inputRunes[j], inputRunes[i]
	}

	// Convert the slice of runes back to a string
	return string(inputRunes)
}
