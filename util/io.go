package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read lines: %v", err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newLine := scanner.Text()
		if newLine == "" {
			continue
		}
		lines = append(lines, strings.TrimSpace(newLine))
	}
	return lines, nil
}
