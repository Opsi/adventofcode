package nine

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func One(document string) (int, error) {
	trimmed := strings.TrimSpace(document)
	sum := 0
	for i, line := range strings.Split(trimmed, "\n") {
		value, err := lineValueOne(line)
		if err != nil {
			return 0, fmt.Errorf("line value of line %d: %v", i, err)
		}
		sum += value
	}
	return sum, nil
}

func Two(document string) (int, error) {
	trimmed := strings.TrimSpace(document)
	sum := 0
	for i, line := range strings.Split(trimmed, "\n") {
		// we just reverse the line and use the same function
		intStrings := strings.Split(line, " ")
		slices.Reverse(intStrings)
		line = strings.Join(intStrings, " ")
		value, err := lineValueOne(line)
		if err != nil {
			return 0, fmt.Errorf("line value of line %d: %v", i, err)
		}
		sum += value
	}
	return sum, nil
}

func lineValueOne(line string) (int, error) {
	sequences := make([][]int, 1)
	var err error
	sequences[0], err = toInts(line)
	if err != nil {
		return 0, fmt.Errorf("parse line: %v", err)
	}
	curr := sequences[0]
	for !isZeroSequence(curr) {
		curr, err = nextSequence(curr)
		if err != nil {
			return 0, fmt.Errorf("next sequence: %v", err)
		}
		sequences = append(sequences, curr)
	}
	for i := len(sequences) - 1; i > 0; i-- {
		bottom := sequences[i]
		top := sequences[i-1]
		newVal := top[len(top)-1] + bottom[len(bottom)-1]
		sequences[i-1] = append(top, newVal)
	}
	return sequences[0][len(sequences[0])-1], nil
}

func nextSequence(sequence []int) ([]int, error) {
	switch {
	case len(sequence) == 0:
		return nil, fmt.Errorf("empty sequence")
	case len(sequence) == 1:
		return nil, fmt.Errorf("sequence with one element")
	default:
		newSeq := make([]int, len(sequence)-1)
		for i := 0; i < len(newSeq); i++ {
			newSeq[i] = sequence[i+1] - sequence[i]
		}
		return newSeq, nil
	}
}

func isZeroSequence(sequence []int) bool {
	for _, i := range sequence {
		if i != 0 {
			return false
		}
	}
	return true
}

func toInts(sequence string) ([]int, error) {
	subs := strings.Split(sequence, " ")
	ints := make([]int, len(subs))
	for i, sub := range subs {
		trimmed := strings.TrimSpace(sub)
		var err error
		ints[i], err = strconv.Atoi(trimmed)
		if err != nil {
			return nil, fmt.Errorf("parse int '%s': %v", trimmed, err)
		}
	}
	return ints, nil
}

func asString(ints []int) string {
	asStrings := make([]string, len(ints))
	for i, int := range ints {
		asStrings[i] = strconv.Itoa(int)
	}
	return strings.Join(asStrings, " ")
}
