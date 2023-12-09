package nine

import (
	"fmt"
	"strconv"
	"strings"
)

func lineValueOne(line string) (int, error) {
	return 0, nil
}

func nextSequence(sequence []int) ([]int, error) {
	return nil, nil
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
	var builder strings.Builder
	for _, i := range ints {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(" ")
	}
	return builder.String()
}
