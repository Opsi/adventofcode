package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseSpaceSeparatedInts(s string) ([]int, error) {
	var ints []int
	for _, part := range strings.Split(strings.TrimSpace(s), " ") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		i, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid int '%s': %v", part, err)
		}
		ints = append(ints, i)
	}
	return ints, nil
}
